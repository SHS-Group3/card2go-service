package model

import "gorm.io/gorm"

type Hotel struct {
	gorm.Model
	Name        string  `json:"name" gorm:"not null"`
	Description string  `json:"description" gorm:"not null"`
	Address     string  `json:"address" gorm:"not null"`
	Ratings     float32 `json:"ratings"`
	Beds        int     `json:"beds"`
	Rooms       int     `json:rooms"`

	Bookings []Booking `gorm:"polymorphic:Location"`
	Packages []Package `gorm:"polymorphic:Offerer"`
}
