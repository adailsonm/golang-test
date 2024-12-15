package Infra

import (
	Models "golang-test/models"

	"gorm.io/gorm"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewIWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{
		db: db,
	}
}

func (w WalletRepository) TxStart() (*gorm.DB, error) {
	tx := w.db.Begin()
	return tx, tx.Error
}

func (w WalletRepository) TxCommit(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (w WalletRepository) TxRollback(tx *gorm.DB) {
	tx.Rollback()
}

func (w WalletRepository) Deposit(request *Models.Wallet) error {
	return w.db.Create(request).Error
}

func (w WalletRepository) Withdraw(request *Models.Wallet) error {
	return w.db.Create(request).Error
}

func (w WalletRepository) FetchWalletByUserId(userId string) ([]Models.Wallet, error) {
	results := []Models.Wallet{}
	err := w.db.
		Table("wallets").
		Where("user_id = ?", userId).
		Find(&results).Error
	return results, err
}
