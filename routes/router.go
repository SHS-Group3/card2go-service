package routes

import (
	"card2go_service/controller/admin"
	"card2go_service/controller/auth"
	"card2go_service/controller/bookings"
	"card2go_service/controller/destinations"
	"card2go_service/controller/middleware"
	"card2go_service/controller/user"

	"github.com/gofiber/fiber/v2"
)

func RegisterAPI(app *fiber.App) {
	RegisterAdmin(app)
	RegisterAuth(app)
	RegisterBookings(app)
	RegisterDestinations(app)
	RegisterUser(app)
}

func RegisterUser(app *fiber.App) {
	path := app.Group("/user", middleware.RequireDatabase, middleware.RequireAuth)
	path.Get("/", user.HandleInfo)
}

func RegisterBookings(app *fiber.App) {
	path := app.Group("/bookings", middleware.RequireDatabase, middleware.RequireAuth)
	path.Get("/", bookings.HandleBookings)
	path.Get("/:id", bookings.HandleBooking)
	path.Delete("/:id", bookings.HandleCancel)
}

func RegisterAdmin(app *fiber.App) {
	path := app.Group("/admin", middleware.RequireDatabase)
	path.Get("/destinations/clear", admin.HandleClearDestinations)
	path.Get("/destinations/dummy", admin.HandleCreateDummyDestinations)
}

func RegisterAuth(app *fiber.App) {
	path := app.Group("/auth", middleware.RequireDatabase)
	path.Post("/", auth.HandleAuthentication)
	path.Post("/register", auth.HandleRegister)
}

func RegisterDestinations(app *fiber.App) {
	path := app.Group("/destinations", middleware.RequireDatabase)
	path.Get("/", destinations.HandleFeed)
	path.Get("/:id", destinations.HandleDestination)

	path.Post("/:id/book", middleware.RequireAuth, destinations.HandleBook)
	path.Post("/:id/book/:pid", middleware.RequireAuth, destinations.HandleBook)
}
