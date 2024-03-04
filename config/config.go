package config

import (
	"log"
	"os"
	"strconv"
)

var (
	Host       = "localhost"
	DBUser     = "postgres"
	DBPassword = "user"
	DBName     = "card2go"
	DBPort     = 5432
	TokenKey   []byte
)

func LoadFromEnv() {
	host, defined := os.LookupEnv("HOST")
	if defined {
		Host = host
	}
	dbUser, defined := os.LookupEnv("DBUSER")
	if defined {
		DBUser = dbUser
	}
	dbPassword, defined := os.LookupEnv("PASSWORD")
	if defined {
		DBPassword = dbPassword
	}
	dbName, defined := os.LookupEnv("DBNAME")
	if defined {
		DBName = dbName
	}
	dbPort, defined := os.LookupEnv("DBPORT")
	if defined {
		port, err := strconv.ParseInt(dbPort, 10, 32)
		if err != nil {
			log.Fatal("Error while converting port to int ", err.Error())
		}
		DBPort = int(port)
	}
	tokenKey, defined := os.LookupEnv("SECRETKEY")
	if defined {
		TokenKey = []byte(tokenKey)
	} else {
		log.Fatal("No secret key defined!")
	}
}
