package model

import "gorm.io/gorm"

type Booking struct {
	gorm.Model

	ClientID uint
	Client   User `json:"user_id" gorm:"not null"`

	LocationID   uint `json:"location_id" gorm:"not null"`
	LocationType string

	PackageID uint `json:"package_id" gorm:"not null"`
	Package   Package
}
