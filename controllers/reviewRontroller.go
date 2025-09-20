package controllers

import (
	"review-products/database"
	"review-products/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateReview(c *fiber.Ctx) error {
	type Input struct {
		ProductID string `json:"productID"`
		UserID    string `json:"userID"`
		Title     string `json:"title"`
		Body      string `json:"body"`
		Rating    int    `json:"rating"`
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

	if input.UserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "UserID is required",
		})
	}

	if input.Rating < 1 || input.Rating > 5 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Rating must be between 1 and 5",
		})
	}

	if input.Body == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Review body is required",
		})
	}

	if _, err := uuid.Parse(input.ProductID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ProductID format",
		})
	}

	if _, err := uuid.Parse(input.UserID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UserID format",
		})
	}

	review := models.Review{
		ProductID: input.ProductID,
		UserID:    input.UserID,
		Title:     &input.Title,
		Body:      input.Body,
		Rating:    input.Rating,
	}

	if err := database.DB.Create(&review).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create review",
		})
	}

	return c.JSON(fiber.Map{
		"ok":      true,
		"message": "Review created successfully",
		"review":  review,
	})
}

func GetReviewByProductId(c *fiber.Ctx) error {
	productId := c.Query("productId")
	if productId == "" {
		return c.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "productId is required",
		})
	}

	// ตรวจสอบ UUID format ไหม
	uid, err := uuid.Parse(productId)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "invalid productId format",
		})
	}

	// ตรวจสอบว่ามี product อยู่จริงไหม
	var product models.Product
	if err := database.DB.Where("id = ?", uid).First(&product).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Product not found",
		})
	}

	var reviews []models.Review
	if err := database.DB.Preload("User").Where("product_id = ?", productId).Find(&reviews).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": "Failed to fetch reviews",
		})
	}

	return c.JSON(fiber.Map{
		"ok":         true,
		"reviews":    reviews,
		"count":      len(reviews),
		"product_id": productId,
		"product":    product.Name,
	})
}

func GetAllReviews(c *fiber.Ctx) error {
	var reviews []models.Review
	if err := database.DB.Preload("User").Find(&reviews).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": "Failed to fetch all reviews",
		})
	}

	return c.JSON(fiber.Map{
		"ok":      true,
		"reviews": reviews,
		"count":   len(reviews),
	})
}

func UpdateReview(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Review ID is required",
		})
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Invalid review ID format",
		})
	}

	var review models.Review
	if err := database.DB.First(&review, "id = ?", uid).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Review not found",
		})
	}

	type Input struct {
		Title  *string `json:"title"`
		Body   *string `json:"body"`
		Rating *int    `json:"rating"`
	}

	var input Input
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Invalid JSON body",
		})
	}

	if input.Title != nil {
		review.Title = input.Title
	}
	if input.Body != nil {
		review.Body = *input.Body
	}
	if input.Rating != nil {
		if *input.Rating < 1 || *input.Rating > 5 {
			return c.Status(400).JSON(fiber.Map{
				"ok":    false,
				"error": "Rating must be between 1 and 5",
			})
		}
		review.Rating = *input.Rating
	}

	if err := database.DB.Save(&review).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": "Failed to update review",
		})
	}

	return c.JSON(fiber.Map{
		"ok":      true,
		"message": "Review updated successfully",
		"review":  review,
	})
}

func DeleteReview(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Review ID is required",
		})
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"ok":    false,
			"error": "Invalid review ID format",
		})
	}

	var review models.Review
	if err := database.DB.First(&review, "id = ?", uid).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Review not found",
		})
	}

	if err := database.DB.Delete(&review).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": "Failed to delete review",
		})
	}

	return c.JSON(fiber.Map{
		"ok":      true,
		"message": "Review deleted successfully",
	})
}
