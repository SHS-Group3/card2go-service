package main

import (
	"card2go_service/config"
	"card2go_service/database"
	"card2go_service/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.LoadFromEnv()

	database.Connect()

	routes.RegisterAPI(app)

	app.Listen(":8080")
}
