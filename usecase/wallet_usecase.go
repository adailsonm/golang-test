package Usecase

import (
	Infra "golang-test/infra"
	Models "golang-test/models"
	Common "golang-test/models/common"
	"time"

	"gorm.io/gorm"
)

type WalletUseCase struct {
	IWalletRepository Models.IWalletRepository
	IUserRepository   Models.IUserRepository
}

func NewWalletUseCase(db *gorm.DB) *WalletUseCase {
	return &WalletUseCase{
		IWalletRepository: Infra.NewIWalletRepository(db),
		IUserRepository:   Infra.NewIUserRepository(db),
	}
}

func (w WalletUseCase) Deposit(identity string, request *Models.Wallet) error {
	user, err := w.IUserRepository.FetchUserById(identity)
	if err != nil {
		return err
	}
	tx, txErr := w.IWalletRepository.TxStart()
	if txErr != nil {
		return txErr
	}

	wallet := &Models.Wallet{
		WalletTable: Models.WalletTable{
			Amount:      request.Amount,
			UserId:      string(identity),
			Transaction: "Deposit",
			User:        user,
			Times:       Common.Times{CreatedAt: time.Now(), UpdatedAt: time.Now()},
		},
	}

	if err := w.IWalletRepository.Deposit(wallet); err != nil {
		w.IWalletRepository.TxRollback(tx)
		return err
	}

	if err := w.IWalletRepository.TxCommit(tx); err != nil {
		w.IWalletRepository.TxRollback(tx)
		return err
	}
	return nil
}

func (w WalletUseCase) Withdraw(identity string, request *Models.Wallet) error {
	user, err := w.IUserRepository.FetchUserById(identity)
	if err != nil {
		return err
	}
	tx, txErr := w.IWalletRepository.TxStart()
	if txErr != nil {
		return txErr
	}

	walletUpdate := &Models.Wallet{
		WalletTable: Models.WalletTable{
			Amount:      request.Amount,
			Transaction: "Withdraw",
			UserId:      user.ID.String(),
			User:        user,
			Times:       Common.Times{CreatedAt: time.Now(), UpdatedAt: time.Now()},
		},
	}

	if err := w.IWalletRepository.Withdraw(walletUpdate); err != nil {
		w.IWalletRepository.TxRollback(tx)
		return err
	}

	if err := w.IWalletRepository.TxCommit(tx); err != nil {
		w.IWalletRepository.TxRollback(tx)
		return err
	}
	return nil
}