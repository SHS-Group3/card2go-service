package auth

import (
	"card2go_service/database"
	"card2go_service/model"
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthenticationInfo struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

// /auth endpoint
func HandleAuthentication(c *fiber.Ctx) error {
	DB, err := database.GetConnection()

	if err != nil {
		fmt.Errorf("Error connecting to database %s", err.Error())
		return err
	}

	var info AuthenticationInfo

	if err := c.BodyParser(&info); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
		return err
	}

	var user model.User

	DB.Where("username = ? AND password = ?", info.Username, info.Password).Select("id").Limit(1).First(&user)

	if user.ID == 0 {
		err := errors.New("authentication failed")
		c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	c.Status(http.StatusOK).JSON(fiber.Map{
		"token": "i would put a token here but no!",
	})

	return nil
}
