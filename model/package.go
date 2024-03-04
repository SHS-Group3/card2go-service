package model

import "gorm.io/gorm"

type Package struct {
	gorm.Model
	OffererID   uint
	OffererType string
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
}
