package main

import (
	Database "golang-test/config"
	Models "golang-test/models"
	Routes "golang-test/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("FRONTEND_URL"),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	Database.Initialize()
	db := Database.GetDB()
	if db == nil {
		log.Fatal("Database connection is not initialized")
	}
	Routes.Initialize(app, db)

	db.AutoMigrate(&Models.User{})
	db.AutoMigrate(&Models.Wallet{})
	serverPort := os.Getenv("SERVER_PORT")
	log.Fatal(app.Listen(":" + serverPort))
}
