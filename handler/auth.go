package Handler

import (
	Models "golang-test/models"
	Usecase "golang-test/usecase"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthHandler struct {
	IAuthUseCase Models.IAuthUseCase
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{
		IAuthUseCase: Usecase.NewAuthUseCase(db),
	}
}

func (a AuthHandler) Login(c *fiber.Ctx) error {
	request := &Models.Auth{}
	if err := c.BodyParser(request); err != nil {
		log.Fatal(err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "error",
			"message": fiber.ErrBadRequest.Message,
		})
	}
	tokenString, err := a.IAuthUseCase.Login(request)

	if err != nil {
		data := map[string]interface{}{
			"error": err.Error(),
		}
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":  "error",
			"message": data,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"token": tokenString,
		},
	})
}