package model

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	Client    User
	Location  Location
	PackageID uint
	Package   Package
}
