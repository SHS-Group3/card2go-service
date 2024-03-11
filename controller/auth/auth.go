package auth

import (
	auth "card2go_service/authentication"
	"card2go_service/model"
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
		return err
	}

	var user model.User

	DB.Where("username = ? AND password = ?", info.Username, info.Password).Select("id").Limit(1).Find(&user)

	if user.ID == 0 {
		return fiber.NewError(fiber.StatusUnauthorized, "authentication failed")
	}

	c.Status(http.StatusOK).JSON(fiber.Map{
		"token": auth.CreateSignedToken(user.ID),
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
		return err
	}

	// check if username and password field were initialized
	if info.Username == "" || info.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "missing username or password")
	}

	// check if username is unique
	var count int64

	DB.Where("username = ?", info.Username).Model(&model.User{}).Count(&count)

	if count > 0 {
		return fiber.NewError(fiber.StatusConflict, "username already taken")
	}

	user := model.User{
		Username: info.Username,
		Password: info.Password,
		Admin:    false,
	}

	DB.Create(&user)

	c.Status(http.StatusCreated).JSON(fiber.Map{
		"id": user.ID,
	})
	return nil
}
