package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phatthakarn-ji/website-sec-api/database"
	"github.com/phatthakarn-ji/website-sec-api/models"
)

func GetContact(c *fiber.Ctx) error {
	db := database.DB
	var contact models.Contact
	db.Find(&contact, 1)
	return c.JSON(contact)
}

func UpdateContact(c *fiber.Ctx) error {
	c.Accepts("application/json")
	contact := new(models.Contact)
	database.DB.Find(&contact, 1)
	if contact.Company == "" {
		response := &models.Response{Status: "error", Message: "Data not found", Data: ""}
		return c.Status(404).JSON(&response)
	}

	if err := c.BodyParser(&contact); err != nil {
		response := &models.Response{Status: "error", Message: "Data not found", Data: err.Error()}
		return c.Status(400).JSON(&response)
	}

	database.DB.Save(&contact)
	return c.JSON(&contact)
}
