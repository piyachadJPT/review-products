package controllers

import (
	"review-products/database"
	"review-products/models"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	id := c.Params("ID")
	var user models.User

	result := database.DB.First(&user, id)
	if result.Error != nil {
		c.Status(404)
		return c.SendString("User not found")
	}

	return c.JSON(user)
}
