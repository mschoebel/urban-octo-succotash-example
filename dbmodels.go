package main

import (
	"gorm.io/gorm"
)

type person struct {
	gorm.Model

	Name string
	City string
}
