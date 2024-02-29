package auth

import (
	"card2go_service/database"
	"card2go_service/model"
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type RegistrationInfo struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

// /auth/register endpoint
func HandleRegister(c *fiber.Ctx) error {
	DB, err := database.GetConnection()

	if err != nil {
		fmt.Errorf("Error connecting to database %s", err.Error())
		return err
	}

	var info RegistrationInfo

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
