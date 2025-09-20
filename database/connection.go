package database

import (
	"fmt"
	"log"
	"os"
	"review-products/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found, using system env")
	}
}

func Connect() {
	// PostgreSQL DSN format
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}
	log.Println("✅ Database connected")

	autoMigrate()
}

func autoMigrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.ProductImage{},
		&models.Review{},
	)
	if err != nil {
		log.Fatalf("❌ Auto migration failed: %v", err)
	}
}
