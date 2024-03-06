package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null"`
	Password string `gorm:"not null;->:false;<-"`
	Admin    bool   `json:"admin" gorm:"not null"`

	Bookings []Booking
}
