package controllers

import (
	"review-products/database"
	"review-products/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateProduct(c *fiber.Ctx) error {
	type Input struct {
		SKU         string  `json:"sku"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float32 `json:"price"`
		Stock       int     `json:"stock"`
	}

	var input Input
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON body",
			"error":   err.Error(),
		})
	}

	fields := map[string]string{
		"SKU":   input.SKU,
		"Name":  input.Name,
		"Price": strconv.FormatFloat(float64(input.Price), 'f', 2, 32),
	}

	for key, value := range fields {
		if value == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": key + " is required",
			})
		}
	}

	var product = models.Product{
		SKU:         &input.SKU,
		Name:        input.Name,
		Description: &input.Description,
		Price:       float64(input.Price),
		Stock:       input.Stock,
	}

	if err := database.DB.Create(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create product",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"ok":      true,
		"message": "Product created successfully",
		"product": product,
	})
}

func GetAllProducts(c *fiber.Ctx) error {
	var products []models.Product

	result := database.DB.Preload("Images").Find(&products)
	if result.Error != nil {
		return c.Status(404).SendString("Product not found")
	}

	if err := database.DB.
		Preload("Images").
		Preload("Review").
		Preload("Review.User").
		Find(&products).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Product not found",
		})
	}

	return c.JSON(fiber.Map{
		"ok":       true,
		"products": products,
	})
}

func GetProductById(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Product ID is required",
		})
	}

	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Invalid product ID format",
		})
	}

	var product models.Product

	if err := database.DB.
		Preload("Images").
		Preload("Review").
		Preload("Review.User").
		First(&product, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Product not found",
		})
	}

	return c.JSON(fiber.Map{
		"ok":      true,
		"product": product,
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Product ID is required",
		})
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Invalid Product ID format",
		})
	}

	var product models.Product
	if err := database.DB.First(&product, "id = ?", uid).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Product not found",
		})
	}

	type Input struct {
		SKU         string  `json:"sku"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float32 `json:"price"`
		Stock       int     `json:"stock"`
	}

	var input Input
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Invalid JSON body",
		})
	}

	fields := map[string]string{
		"SKU":   input.SKU,
		"Name":  input.Name,
		"Price": strconv.FormatFloat(float64(input.Price), 'f', 2, 32),
	}

	for key, value := range fields {
		if value == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": key + " is required",
			})
		}
	}

	if err := database.DB.Save(&product).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": "Failed to update product",
		})
	}

	return c.JSON(fiber.Map{
		"ok":      true,
		"message": "Product updated successfully",
		"product": product,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Product ID is required",
		})
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Invalid product ID format",
		})
	}

	var product models.Product
	if err := database.DB.First(&product, "id = ?", uid).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Product not found",
		})
	}

	if err := database.DB.Delete(&product).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": "Failed to delete Product",
		})
	}

	return c.JSON(fiber.Map{
		"ok":      true,
		"message": "Product deleted successfully",
	})
}
