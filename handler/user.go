package Handler

import (
	Models "golang-test/models"
	Usecase "golang-test/usecase"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserHandler struct {
	IUserUseCase Models.IUserUseCase
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		IUserUseCase: Usecase.NewUserUseCase(db),
	}
}

func (u UserHandler) GetUser(c *fiber.Ctx) error {
	userJwt := c.Locals("user").(*jwt.Token)
	claims := userJwt.Claims.(jwt.MapClaims)
	identity, _ := uuid.Parse(claims["id"].(string))

	user, err := u.IUserUseCase.GetUser(&Models.SingleUserInput{ID: identity.String()})
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
		"data":   user,
	})
}

func (u UserHandler) CreateUser(c *fiber.Ctx) error {
	request := &Models.User{}
	if err := c.BodyParser(request); err != nil {
		log.Fatal(err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "error",
			"message": fiber.ErrBadRequest.Message,
		})
	}

	err := u.IUserUseCase.Create(request)
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
		"status":  "success",
		"message": "Create Successfully!",
	})
}