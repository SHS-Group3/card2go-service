package middleware

import (
	auth "card2go_service/authentication"
	"card2go_service/database"
	"card2go_service/model"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func RequireAuth(c *fiber.Ctx) error {
	DB, err := database.GetConnection()
	if err != nil {
		return err
	}

	authorization := c.Get("Authorization")
	if authorization == "" {
		return fiber.NewError(fiber.ErrUnauthorized.Code, "no auth token provided")
	}

	if str := strings.Split(authorization, " "); str[0] == "Bearer" {
		id, err := auth.GetIDFromToken(str[1])
		if err != nil {
			return fiber.NewError(fiber.ErrBadRequest.Code, fmt.Sprintf("error while parsing token %s", err.Error()))
		}

		var user model.User
		if err = DB.Limit(1).Find(&user, id).Error; err != nil {
			return err
		}
		c.Locals("user", user)
		return c.Next()
	} else {
		return fiber.NewError(fiber.ErrBadRequest.Code, "invalid authorization header")
	}
}

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
