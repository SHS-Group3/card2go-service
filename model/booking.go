package model

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model

	On time.Time

	UserID uint `gorm:"not null"`

	DestinationID uint `gorm:"not null"`

	PackageID uint
}
