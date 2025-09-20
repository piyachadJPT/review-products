package routers

import (
	"review-products/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	app.Post("/api/auth/login", controllers.LoginHandler)
	app.Post("/api/auth/register", controllers.RegisterHandler)
}
