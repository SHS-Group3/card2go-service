package model

import "gorm.io/gorm"

type Hotel struct {
	gorm.Model
	Name        string
	Description string
	Address     string
}
