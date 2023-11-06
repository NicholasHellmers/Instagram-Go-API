package handlers

import (
	"github.com/NicholasHellmers/Instagram-Go-API/api/database"
	"github.com/NicholasHellmers/Instagram-Go-API/api/models"
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func CreateMessage(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	// Send test message

	fact = &models.Fact{
		Message: "This is a test message",
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}

func GetMessages(c *fiber.Ctx) error {
	var facts []models.Fact
	database.DB.Db.Find(&facts)
	return c.Status(200).JSON(facts)
}
