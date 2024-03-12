package destinations

import (
	auth "card2go_service/authentication"
	"card2go_service/model"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GET /destination/:id
// requires database
func HandleDestination(c *fiber.Ctx) error {
	DB := c.Locals("database").(*gorm.DB)

	id, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "invalid id")
	}

	var dest model.Destination
	DB.Model(&dest).Where("id = ?", id).Preload("Packages").Limit(1).Find(&dest)

	if dest.ID == 0 {
		return fiber.NewError(fiber.ErrBadRequest.Code, "destination id not found")
	}

	booked := false

	authorization := c.Get("Authorization")
	if authorization != "" {
		if str := strings.Split(authorization, " "); str[0] == "Bearer" {
			id, err := auth.GetIDFromToken(str[1])
			if err != nil {
				return fiber.NewError(fiber.ErrBadRequest.Code, fmt.Sprintf("error while parsing token %s", err.Error()))
			}

			var count int64
			DB.Model(&model.Booking{}).Where("destination_id = ? AND user_id = ?", dest.ID, id).Count(&count)

			booked = count < 0
		}
	}

	// response representations
	// TODO: unify these
	type returnPackage struct {
		ID          uint   `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	type returnDestination struct {
		ID          uint   `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Address     string `json:"address"`
		IsLodging   bool   `json:"is_lodging"`
		Booked      bool   `json:"booked"`
		Beds        int    `json:"beds"`
		Rooms       int    `json:"rooms"`

		Packages []returnPackage `json:"packages"`
	}

	a := returnDestination{
		ID:          dest.ID,
		Name:        dest.Name,
		Address:     dest.Address,
		Description: dest.Description,
		Booked:      booked,
	}

	for _, i := range dest.Packages {
		p := returnPackage{
			ID:          i.ID,
			Title:       i.Title,
			Description: i.Description,
		}

		a.Packages = append(a.Packages, p)
	}

	c.Status(fiber.StatusOK).JSON(a)

	return nil
}

// GET /destinations
// requires database
func HandleFeed(c *fiber.Ctx) error {
	DB := c.Locals("database").(*gorm.DB)

	offset := (c.QueryInt("page", 1) - 1) * 20
	hotels := c.QueryBool("hotels")

	if offset < 0 {
		return fiber.NewError(fiber.ErrBadRequest.Code, "invalid page")
	}

	var dests []model.Destination
	err := DB.Model(&model.Destination{}).Where("is_lodging = ?", hotels).Preload("Packages").Order("created_at desc").Offset(offset).Limit(20).Find(&dests).Error
	if err != nil {
		return err
	}

	// response representations
	type returnPackage struct {
		ID          uint   `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	type returnDestination struct {
		ID          uint   `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Address     string `json:"address"`
		IsLodging   bool   `json:"is_lodging"`
		Beds        int    `json:"beds"`
		Rooms       int    `json:"rooms"`

		Packages []returnPackage `json:"packages"`
	}

	var returnDests []returnDestination

	for _, dest := range dests {
		a := returnDestination{
			ID:          dest.ID,
			Name:        dest.Name,
			Address:     dest.Address,
			Description: dest.Description,
			IsLodging:   dest.IsLodging,
			Beds:        dest.Beds,
			Rooms:       dest.Rooms,
		}

		for _, j := range dest.Packages {
			p := returnPackage{
				ID:          j.ID,
				Title:       j.Title,
				Description: j.Description,
			}

			a.Packages = append(a.Packages, p)
		}

		returnDests = append(returnDests, a)
	}

	c.JSON(returnDests)

	return nil
}

// POST /destination/:id/book
// POST /destination/:id/book/:pid
// requires database
// requires authorization
func HandleBook(c *fiber.Ctx) error {
	DB := c.Locals("database").(*gorm.DB)
	user := c.Locals("user").(model.User)

	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "invalid destination id")
	}

	var dest model.Destination
	DB.Model(&dest).Where("id = ?", id).Preload("Packages").Limit(1).Find(&dest)

	if dest.ID == 0 {
		return fiber.NewError(fiber.ErrBadRequest.Code, "destination id not found")
	}

	pid, err := c.ParamsInt("pid", 0)
	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "invalid package id")
	}

	// look up pkg before making creation query
	var pkg model.Package
	if pid != 0 {
		for _, i := range dest.Packages {
			if i.ID == uint(pid) {
				pkg = i
				break
			}
		}

		if pkg.ID == 0 {
			return fiber.NewError(fiber.ErrBadRequest.Code, "package id not found")
		}
	}

	booking := model.Booking{
		// On: ...
		Destination: dest,
		User:        user,
	}
	if pid != 0 {
		booking.Package = &pkg
	}

	err = DB.Transaction(func(tx *gorm.DB) error {
		tx.Create(&booking)

		return nil
	})
	if err != nil {
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

	c.Status(fiber.StatusCreated).JSON(returnBooking{
		ID: booking.ID,
		User: returnUser{
			ID:       booking.UserID,
			Username: booking.User.Username,
		},
		Destination: returnDestination{
			ID:   booking.DestinationID,
			Name: booking.Destination.Name,
		},
		Package: *p,
	})

	return nil
}
