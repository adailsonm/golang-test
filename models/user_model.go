package Models

import (
	Common "golang-test/models/common"
)

type User struct {
	Balance float64 `json:"balance,omitempty"`
	UserTable
}

type UserWithoutBalance struct {
	UserTable
}

type UserTable struct {
	Common.Identify
	FirstName string `json:"first_name" gorm:"first_name"`
	LastName  string `json:"last_name" gorm:"last_name"`
	Password  string `json:"-" gorm:"password"`
	Email     string `json:"email" gorm:"email,unique"`
	Common.Times
}
type SingleUserInput struct {
	ID string `json:"id" param:"id" validate:"required"`
}

type IUserUseCase interface {
	Create(request *User) error
	GetUser(request *SingleUserInput) (User, error)
}

type IUserRepository interface {
	Common.Repository
	CreateUser(request *User) error
	FetchUserById(userId string) (User, error)
	FetchUserByEmail(email string) (User, error)
}
