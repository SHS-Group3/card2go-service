package destinations

import (
	"card2go_service/database"
	"card2go_service/model"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func HandleDestination(c *fiber.Ctx) error {
	DB, err := database.GetConnection()
	if err != nil {
		fmt.Errorf("Error connecting to database %s", err.Error())
		return err
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "invalid id")
	}

	var destination model.Destination
	DB.Model(&destination).Where("id = ?", id).Preload("Packages").Limit(1).Find(&destination)

	if destination.ID == 0 {
		return fiber.NewError(fiber.ErrBadRequest.Code, "destination id not found")
	}

	c.Status(fiber.StatusOK).JSON(destination)

	return nil
}

func HandleFeed(c *fiber.Ctx) error {
	DB, err := database.GetConnection()
	if err != nil {
		fmt.Errorf("Error connecting to database %s", err.Error())
		return err
	}

	offset, err := c.ParamsInt("page", 1)
	offset = (offset - 1) * 20

	if err != nil {
		return err
	} else if offset < 0 {
		return fiber.NewError(fiber.ErrBadRequest.Code, "invalid page")
	}

	var dest []model.Destination
	err = DB.Model(&model.Destination{}).Order("created_at desc").Limit(20).Offset(offset).Find(&dest).Error
	if err != nil {
		return err
	}

	c.JSON(dest)

	return nil
}
