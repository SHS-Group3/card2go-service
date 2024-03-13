package admin

import (
	"card2go_service/model"
	"fmt"
	"math"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GET /admin/destinations/clear
// requires database
func HandleClearDestinations(c *fiber.Ctx) error {
	DB := c.Locals("database").(*gorm.DB)

	DB.Delete(&model.Destination{})

	c.Status(fiber.StatusOK)
	return nil
}

// GET /admin/destinations/dummy
// requires database
func HandleCreateDummyDestinations(c *fiber.Ctx) error {
	DB := c.Locals("database").(*gorm.DB)

	var dests []model.Destination

	for i := range 40 {
		price := float64(i) * 100.
		dests = append(dests, model.Destination{
			Name:        fmt.Sprintf("dummy%d", i),
			Description: fmt.Sprintf("entry no. %d", i),
			Address:     "your house",
			IsLodging:   i%2 == 1,
			Rooms:       i,
			Beds:        i * 2,
			Ratings:     math.Mod(float64(i)*1.3, 5.0),
			Packages: []model.Package{
				{
					Title:       "your organs",
					Description: "we receive your organs",
				},
				{
					Title:       "your soul",
					Description: "we receive your soul",
					Price:       &price,
				},
				{
					Title:       "offered to demons bundle",
					Description: "lose your soul and your organs all in one bundle!",
					Price:       &price,
				},
			},
		})
	}
	DB.Create(&dests)

	c.Status(fiber.StatusOK)
	return nil
}
