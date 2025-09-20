package routers

import (
	"review-products/controllers"

	"github.com/gofiber/fiber/v2"
)

func ReviewRouters(app *fiber.App) {
	app.Get("/api/all-reviews", controllers.GetAllReviews)
	app.Get("/api/review", controllers.GetReviewByProductId)
	app.Post("/api/add-review", controllers.CreateReview)
	app.Patch("/api/update-review", controllers.UpdateReview)
	app.Delete("/api/delete-review", controllers.DeleteReview)
}
