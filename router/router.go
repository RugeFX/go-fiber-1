package router

import (
	routes "github.com/RugeFX/go-fiber-1.git/routes/user"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	routes.RegisterUserRoute(api)
}
