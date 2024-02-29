package model

import "gorm.io/gorm"

type Booking struct {
	gorm.Model

	ClientID uint
	Client   User

	LocationID uint
	Location   Location

	PackageID uint
	Package   Package
}
