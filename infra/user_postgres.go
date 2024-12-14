package Infra

import (
	Models "golang-test/models"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewIUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u UserRepository) TxStart() (*gorm.DB, error) {
	tx := u.db.Begin()
	return tx, tx.Error
}

func (u UserRepository) TxCommit(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (u UserRepository) TxRollback(tx *gorm.DB) {
	tx.Rollback()
}

func (u UserRepository) FetchAllUser() ([]Models.User, error) {
	var users []Models.User
	result := u.db.Find(&users)
	if result.Error != nil {
		return users, result.Error
	}

	return users, nil
}

func (u UserRepository) FetchUserById(userId string) (Models.User, error) {
	result := Models.User{}
	err := u.db.
		Table("users").
		Where("id = ?", userId).
		First(&result).Error
	return result, err
}

func (u UserRepository) FetchUserByEmail(email string) (Models.User, error) {
	result := Models.User{}
	err := u.db.
		Table("users").
		Where("email = ?", email).
		First(&result).Error
	return result, err
}

func (u UserRepository) CreateUser(request *Models.User) error {
	return u.db.Create(request).Error
}

func (u UserRepository) UpdateUser(userId string, request *Models.User) error {
	return u.db.Model(&Models.User{}).
		Where("id = ?", userId).
		Updates(map[string]interface{}{
			"first_name": request.FirstName,
			"last_name":  request.LastName,
			"email":      request.Email,
			"category":   request.Password,
			"updated_at": time.Now(),
		}).Error
}

func (u UserRepository) DeleteUser(userId string) error {
	return u.db.Delete(&Models.User{}, "id = ?", userId).Error
}
