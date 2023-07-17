package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(&fiber.Map{
			"hello": &fiber.Map{"bjir": 1},
		})
	})

	app.Use("/test", func(c *fiber.Ctx) error {
		if c.Method() == "GET" {
			return c.JSON(&fiber.Map{
				"rightio": true,
			})
		}

		return c.Status(http.StatusMethodNotAllowed).JSON(&fiber.Map{
			"message": fmt.Sprintf("Method %q is not allowed", c.Method()),
		})
	})

	log.Fatal(app.Listen(":8000"))
}
