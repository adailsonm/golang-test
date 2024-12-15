package Models

import (
	Common "golang-test/models/common"
)

type Wallet struct {
	WalletTable
}

type WalletTable struct {
	Common.Identify
	Amount      float64 `json:"amount" gorm:"amount"`
	Transaction string  `json:"transaction" gorm:"transaction"`
	UserId      string
	User        User
	Common.Times
}

type IWalletUseCase interface {
	Deposit(identity string, request *Wallet) error
	Withdraw(identity string, request *Wallet) error
	CreateTransaction(identity string, request *Wallet) error
	GetWallet(identity string) ([]Wallet, error)
}

type IWalletRepository interface {
	Common.Repository
	Deposit(request *Wallet) error
	Withdraw(request *Wallet) error
	CreateTransaction(request *Wallet) error
	FetchWalletByUserId(userId string) ([]Wallet, error)
}
