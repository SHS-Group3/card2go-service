package routes

import (
	"card2go_service/controller/auth"
	"card2go_service/controller/hotels"
	"card2go_service/controller/packages"
	"card2go_service/controller/poi"

	"github.com/gofiber/fiber/v2"
)

func RegisterAPI(app *fiber.App) {
	RegisterAdmins(app)
	RegisterAuth(app)
	RegisterPackages(app)
	RegisterHotels(app)
	RegisterPOI(app)
}

func RegisterAdmins(app *fiber.App) {

}

func RegisterAuth(app *fiber.App) {
	path := app.Group("/auth")
	path.Post("/", auth.HandleAuthentication)
	path.Post("/register", auth.HandleRegister)
}

func RegisterPackages(app *fiber.App) {
	path := app.Group("/packages")
	path.Get("/:id", packages.HandleBooking)

}

func RegisterHotels(app *fiber.App) {
	path := app.Group("/hotels")
	path.Get("/", hotels.HandleFeed)
	path.Get("/:id", hotels.HandleHotel)
}

func RegisterPOI(app *fiber.App) {
	path := app.Group("/poi")
	path.Get("/", poi.HandleFeed)
	path.Get("/:id", poi.HandlePOI)
}
