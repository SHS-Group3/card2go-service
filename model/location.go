package model

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	Name        string
	Description string
	Address     string

	Packages []Package
}
