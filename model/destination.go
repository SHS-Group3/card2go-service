package model

import (
	"gorm.io/gorm"
)

type Destination struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Address     string `json:"address" gorm:"not null"`

	IsLodging bool    `json:"is_lodging"`
	Ratings   float64 `json:"ratings"`
	Beds      int     `json:"beds"`
	Rooms     int     `json:"rooms"`

	Packages []Package `json:"packages"`
}

// func NewDestination() Destination {
// 	return Destination{
// 		Packages: []Package{
// 			Package{
// 				Title:       "Plan",
// 				Description: "Plan your trip",
// 			},
// 		},
// 	}
// }
