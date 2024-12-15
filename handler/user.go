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

// Profile godoc
// @Summary Get user profile
// @Description Get the details of the logged-in user's profile
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {string} string "User profile"
// @Router /profile [get]
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

// Register godoc
// @Summary Register a new user
// @Description Register a new user with their details
// @Tags Auth
// @Accept json
// @Produce json
// @Param first_name body string true "First Name"
// @Param last_name body string true "Last Name"
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 200 {string} string "Registration successful"
// @Router /register [post]
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
