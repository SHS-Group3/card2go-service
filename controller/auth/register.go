package auth

import (
	"card2go_service/database"
	"card2go_service/model"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type RegistrationInfo struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

func HandleRegister(c *fiber.Ctx) error {
	var info RegistrationInfo

	if err := c.BodyParser(&info); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	if info.Username == "" || info.Password == "" {
		err := errors.New("missing fields")
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	var count int64
	database.DB.Where("username = ?", info.Username).Model(&model.User{}).Count(&count)

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

	database.DB.Create(&user)

	c.Status(http.StatusCreated).JSON(fiber.Map{
		"id": user.ID,
	})
	return nil
}
