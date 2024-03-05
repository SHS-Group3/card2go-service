package main

import (
	"card2go_service/config"
	"card2go_service/database"
	"card2go_service/model"
	"card2go_service/routes"
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

			// Return status code with error message
			return c.Status(code).JSON(fiber.Map{
				"code":  code,
				"error": err.Error(),
			})
		},
	})

	app.Use(cors.New())

	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "[${ip}] ${locals:requestid} ${status} - ${method} ${path}\n",
	}))

	app.Use(recover.New())

	if err := godotenv.Load(); err != nil {
		fmt.Errorf("Error loading .env file ")
	}

	config.LoadFromEnv()

	setupDB()

	routes.RegisterAPI(app)

	log.Fatal(app.Listen(":8080"))
}

func setupDB() {
	DB, err := database.GetConnection()
	if err != nil {
		log.Fatal("Failed to connect to database! ", err.Error())
	}

	DB.AutoMigrate(&model.User{}, &model.POI{}, &model.Booking{}, &model.Package{})
}
