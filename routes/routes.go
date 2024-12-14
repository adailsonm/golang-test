package Routes

import (
	Handler "golang-test/handler"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Initialize(app *fiber.App, database *gorm.DB) {
	UserHandler := Handler.NewUserHandler(database)
	AuthHandler := Handler.NewAuthHandler(database)
	WalletHandler := Handler.NewWalletHandler(database)

	api := app.Group("/api")

	api.Post("/register", UserHandler.CreateUser)
	api.Post("/login", AuthHandler.Login)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("SECRET_JWT_KEY"))},
	}))
	api.Get("/profile", UserHandler.GetUser)
	api.Post("/wallet/deposit", WalletHandler.Deposit)
	api.Post("/wallet/withdraw", WalletHandler.Withdraw)
}
