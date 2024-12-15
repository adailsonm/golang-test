package Usecase

import (
	"fmt"
	Infra "golang-test/infra"
	Models "golang-test/models"
	Common "golang-test/models/common"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type GameUseCase struct {
	IGameRepository Models.IGameRepository
	IUserRepository Models.IUserRepository
	IUserUseCase    Models.IUserUseCase
	IWalletuseCase  Models.IWalletUseCase
}

func NewGameUseCase(db *gorm.DB) *GameUseCase {
	return &GameUseCase{
		IGameRepository: Infra.NewIGameRepository(db),
		IUserRepository: Infra.NewIUserRepository(db),
		IUserUseCase:    NewUserUseCase(db),
		IWalletuseCase:  NewWalletUseCase(db),
	}
}

func (g GameUseCase) Spin(identity string, request *Models.Game) (*Models.BetResult, error) {
	user, _ := g.IUserUseCase.GetUser(&Models.SingleUserInput{ID: identity})
	if user.Balance < request.BetAmount {
		return nil, fmt.Errorf("insufficient balance")
	}

	tx, txErr := g.IGameRepository.TxStart()
	if txErr != nil {
		return nil, txErr
	}

	result, payout := spinSlotMachine(request.BetAmount)

	game := &Models.Game{
		GameTable: Models.GameTable{
			BetAmount: request.BetAmount,
			Result:    fmt.Sprintf("%d %d %d", result[0], result[1], result[2]),
			Payout:    payout,
			UserId:    string(identity),
			Times:     Common.Times{CreatedAt: time.Now(), UpdatedAt: time.Now()},
		},
	}

	if err := g.IGameRepository.CreateBet(game); err != nil {
		g.IGameRepository.TxRollback(tx)
		return nil, err
	}

	if payout == 0 {
		wallet := &Models.Wallet{
			WalletTable: Models.WalletTable{
				Amount:      -request.BetAmount,
				UserId:      string(identity),
				Transaction: "Bet",
				Times:       Common.Times{CreatedAt: time.Now(), UpdatedAt: time.Now()},
			},
		}
		if err := g.IWalletuseCase.CreateTransaction(user.ID.String(), wallet); err != nil {
			g.IGameRepository.TxRollback(tx)
			return nil, err
		}
	}

	if payout > 0 {
		wallet := &Models.Wallet{
			WalletTable: Models.WalletTable{
				Amount:      payout,
				UserId:      string(identity),
				Transaction: "Bet",
				Times:       Common.Times{CreatedAt: time.Now(), UpdatedAt: time.Now()},
			},
		}
		if err := g.IWalletuseCase.CreateTransaction(user.ID.String(), wallet); err != nil {
			g.IGameRepository.TxRollback(tx)
			return nil, err
		}
	}

	if err := g.IGameRepository.TxCommit(tx); err != nil {
		g.IGameRepository.TxRollback(tx)
		return nil, err
	}
	return &Models.BetResult{
		Result: fmt.Sprintf("%d %d %d", result[0], result[1], result[2]),
		Payout: payout,
		Status: func() string {
			if payout > 0 {
				return "win"
			}
			return "lose"
		}(),
		Balance: user.Balance,
	}, nil
}

func spinSlotMachine(bet float64) ([3]int, float64) {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := [3]int{
		rand.Intn(9) + 1,
		rand.Intn(9) + 1,
		rand.Intn(9) + 1,
	}

	payout := calculatePayout(result, bet)

	return result, payout
}

func calculatePayout(result [3]int, bet float64) float64 {
	if result[0] == result[1] && result[1] == result[2] {
		return bet * 10
	} else if result[0] == result[1] || result[1] == result[2] || result[0] == result[2] {
		return bet * 2
	}
	return 0
}
