package model

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model

	On time.Time

	UserID uint
	User   User `json:"user_id" gorm:"not null"`

	DestinationID uint `json:"location_id" gorm:"not null"`
	Destination   Destination

	PackageID uint `json:"package_id" gorm:"not null"`
	Package   Package
}
