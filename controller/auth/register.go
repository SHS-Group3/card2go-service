package auth

import (
	"card2go_service/model"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// /auth/register endpoint
// requires database
func HandleRegister(c *fiber.Ctx) error {
	DB := c.Locals("database").(*gorm.DB)

	type registrationInfo struct {
		Username string `json: "username"`
		Password string `json: "password"`
	}
	var info registrationInfo

	if err := c.BodyParser(&info); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	// check if username and password field were initialized
	if info.Username == "" || info.Password == "" {
		err := errors.New("missing fields")
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	// check if username is unique
	var count int64

	DB.Where("username = ?", info.Username).Model(&model.User{}).Count(&count)

	if count > 0 {
		err := errors.New("username taken")
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	var user model.User

	user.Username = info.Username
	user.Password = info.Password
	user.Admin = false

	DB.Create(&user)

	c.Status(http.StatusCreated).JSON(fiber.Map{
		"id": user.ID,
	})
	return nil
}
