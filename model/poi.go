package model

import "gorm.io/gorm"

type POI struct {
	gorm.Model
	Location
}
