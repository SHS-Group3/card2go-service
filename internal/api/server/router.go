package server

import "github.com/gofiber/fiber/v2"

func RegisterAPI(app *fiber.App) {
	RegisterAdmins(app)
	RegisterTours(app)
	RegisterHotels(app)
	RegisterPOI(app)
}

func RegisterAdmins(app *fiber.App) {

}

func RegisterTours(app *fiber.App) {

}

func RegisterHotels(app *fiber.App) {

}

func RegisterPOI(app *fiber.App) {

}
