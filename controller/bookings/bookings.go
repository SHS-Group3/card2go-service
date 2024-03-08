package bookings

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GET /bookings
// requires authorization
// requires database
func HandleBookings(c *fiber.Ctx) error {
	DB := c.Locals("database").(*gorm.DB)
	_ = DB
	return nil
}
