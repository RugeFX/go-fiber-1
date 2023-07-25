package router

import (
	"github.com/RugeFX/go-fiber-1.git/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	routes.RegisterUserRoute(api)
}
