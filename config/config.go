package config

import "os"

var (
	Host       = "localhost"
	DBUser     = "postgres"
	DBPassword = "user"
	DBName     = "card2go"
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
}
