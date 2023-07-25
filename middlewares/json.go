package middlewares

import "github.com/gofiber/fiber/v2"

func Json() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return c.Next()
	}
}
