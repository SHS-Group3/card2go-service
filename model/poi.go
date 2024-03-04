package model

import "gorm.io/gorm"

type POI struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Address     string `json:"address" gorm:"not null"`

	Bookings []Booking `gorm:"polymorphic:Location"`
	Packages []Package `gorm:"polymorphic:Offerer"`
}
