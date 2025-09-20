package main

import (
	"log"
	"os"
	"review-products/database"
	"review-products/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	database.Connect()

	routers.AuthRoutes(app)
	routers.UserRouter(app)
	routers.ProductRoutes(app)
	routers.ProductImageRoutes(app)
	routers.ReviewRouters(app)

	// ดึง SERVER_PORT จาก env
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server is running on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
