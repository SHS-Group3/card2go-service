package model

import "gorm.io/gorm"

type Hotel struct {
	gorm.Model
	Location
	Rates float32
}
