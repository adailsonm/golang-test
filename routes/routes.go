package Routes

import (
	Handler "golang-test/handler"
	"os"

	_ "golang-test/docs"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"gorm.io/gorm"
)

func Initialize(app *fiber.App, database *gorm.DB) {
	UserHandler := Handler.NewUserHandler(database)
	AuthHandler := Handler.NewAuthHandler(database)
	WalletHandler := Handler.NewWalletHandler(database)
	GameHandler := Handler.NewGameHandler(database)

	api := app.Group("/api")
	app.Get("/swagger/*", swagger.HandlerDefault)
	api.Post("/register", UserHandler.CreateUser)
	api.Post("/login", AuthHandler.Login)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("SECRET_JWT_KEY"))},
	}))
	api.Get("/profile", UserHandler.GetUser)
	api.Post("/wallet/deposit", WalletHandler.Deposit)
	api.Post("/wallet/withdraw", WalletHandler.Withdraw)

	api.Post("/slot/spin", GameHandler.Spin)
	api.Get("/slot/history", GameHandler.GetHistory)

}
