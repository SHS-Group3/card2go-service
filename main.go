package main

import (
	"card2go_service/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.RegisterAPI(app)

	app.Listen(":8080")
}
