package main

import (
	uos "github.com/mschoebel/urban-octo-succotash"
)

type hintDialog struct{}

func (d hintDialog) Name() string {
	return "hint"
}

func (d hintDialog) DisplayTitle() string {
	return "Hint"
}

func (d hintDialog) Footer() uos.DialogFooter {
	return uos.DialogFooter{
		Left: []uos.DialogFooterElement{
			uos.DialogFooterText("This button does not work: ", "warning"),
			uos.DialogFooterButton("Help", "info"),
		},
		Right: []uos.DialogFooterElement{
			uos.DialogFooterPrimaryButton("Ok").Closing(),
		},
	}
}

type dataDialog struct{}

func (d dataDialog) Name() string {
	return "data"
}

func (d dataDialog) DisplayTitle() string {
	return "Data input form"
}

func (d dataDialog) Footer() uos.DialogFooter {
	return uos.DialogFooter{
		Right: []uos.DialogFooterElement{
			uos.DialogFooterButton("Cancel", "white").Closing(),
			uos.DialogFooterPrimaryButton("Save").Saving("data"),
		},
	}
}

type loginDialog struct{}

func (d loginDialog) Name() string {
	return "login"
}

func (d loginDialog) DisplayTitle() string {
	return "Login"
}

func (d loginDialog) Footer() uos.DialogFooter {
	return uos.DialogFooter{
		Right: []uos.DialogFooterElement{
			uos.DialogFooterButton("Cancel", "white").Closing(),
			uos.DialogFooterPrimaryButton("Login").Saving("login"),
		},
	}
}
