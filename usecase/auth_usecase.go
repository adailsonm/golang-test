package Usecase

import (
	"fmt"
	Infra "golang-test/infra"
	Models "golang-test/models"
	"golang-test/utils"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthUseCase struct {
	IUserRepository Models.IUserRepository
}

func NewAuthUseCase(db *gorm.DB) *AuthUseCase {
	return &AuthUseCase{
		IUserRepository: Infra.NewIUserRepository(db),
	}
}

func (u AuthUseCase) Login(request *Models.Auth) (string, error) {
	user, err := u.IUserRepository.FetchUserByEmail(request.Email)
	if err != nil {
		return "", err
	}

	token, err := utils.CreateToken(user)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (u AuthUseCase) verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return os.Getenv("SECRET_JWT_KEY"), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid verify token")
	}

	return nil
}
