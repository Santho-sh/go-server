package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// Define a simple route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! Go app is running!")
	})

	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// Start the server
	port := ":4000"
	if err := app.Listen(port); err != nil {
		panic(err)
	}
}
