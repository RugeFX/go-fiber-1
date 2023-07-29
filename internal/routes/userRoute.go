package routes

import (
	userHandlers "github.com/RugeFX/go-fiber-1.git/internal/handlers/user"
	"github.com/gofiber/fiber/v2"
)

// Registers the /users endpoint into the router
func RegisterUserRoute(r fiber.Router) {
	user := r.Group("/users")
	user.Get("/", userHandlers.GetAllUsers)
	user.Get("/:username", userHandlers.GetUserByUsername)
	user.Post("/", userHandlers.CreateUser)
	user.Delete("/:id", userHandlers.DeleteUserByID)
	user.Put("/:id", userHandlers.UpdateUser)
}
