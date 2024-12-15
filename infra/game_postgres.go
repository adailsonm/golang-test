package Infra

import (
	Models "golang-test/models"

	"gorm.io/gorm"
)

type GameRepository struct {
	db *gorm.DB
}

func NewIGameRepository(db *gorm.DB) *GameRepository {
	return &GameRepository{
		db: db,
	}
}

func (g GameRepository) TxStart() (*gorm.DB, error) {
	tx := g.db.Begin()
	return tx, tx.Error
}

func (g GameRepository) TxCommit(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (g GameRepository) TxRollback(tx *gorm.DB) {
	tx.Rollback()
}

func (g GameRepository) CreateBet(request *Models.Game) error {
	return g.db.Create(request).Error
}

func (g GameRepository) FindHistorical(identity string) ([]Models.Game, error) {
	results := []Models.Game{}
	err := g.db.Preload("User").
		Table("games").
		Where("user_id = ?", identity).
		Find(&results).Error
	return results, err
}
