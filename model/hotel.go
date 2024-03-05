package model

import (
	"card2go_service/database"

	"gorm.io/gorm"
)

type Hotel struct {
	gorm.Model
	Name        string  `json:"name" gorm:"not null"`
	Description string  `json:"description" gorm:"not null"`
	Address     string  `json:"address" gorm:"not null"`
	Ratings     float32 `json:"ratings"`
	Beds        int     `json:"beds"`
	Rooms       int     `json:"rooms"`

	Bookings []Booking `json:"bookings" gorm:"polymorphic:Location"`
	Packages []Package `json:"packages" gorm:"polymorphic:Offerer"`
}

func (hotel *Hotel) GetBookings() ([]Booking, error) {
	DB, err := database.GetConnection()
	if err != nil {
		return nil, err
	}

	var bookings []Booking

	err = DB.Model(hotel).Association("Bookings").Find(&bookings)

	return bookings, err
}

func (hotel *Hotel) GetPackages() ([]Package, error) {
	DB, err := database.GetConnection()
	if err != nil {
		return nil, err
	}

	var packages []Package

	err = DB.Model(hotel).Association("Packages").Find(&packages)

	return packages, err
}
