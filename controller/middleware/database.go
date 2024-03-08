package middleware

import (
	"card2go_service/database"

	"github.com/gofiber/fiber/v2"
)

// middleware to make a request have a database connection
// assigns "database" local to the database instance
func RequireDatabase(c *fiber.Ctx) error {
	DB, err := database.GetConnection()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to connect to database!")
	}

	c.Locals("database", DB)

	return c.Next()
}
