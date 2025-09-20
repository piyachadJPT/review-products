package controllers

import (
	"review-products/database"
	"review-products/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UploadProductImage(c *fiber.Ctx) error {
	type Input struct {
		ProductID string `json:"productID"`
		URL       string `json:"url"`
		Alt       string `json:"alt"`
		Position  int    `json:"position"`
	}

	var input Input
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON body",
			"error":   err.Error(),
		})
	}

	if input.ProductID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ProductID is required",
		})
	}

	// แปลงรูปเป็น Base64
	base64Image, err := loadImageAsBase64(input.URL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to load image",
		})
	}

	productUUID, err := uuid.Parse(input.ProductID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ProductID UUID",
		})
	}

	var lastImage models.ProductImage
	position := 1

	if err := database.DB.
		Where("product_id = ?", productUUID).
		Order("position DESC").
		First(&lastImage).Error; err == nil {
		position = lastImage.Position + 1
	}

	productImage := models.ProductImage{
		ProductID: input.ProductID,
		URL:       base64Image,
		Alt:       &input.Alt,
		Position:  position,
	}

	if err := database.DB.Create(&productImage).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to save product image",
		})
	}

	return c.JSON(fiber.Map{
		"ok":      true,
		"message": "Upload image successfully",
	})
}
