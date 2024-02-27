package server

import "github.com/gofiber/fiber/v2"

func New() *fiber.App {
	server := fiber.New()

	SetupRoutes(server)

	return server
}
