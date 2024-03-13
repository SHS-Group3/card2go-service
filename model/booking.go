package model

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model

	Tickets uint      `gorm:"not null"`
	On      time.Time `gorm:"not null"`

	UserID uint
	User   User `gorm:"not null"`

	DestinationID uint
	Destination   Destination `gorm:"not null"`

	PackageID *uint
	Package   *Package
}
