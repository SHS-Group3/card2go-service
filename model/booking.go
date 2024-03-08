package model

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	BookingID uint `gorm:"unique;autoIncrement"`

	On time.Time

	UserID uint
	User   User

	DestinationID uint
	Destination   Destination

	PackageID *uint
	Package   *Package
}
