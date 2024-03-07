package model

import (
	"gorm.io/gorm"
)

type Package struct {
	gorm.Model
	DestinationID uint
	Destination   Destination
	Title         string `json:"title" gorm:"not null"`
	Description   string `json:"description" gorm:"not null"`
}
