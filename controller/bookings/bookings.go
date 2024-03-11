package bookings

import (
	"card2go_service/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GET /bookings
// requires authorization
// requires database
func HandleBookings(c *fiber.Ctx) error {
	DB := c.Locals("database").(*gorm.DB)
	user := c.Locals("user").(model.User)
	_, _ = DB, user
	return nil
}
