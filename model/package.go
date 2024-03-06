package model

import (
	"card2go_service/database"
	"log"

	"gorm.io/gorm"
)

type Package struct {
	gorm.Model
	OffererID   uint
	OffererType string
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
}

func (pkg *Package) GetOfferer() (interface{}, error) {
	DB, err := database.GetConnection()
	if err != nil {
		return nil, err
	}

	query := DB.Model(pkg).Limit(1).Association("Offerer")

	// why do i do this to myself
	switch pkg.OffererType {
	case "hotels":
		var hotel Hotel
		err = query.Find(&hotel)
		return hotel, err
	case "pois":
		var poi POI
		err = query.Find(&poi)
		return poi, err
	default:
		log.Panic("wtf")
		return nil, err // aasfaesfwefa
	}
}
