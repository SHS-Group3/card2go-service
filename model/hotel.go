package model

import "gorm.io/gorm"

type Hotel struct {
	gorm.Model
	Name        string
	Description string
	Address     string
	Ratings     float32
	Beds        int
	Rooms       int

	Packages []Package
}
