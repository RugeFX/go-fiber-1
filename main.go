package main

import (
	"log"

	"github.com/RugeFX/go-fiber-1.git/database"
	"github.com/RugeFX/go-fiber-1.git/middlewares"
	"github.com/RugeFX/go-fiber-1.git/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	if envErr := godotenv.Load(".env"); envErr != nil {
		panic("Error loading env")
	}
	database.ConnectDB()

	// Initialize the new Fiber App
	app := fiber.New()
	// Use the middlewares globally
	app.Use(cors.New())
	app.Use(middlewares.Json())
	// Testing the endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"welcome": &fiber.Map{"num": 1},
		})
	})

	// Setup the routes from the router
	router.SetupRoutes(app)

	log.Fatal(app.Listen(":8001"))
}
