package routes

import (
	userHandlers "github.com/RugeFX/go-fiber-1.git/internal/handlers/user"
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoute(r fiber.Router) {
	user := r.Group("/users")
	user.Get("/", userHandlers.GetAllUsers)
	user.Get("/:id", userHandlers.GetUserByID)
	user.Post("/", userHandlers.CreateUser)
}
