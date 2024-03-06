package packages

import (
	auth "card2go_service/authentication"
	"card2go_service/database"
	"card2go_service/model"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func HandleBooking(c *fiber.Ctx) error {
	//TODO: add from and to date in request to add to booking
	DB, err := database.GetConnection()
	if err != nil {
		fmt.Errorf("Error connecting to database %s", err.Error())
		return err
	}

	// TODO: create authorization required header middleware
	// #region authorization
	authorization := c.Get("Authorization")
	if authorization == "" {
		return fiber.ErrUnauthorized
	}

	str := strings.Split(authorization, " ")
	if str[0] != "Bearer" {
		return fiber.NewError(fiber.ErrBadRequest.Code, fmt.Sprintf("invalid authorization header! %s", err.Error()))
	}

	userID, err := auth.GetIDFromToken(str[1])
	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, fmt.Sprintf("error while parsing token %s", err.Error()))
	}
	// #endregion authorization

	id, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "invalid id")
	}

	var pkg model.Package
	err = DB.Where("id = ?", id).Limit(1).Find(&pkg).Error
	if err != nil {
		return fiber.NewError(fiber.ErrNotFound.Code, "package not found")
	}

	var booking model.Booking
	booking.UserID = userID
	booking.Package = pkg
	DB.Create(&booking)
	association := DB.Model(&booking).Association("Location")
	offerer, _ := pkg.GetOfferer()

	//TODO: this is awful
	type responseBooking struct {
		id          uint
		user_id     uint
		location_id uint
	}

	var response responseBooking

	//TODO: contemplate life choices and merge hotels and poi
	switch pkg.OffererType {
	case "hotels":
		hotel := offerer.(model.Hotel)
		association.Append(&hotel)

		response.id = booking.ID
		response.location_id = hotel.ID
		response.user_id = userID

		c.Status(fiber.StatusOK).JSON(response)
	case "pois":
		poi := offerer.(model.POI)
		association.Append(&poi)

		response.id = booking.ID
		response.location_id = poi.ID
		response.user_id = userID

		c.Status(fiber.StatusOK).JSON(response)
	}

	return nil
}
