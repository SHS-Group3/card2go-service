package bookings

import (
	"card2go_service/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GET /bookings
// requires authorization
// requires database
func HandleBookings(c *fiber.Ctx) error {
	DB := c.Locals("database").(*gorm.DB)
	user := c.Locals("user").(model.User)

	var bookings []model.Booking
	DB.Where("user_id = ?", user.ID).Preload("User").Preload("Destination").Preload("Packages").Find(&bookings)

	type returnUser struct {
		ID       uint
		Username string
	}

	type returnDestination struct {
		ID   uint
		Name string
	}

	type returnPackage struct {
		ID          uint
		Title       string
		Description string
	}

	type returnBooking struct {
		ID          uint `json:"id"`
		User        returnUser
		Destination returnDestination
		Package     returnPackage
		On          time.Time
	}

	returnBookings := []returnBooking{}
	for _, i := range bookings {
		var p *returnPackage
		if i.Package != nil {
			p = &returnPackage{
				ID:          *i.PackageID,
				Title:       i.Package.Title,
				Description: i.Package.Description,
			}
		}

		a := returnBooking{
			ID: i.ID,
			User: returnUser{
				ID:       i.UserID,
				Username: i.User.Username,
			},
			Destination: returnDestination{
				ID:   i.DestinationID,
				Name: i.Destination.Name,
			},
			Package: *p,
		}
		returnBookings = append(returnBookings, a)
	}

	c.Status(fiber.StatusOK).JSON(returnBookings)

	return nil
}

// GET /booking/:id
// require auth
// require database
// TODO: less duplicate code
func HandleBooking(c *fiber.Ctx) error {
	DB := c.Locals("database").(*gorm.DB)
	user := c.Locals("user").(model.User)

	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "invalid booking id")
	}

	var booking model.Booking

	if err := DB.Where("user_id = ?", user.ID).Preload("User").Preload("Destination").Preload("Packages").Limit(1).Find(&booking, id).Error; err != nil {
		return err
	}
	if booking.ID == 0 {
		return fiber.NewError(fiber.StatusNotFound, "booking not found")
	}

	type returnUser struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
	}

	type returnDestination struct {
		ID   uint   `json:"id"`
		Name string `json:"username"`
	}

	type returnPackage struct {
		ID          uint   `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	type returnBooking struct {
		ID          uint              `json:"id"`
		User        returnUser        `json:"user"`
		Destination returnDestination `json:"destination"`
		Package     returnPackage     `json:"package"`
		On          time.Time         `json:"on"`
	}

	var p *returnPackage
	if booking.Package != nil {
		p = &returnPackage{
			ID:          *booking.PackageID,
			Title:       booking.Package.Title,
			Description: booking.Package.Description,
		}
	}

	c.Status(fiber.StatusOK).JSON(returnBooking{
		ID: booking.ID,
		User: returnUser{
			ID:       booking.UserID,
			Username: booking.User.Username,
		},
		Destination: returnDestination{
			ID:   booking.DestinationID,
			Name: booking.Destination.Name,
		},
		On:      booking.On,
		Package: *p,
	})

	return nil
}

// DELETE /booking/:id
// require auth
// require database
func HandleCancel(c *fiber.Ctx) error {
	DB := c.Locals("database").(*gorm.DB)
	user := c.Locals("user").(model.User)

	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "invalid booking id")
	}

	var booking model.Booking

	if err := DB.Where("user_id = ?", user.ID).Preload("User").Preload("Destination").Preload("Packages").Limit(1).Find(&booking, id).Error; err != nil {
		return err
	}
	if booking.ID == 0 {
		return fiber.NewError(fiber.StatusNotFound, "booking not found")
	}

	if err := DB.Delete(&booking).Error; err != nil {
		return err
	}

	type returnUser struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
	}

	type returnDestination struct {
		ID   uint   `json:"id"`
		Name string `json:"username"`
	}

	type returnPackage struct {
		ID          uint   `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	type returnBooking struct {
		ID          uint              `json:"id"`
		User        returnUser        `json:"user"`
		Destination returnDestination `json:"destination"`
		Package     returnPackage     `json:"package"`
		On          time.Time         `json:"on"`
	}

	var p *returnPackage
	if booking.Package != nil {
		p = &returnPackage{
			ID:          *booking.PackageID,
			Title:       booking.Package.Title,
			Description: booking.Package.Description,
		}
	}

	c.Status(fiber.StatusOK).JSON(returnBooking{
		ID: booking.ID,
		User: returnUser{
			ID:       booking.UserID,
			Username: booking.User.Username,
		},
		Destination: returnDestination{
			ID:   booking.DestinationID,
			Name: booking.Destination.Name,
		},
		On:      booking.On,
		Package: *p,
	})

	return nil
}
