package main

import (
	"github.com/NicholasHellmers/Instagram-Go-API/api/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Start the server
	database.Connect()

	app := fiber.New()

	setupRoutes(app)

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World 👋!")
	// })

	app.Listen(":3000")
}
