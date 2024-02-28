package model

import "gorm.io/gorm"

type Package struct {
	gorm.Model
	LocationID  uint
	Title       string
	Description string
}
