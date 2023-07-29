package router

import (
	"github.com/RugeFX/go-fiber-1.git/internal/routes"
	"github.com/gofiber/fiber/v2"
)

// Sets up all of the routes for the API
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	routes.RegisterUserRoute(api)
}
