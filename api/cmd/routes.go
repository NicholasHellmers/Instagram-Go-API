package main

import (
	"github.com/NicholasHellmers/Instagram-Go-API/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)

	app.Post("/fact", handlers.CreateMessage)

	app.Get("/facts", handlers.GetMessages)
}
