package database

import (
	"card2go_service/config"
	"card2go_service/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.DBUser, config.DBPassword, config.DBName)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	// create schemas on the database
	DB.AutoMigrate(&model.User{}, &model.Hotel{})

	return nil
}
