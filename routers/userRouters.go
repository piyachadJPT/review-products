package routers

import (
	"review-products/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app *fiber.App) {
	app.Get("/api/user/:id", controllers.GetUser)
}
