package routers

import (
	"review-products/controllers"

	"github.com/gofiber/fiber/v2"
)

func ProductImageRoutes(app *fiber.App) {
	app.Post("/api/upload-image-product", controllers.UploadProductImage)
}
