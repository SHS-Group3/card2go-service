package database

import (
	"card2go_service/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// gorm connections are pooled so it is fine to use a connection for each operation?
func GetConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable", config.Host, config.DBUser, config.DBPassword, config.DBPort, config.DBName)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

// func Connection() *gorm.DB {
// 	db, err := GetConnection()
// 	err
// }
