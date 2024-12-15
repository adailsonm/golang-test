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

type WalletHandler struct {
	IWalletUseCase Models.IWalletUseCase
}

func NewWalletHandler(db *gorm.DB) *WalletHandler {
	return &WalletHandler{
		IWalletUseCase: Usecase.NewWalletUseCase(db),
	}
}

// Deposit godoc
// @Summary Deposit money into wallet
// @Description Deposit a specified amount into the wallet
// @Tags Wallet
// @Accept json
// @Produce json
// @Param amount body float64 true "Deposit Amount"
// @Security BearerAuth
// @Success 200 {string} string "Deposit successful"
// @Router /wallet/deposit [post]
func (w WalletHandler) Deposit(c *fiber.Ctx) error {
	userJwt := c.Locals("user").(*jwt.Token)
	claims := userJwt.Claims.(jwt.MapClaims)
	identity, _ := uuid.Parse(claims["id"].(string))
	request := &Models.Wallet{}
	if err := c.BodyParser(request); err != nil {
		log.Fatal(err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "error",
			"message": fiber.ErrBadRequest.Message,
		})
	}
	err := w.IWalletUseCase.Deposit(identity.String(), request)
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
		"data":   "Deposit with Successfully!",
	})
}

// Withdraw godoc
// @Summary Withdraw money from wallet
// @Description Withdraw a specified amount from the wallet
// @Tags Wallet
// @Accept json
// @Produce json
// @Param amount body float64 true "Withdraw Amount"
// @Security BearerAuth
// @Success 200 {string} string "Withdraw successful"
// @Router /wallet/withdraw [post]
func (w WalletHandler) Withdraw(c *fiber.Ctx) error {
	userJwt := c.Locals("user").(*jwt.Token)
	claims := userJwt.Claims.(jwt.MapClaims)
	identity, _ := uuid.Parse(claims["id"].(string))
	request := &Models.Wallet{}

	if err := c.BodyParser(request); err != nil {
		log.Fatal(err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":  "error",
			"message": fiber.ErrBadRequest.Message,
		})
	}
	err := w.IWalletUseCase.Withdraw(identity.String(), request)
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
		"message": "Withdraw with Successfully!",
	})
}
