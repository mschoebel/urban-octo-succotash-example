package main

import (
	"errors"
	"strconv"

	uos "github.com/mschoebel/urban-octo-succotash"
	"gorm.io/gorm"
)

type loginForm struct{}

func (f loginForm) Name() string {
	return "login"
}

func (f loginForm) Read(id string) (uos.FormItems, error) {
	// login form does not expect any id
	if id != "" {
		return nil, uos.ErrorFormInvalidRequest
	}

	return uos.FormItems{
			uos.FormItem{
				InputType:     "input",
				InputTypeHTML: "text",
				Name:          "name",
				Placeholder:   "max.mustermann@example.com",
				Label:         "E-Mail",
				IsHorizontal:  true,
				HasFocus:      true,
				Constraints: &uos.FormItemConstraints{
					IsMandatory: true,
					MaxLength:   20,
				},
			},
			uos.FormItem{
				InputType:     "input",
				InputTypeHTML: "password",
				Name:          "password",
				Label:         "Password",
				IsHorizontal:  true,
				Constraints: &uos.FormItemConstraints{
					IsMandatory: true,
				},
			},
		},
		nil
}

func (f loginForm) Save(is string, items uos.FormItems) (*uos.ResponseAction, error) {
	var (
		name     = items.Get("name").Value
		password = items.Get("password").Value
	)

	// try to authenticate user
	user, err := uos.GetAppUser(name, password)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// create user
		user, err = uos.CreateAppUser(name, password)
	}
	if errors.Is(err, uos.ErrorInvalidPassword) {
		return uos.ResponseFormError("Invalid username or password"), nil
	}
	if err != nil {
		return nil, err
	}

	uos.Metrics.CounterInc(mLogins)
	return uos.ResponseSetSessionCookie(user.ID), nil
}

type dataForm struct{}

func (f dataForm) Name() string {
	return "data"
}

func (f dataForm) Read(id string) (uos.FormItems, error) {
	items := uos.FormItems{
		uos.FormItem{
			InputType:     "input",
			InputTypeHTML: "text",
			Name:          "id",
			IsHorizontal:  true,
			IsHidden:      true,
		},
		uos.FormItem{
			InputType:     "input",
			InputTypeHTML: "text",
			Name:          "name",
			Label:         "Name",
			Placeholder:   "Name",
			IsHorizontal:  true,
			HasFocus:      true,
			Constraints: &uos.FormItemConstraints{
				IsMandatory: true,
			},
		},
		uos.FormItem{
			InputType:     "input",
			InputTypeHTML: "text",
			Name:          "city",
			Label:         "City",
			Placeholder:   "City",
			IsHorizontal:  true,
			Constraints: &uos.FormItemConstraints{
				IsMandatory: true,
			},
		},
	}

	if id == "" {
		// return empty form
		return items, nil
	}

	// load form content from database
	var p person
	err := uos.DB.First(&p, id).Error
	if err != nil {
		return nil, err
	}

	return uos.DBExtractForm("person", p, items), nil
}

func (f dataForm) Save(id string, items uos.FormItems) (*uos.ResponseAction, error) {
	p := person{
		Name: items.Get("name").Value,
		City: items.Get("city").Value,
	}

	if id != "" {
		idNr, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			return nil, err
		}
		p.ID = uint(idNr)

		return uos.ResponseRefresh(), uos.DB.Save(&p).Error
	}

	return uos.ResponseMessage("Saved new entry.", "success"), uos.DB.Create(&p).Error
}
