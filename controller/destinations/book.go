package destinations

import (
	"card2go_service/database"
	"card2go_service/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func HandleBook(c *fiber.Ctx) error {
	DB, err := database.GetConnection()
	if err != nil {
		return err
	}

	stored := c.Locals("user")
	user, ok := stored.(model.User)
	if !ok {
		return fiber.NewError(fiber.ErrInternalServerError.Code, "failed to get user")
	}

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
		DB.Limit(1).Find(&pkg, pid)

		if pkg.ID == 0 {
			return fiber.NewError(fiber.ErrBadRequest.Code, "package id not found")
		}
	}

	err = DB.Transaction(func(tx *gorm.DB) error {
		booking := model.Booking{
			// On: ...
			Destination: dest,
			User:        user,
		}

		// if pid != 0 {
		// 	booking.PackageID = pkg.ID
		// }

		// user.Bookings = append(user.Bookings, booking)
		// tx.Save(user)
		// dest.Bookings = append(dest.Bookings, booking)
		// tx.Save(&dest)
		// if err := tx.Model(&user).Association("Bookings").Append(booking); err != nil {
		// 	return err
		// }
		if pid != 0 {
			booking.Package = &pkg
			// pkg.Bookings = append(dest.Bookings, booking)
			// tx.Save(&pkg)
			// if err := tx.First(&pkg).Association("Bookings").Append(&booking); err != nil {
			// 	return err
			// }
		}
		tx.Create(&booking)

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
