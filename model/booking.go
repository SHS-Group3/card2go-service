package model

import (
	"card2go_service/database"
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model

	On time.Time

	UserID uint
	User   User `json:"user_id" gorm:"not null"`

	LocationID   uint `json:"location_id" gorm:"not null"`
	LocationType string

	PackageID uint `json:"package_id" gorm:"not null"`
	Package   Package
}

func (booking *Booking) GetLocation() (interface{}, error) {
	DB, err := database.GetConnection()
	if err != nil {
		return nil, err
	}

	query := DB.Model(booking).Limit(1).Association("Location")

	// why do i do this to myself
	switch booking.LocationType {
	case "hotels":
		var hotel Hotel
		err = query.Find(&hotel)
		return hotel, err
	case "pois":
		var poi POI
		err = query.Find(&poi)
		return poi, err
	default:
		return nil, err // aasfaesfwefa
	}
}
