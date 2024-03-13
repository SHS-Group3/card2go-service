package bookings

import (
	"card2go_service/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type returnDestination struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type returnPackage struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type returnBooking struct {
	ID          uint              `json:"id"`
	Destination returnDestination `json:"destination"`
	Package     *returnPackage    `json:"package"`
	On          time.Time         `json:"on"`
}

// GET /bookings
// requires authorization
// requires database
func HandleBookings(c *fiber.Ctx) error {
	DB := c.Locals("database").(*gorm.DB)
	user := c.Locals("user").(model.User)

	var bookings []model.Booking
	if err := DB.Where("user_id = ?", user.ID).Preload("Destination").Preload("Package").Preload("User").Find(&bookings).Error; err != nil {
		return err
	}

	returnBookings := []returnBooking{}
	for _, i := range bookings {
		var p *returnPackage
		if i.Package != nil {
			p = &returnPackage{
				ID:          *i.PackageID,
				Title:       i.Package.Title,
				Description: i.Package.Description,
				Price:       i.Package.Price,
			}
		}

		a := returnBooking{
			ID: i.ID,
			Destination: returnDestination{
				ID:          i.DestinationID,
				Name:        i.Destination.Name,
				Description: i.Destination.Description,
			},
			Package: p,
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

	if err := DB.Where("user_id = ?", user.ID).Preload("User").Preload("Destination").Preload("Package").Limit(1).Find(&booking, id).Error; err != nil {
		return err
	}
	if booking.ID == 0 {
		return fiber.NewError(fiber.StatusNotFound, "booking not found")
	}

	var p *returnPackage
	if booking.Package != nil {
		p = &returnPackage{
			ID:          *booking.PackageID,
			Title:       booking.Package.Title,
			Description: booking.Package.Description,
			Price:       booking.Package.Price,
		}
	}

	c.Status(fiber.StatusOK).JSON(returnBooking{
		ID: booking.ID,
		Destination: returnDestination{
			ID:   booking.DestinationID,
			Name: booking.Destination.Name,
		},
		On:      booking.On,
		Package: p,
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

	if err := DB.Where("user_id = ?", user.ID).Preload("Package").Preload("Destination").Limit(1).Find(&booking, id).Error; err != nil {
		return err
	}
	if booking.ID == 0 {
		return fiber.NewError(fiber.StatusNotFound, "booking not found")
	}

	if err := DB.Delete(&booking).Error; err != nil {
		return err
	}

	var p *returnPackage
	if booking.Package != nil {
		p = &returnPackage{
			ID:          *booking.PackageID,
			Title:       booking.Package.Title,
			Description: booking.Package.Description,
			Price:       booking.Package.Price,
		}
	}

	c.Status(fiber.StatusOK).JSON(returnBooking{
		ID: booking.ID,
		Destination: returnDestination{
			ID:          booking.DestinationID,
			Name:        booking.Destination.Name,
			Description: booking.Destination.Description,
		},
		On:      booking.On,
		Package: p,
	})

	return nil
}
