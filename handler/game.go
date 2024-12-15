package Handler

import (
	Models "golang-test/models"
	Usecase "golang-test/usecase"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GameHandler struct {
	IGameUseCase Models.IGameUseCase
}

func NewGameHandler(db *gorm.DB) *GameHandler {
	return &GameHandler{
		IGameUseCase: Usecase.NewGameUseCase(db),
	}
}

func (g GameHandler) Spin(c *fiber.Ctx) error {
	userJwt := c.Locals("user").(*jwt.Token)
	claims := userJwt.Claims.(jwt.MapClaims)
	identity, _ := uuid.Parse(claims["id"].(string))
	request := &Models.Game{}
	if err := c.BodyParser(request); err != nil {
		log.Fatal(err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "error",
			"message": fiber.ErrBadRequest.Message,
		})
	}
	betResult, err := g.IGameUseCase.Spin(identity.String(), request)
	if err != nil {
		data := map[string]interface{}{
			"error": err.Error(),
		}
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":  "error",
			"message": data,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   betResult,
	})
}
