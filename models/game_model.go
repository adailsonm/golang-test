package Models

import (
	Common "golang-test/models/common"
)

type Game struct {
	GameTable
}

type GameTable struct {
	Common.Identify
	BetAmount float64 `json:"bet_amount" gorm:"bet_amount"`
	Result    string  `json:"result" gorm:"result"`
	Payout    float64 `json:"payout" gorm:"payout"`
	Status    string  `json:"status" gorm:"status"`
	UserId    string  `gorm:"column:user_id"`
	User      User    `json:"user" gorm:"foreignKey:UserId;references:ID"`
	Common.Times
}

type BetResult struct {
	Result  string  `json:"results"`
	Payout  float64 `json:"payout"`
	Balance float64 `json:"balance"`
	Status  string  `json:"status"`
}
type IGameUseCase interface {
	Spin(identity string, request *Game) (*BetResult, error)
	GetHistory(identity string) ([]Game, error)
}

type IGameRepository interface {
	Common.Repository
	CreateBet(request *Game) error
	FindHistorical(identity string) ([]Game, error)
}
