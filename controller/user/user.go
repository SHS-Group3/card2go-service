package user

import (
	"card2go_service/model"

	"github.com/gofiber/fiber/v2"
)

// GET /user
// requires auth
func HandleInfo(c *fiber.Ctx) error {
	user := c.Locals("user").(model.User)
	type returnInfo struct {
		Username      string
		Email         string
		ContactNumber string
	}

	c.JSON(returnInfo{
		Username:      user.Username,
		Email:         user.Email,
		ContactNumber: user.ContactNumber,
	})

	return nil
}
