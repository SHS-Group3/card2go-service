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

	var bookings []model.Booking
	DB.Where("user_id = ?", user.ID).Preload("User").Preload("Destination").Preload("Packages").Find(&bookings)

	c.Status(fiber.StatusOK).JSON(bookings)

	return nil
}
