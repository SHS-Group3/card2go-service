package admin

import (
	"card2go_service/database"
	"card2go_service/model"
	"fmt"
	"math"

	"github.com/gofiber/fiber/v2"
)

func HandleClearDestinations(c *fiber.Ctx) error {
	DB, err := database.GetConnection()
	if err != nil {
		return err
	}

	DB.Delete(&model.Destination{})

	c.Status(fiber.StatusOK)
	return nil
}

func HandleCreateDummyDestinations(c *fiber.Ctx) error {
	DB, err := database.GetConnection()
	if err != nil {
		return err
	}

	var dests []model.Destination

	for i := range 40 {
		dests = append(dests, model.Destination{
			Name:        fmt.Sprintf("dummy%d", i),
			Description: fmt.Sprintf("entry no. %d", i),
			Address:     "your house",
			IsLodging:   i%2 == 1,
			Rooms:       i,
			Beds:        i * 2,
			Ratings:     math.Mod(float64(i)*1.3, 5.0),
			Packages: []model.Package{
				model.Package{
					Title:       "your organs",
					Description: "we receive your organs",
				},
				model.Package{
					Title:       "your soul",
					Description: "we receive your soul",
				},
			},
		})
	}
	DB.Create(&dests)

	c.Status(fiber.StatusOK)
	return nil
}
