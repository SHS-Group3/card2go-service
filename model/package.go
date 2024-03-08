package model

import (
	"gorm.io/gorm"
)

type Package struct {
	gorm.Model
	DestinationID uint   `gorm:"not null"`
	Title         string `json:"title" gorm:"not null"`
	Description   string `json:"description" gorm:"not null"`
}
