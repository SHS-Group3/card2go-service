package main

import (
	"card2go_service/config"
	"card2go_service/database"
	"card2go_service/routes"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	if err := godotenv.Load(); err != nil {
		fmt.Errorf("Error loading .env file ")
	}

	config.LoadFromEnv()

	if err := database.SetupDB(); err != nil {
		log.Fatal("Error when connecting to database ", err)
	}

	routes.RegisterAPI(app)

	app.Listen(":8080")
}
