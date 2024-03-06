package poi

import (
	auth "card2go_service/authentication"
	"card2go_service/database"
	"card2go_service/model"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

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

	var pois []model.POI
	err = DB.Model(&model.POI{}).Preload("Bookings").Preload("Packages").Order("created_at desc").Limit(20).Offset(offset).Find(&pois).Error
	if err != nil {
		return err
	}

	type returnPOI struct {
		ID          uint
		Name        string
		Description string
		Address     string

		Booked   int
		Packages int
	}

	var returnPOIs []returnPOI
	for _, i := range pois {
		var a returnPOI

		a.ID = i.ID
		a.Name = i.Name
		a.Address = i.Address
		a.Description = i.Description
		a.Booked = len(i.Bookings)
		a.Packages = len(i.Packages)

		returnPOIs = append(returnPOIs, a)
	}

	c.JSON(returnPOIs)

	return nil
}

func HandlePOI(c *fiber.Ctx) error {
	DB, err := database.GetConnection()
	if err != nil {
		fmt.Errorf("Error connecting to database %s", err.Error())
		return err
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "invalid id")
	}

	var poi model.POI
	DB.Model(&poi).Where("id = ?", id).Preload("Packages").Limit(1).Find(&poi)

	if poi.ID == 0 {
		return fiber.NewError(fiber.ErrBadRequest.Code, "poi id not found")
	}

	// check if booked by user
	booked := false

	authorization := c.Get("Authorization")
	if authorization != "" {
		if str := strings.Split(authorization, " "); str[0] == "Bearer" {
			id, err := auth.GetIDFromToken(str[1])
			if err != nil {
				return fiber.NewError(fiber.ErrBadRequest.Code, fmt.Sprintf("error while parsing token %s", err.Error()))
			}

			var count int64
			DB.Model(&model.Booking{}).Where("location_id = ? AND user_id = ?", poi.ID, id).Count(&count)

			booked = count < 0
		}
	}

	type returnPackage struct {
		id          uint
		title       string
		description string
	}

	type returnPOI struct {
		id          uint
		name        string
		description string
		address     string
		booked      bool

		packages []returnPackage
	}

	var a returnPOI

	a.id = poi.ID
	a.name = poi.Name
	a.address = poi.Address
	a.description = poi.Description
	a.booked = booked

	for _, i := range poi.Packages {
		var p returnPackage

		p.id = i.ID
		p.title = i.Title
		p.description = i.Description

		a.packages = append(a.packages, p)
	}

	c.Status(fiber.StatusOK).JSON(a)

	return nil
}
