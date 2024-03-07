package model

import (
	"gorm.io/gorm"
)

type Destination struct {
	gorm.Model
	Name        string  `json:"name" gorm:"not null"`
	Description string  `json:"description" gorm:"not null"`
	Address     string  `json:"address" gorm:"not null"`
	Ratings     float32 `json:"ratings"`
	IsLodging   bool    `json:"is_lodging"`
	Beds        int     `json:"beds"`
	Rooms       int     `json:"rooms"`

	Bookings []Booking `json:"bookings"`
	Packages []Package `json:"packages"`
}
