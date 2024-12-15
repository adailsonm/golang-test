package main

import (
	Database "golang-test/config"
	Models "golang-test/models"
	Routes "golang-test/routes"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/joho/godotenv"
)

// @title NodeArt - Golang
// @version 1.0
// @description Created by Adailson
// @termsOfService http://swagger.io/terms/
// @contact.name Adailson Moreira
// @contact.email adailson.moreira16@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	app := fiber.New()
	app.Use(logger.New())

	app.Use(limiter.New(limiter.Config{
		Max:               20,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

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
	db.AutoMigrate(&Models.Game{})
	serverPort := os.Getenv("SERVER_PORT")
	log.Fatal(app.Listen(":" + serverPort))
}
