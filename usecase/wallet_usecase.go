package Usecase

import (
	"fmt"
	Infra "golang-test/infra"
	Models "golang-test/models"
	Common "golang-test/models/common"
	"time"

	"github.com/gofiber/fiber/v2"
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

func (w WalletUseCase) Deposit(identity string, request *Models.Wallet) (fiber.Map, error) {
	transactions, err := w.GetWallet(identity)
	if err != nil {
		return nil, err
	}

	var balance float64
	for _, transaction := range transactions {
		balance += transaction.Amount
	}

	tx, txErr := w.IWalletRepository.TxStart()
	if txErr != nil {
		return nil, txErr
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
		return nil, err
	}

	if err := w.IWalletRepository.TxCommit(tx); err != nil {
		w.IWalletRepository.TxRollback(tx)
		return nil, err
	}
	return fiber.Map{
		"balance": balance,
	}, nil
}

func (w WalletUseCase) Withdraw(identity string, request *Models.Wallet) (fiber.Map, error) {
	transactions, err := w.GetWallet(identity)
	if err != nil {
		return nil, err
	}

	var balance float64
	for _, transaction := range transactions {
		balance += transaction.Amount
	}

	if balance < request.Amount {
		return nil, fmt.Errorf("insufficient balance: current balance is %.2f, requested amount is %.2f", balance, request.Amount)
	}

	tx, txErr := w.IWalletRepository.TxStart()
	if txErr != nil {
		return nil, txErr
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
		return nil, err
	}

	if err := w.IWalletRepository.TxCommit(tx); err != nil {
		w.IWalletRepository.TxRollback(tx)
		return nil, err
	}
	return fiber.Map{
		"balance": balance,
	}, nil
}

func (w WalletUseCase) CreateTransaction(identity string, request *Models.Wallet) error {
	tx, txErr := w.IWalletRepository.TxStart()
	if txErr != nil {
		return txErr
	}

	walletUpdate := &Models.Wallet{
		WalletTable: Models.WalletTable{
			Amount:      request.Amount,
			Transaction: "Game",
			UserId:      string(identity),
			Times:       Common.Times{CreatedAt: time.Now(), UpdatedAt: time.Now()},
		},
	}

	if err := w.IWalletRepository.CreateTransaction(walletUpdate); err != nil {
		w.IWalletRepository.TxRollback(tx)
		return err
	}

	if err := w.IWalletRepository.TxCommit(tx); err != nil {
		w.IWalletRepository.TxRollback(tx)
		return err
	}
	return nil
}
