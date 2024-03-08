package auth

import (
	auth "card2go_service/authentication"
	"card2go_service/model"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// POST /auth
// requires database
func HandleAuthentication(c *fiber.Ctx) error {
	DB := c.Locals("database").(*gorm.DB)

	type authenticationInfo struct {
		Username string `json: "username"`
		Password string `json: "password"`
	}
	var info authenticationInfo

	if err := c.BodyParser(&info); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	var user model.User

	DB.Where("username = ? AND password = ?", info.Username, info.Password).Select("id").Limit(1).Find(&user)

	if user.ID == 0 {
		err := errors.New("authentication failed")
		c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
		return nil
	}

	token := auth.CreateSignedToken(user.ID)
	c.Status(http.StatusOK).JSON(fiber.Map{
		"token": token,
	})

	return nil
}

// POST /auth/register
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
