package routers

import (
	"review-products/controllers"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App) {
	app.Get("/api/all-product", controllers.GetAllProducts)
	app.Get("/api/product", controllers.GetProductById)
	app.Post("/api/product/create", controllers.CreateProduct)
	app.Patch("/api/product/update", controllers.UpdateProduct)
	app.Delete("/api/product/delete", controllers.DeleteProduct)
}
