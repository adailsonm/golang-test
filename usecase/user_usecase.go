package Usecase

import (
	"errors"
	"fmt"
	Infra "golang-test/infra"
	Models "golang-test/models"
	Common "golang-test/models/common"
	"golang-test/utils"
	"time"

	"gorm.io/gorm"
)

type UserUseCase struct {
	IUserRepository Models.IUserRepository
}

func NewUserUseCase(db *gorm.DB) *UserUseCase {
	return &UserUseCase{
		IUserRepository: Infra.NewIUserRepository(db),
	}
}

func (u UserUseCase) GetUser(request *Models.SingleUserInput) (Models.User, error) {
	user, err := u.IUserRepository.FetchUserById(request.ID)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u UserUseCase) Create(request *Models.User) error {
	tx, txErr := u.IUserRepository.TxStart()
	if txErr != nil {
		return txErr
	}

	_, err := u.IUserRepository.FetchUserByEmail(request.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		u.IUserRepository.TxRollback(tx)
		return fmt.Errorf("error fetching user by email: %v", err)
	}

	if err == nil {
		u.IUserRepository.TxRollback(tx)
		return fmt.Errorf("user with email %s already exists", request.Email)
	}

	passwordHash, err := utils.HashPassword(request.Password)
	if err != nil {
		u.IUserRepository.TxRollback(tx)
		return fmt.Errorf("error in create hashpassword: %v", err)
	}

	user := &Models.User{
		UserTable: Models.UserTable{
			FirstName: request.FirstName,
			LastName:  request.LastName,
			Email:     request.Email,
			Password:  passwordHash,
			Times:     Common.Times{CreatedAt: time.Now(), UpdatedAt: time.Now()},
		},
	}

	if err := u.IUserRepository.CreateUser(user); err != nil {
		u.IUserRepository.TxRollback(tx)
		return err
	}

	if err := u.IUserRepository.TxCommit(tx); err != nil {
		u.IUserRepository.TxRollback(tx)
		return err
	}
	return nil
}
