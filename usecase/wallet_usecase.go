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

func (w WalletUseCase) GetWallet(identity string) ([]Models.Wallet, error) {
	wallets, err := w.IWalletRepository.FetchWalletByUserId(identity)
	if err != nil {
		return []Models.Wallet{}, err
	}
	return wallets, nil
}

func (w WalletUseCase) Deposit(identity string, request *Models.Wallet) error {
	tx, txErr := w.IWalletRepository.TxStart()
	if txErr != nil {
		return txErr
	}

	wallet := &Models.Wallet{
		WalletTable: Models.WalletTable{
			Amount:      request.Amount,
			UserId:      string(identity),
			Transaction: "Deposit",
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
	tx, txErr := w.IWalletRepository.TxStart()
	if txErr != nil {
		return txErr
	}

	walletUpdate := &Models.Wallet{
		WalletTable: Models.WalletTable{
			Amount:      -request.Amount,
			Transaction: "Withdraw",
			UserId:      string(identity),
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
