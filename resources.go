package main

import (
	uos "github.com/mschoebel/urban-octo-succotash"
)

type personResource struct {
}

func (r personResource) Name() string {
	return "person"
}

func (r personResource) Read(id string) (interface{}, error) {
	var p person

	err := uos.DB.First(&p, id).Error
	if err != nil {
		return nil, err
	}

	return p, nil
}
