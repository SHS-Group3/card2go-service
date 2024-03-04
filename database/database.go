package database

import (
	"card2go_service/config"
	"card2go_service/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// setup database schema
func SetupDB() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable", config.Host, config.DBUser, config.DBPassword, config.DBPort, config.DBName)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	// create schemas/update schemas on the database
	DB.AutoMigrate(&model.User{}, &model.POI{}, &model.Booking{}, &model.Package{})

	return nil
}

func GetConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable", config.Host, config.DBUser, config.DBPassword, config.DBPort, config.DBName)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

// func Connection() *gorm.DB {
// 	db, err := GetConnection()
// 	err
// }
