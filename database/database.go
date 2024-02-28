package database

import (
	"card2go_service/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.DBUser, config.DBPassword, config.DBName)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	_ = err
}
