package Common

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Identify struct {
	ID uuid.UUID `gorm:"primary_key" json:"id"`
}

type Times struct {
	CreatedAt time.Time `json:"created_at" gorm:"type:time"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:time"`
}

type Repository interface {
	TxStart() (*gorm.DB, error)
	TxCommit(tx *gorm.DB) error
	TxRollback(tx *gorm.DB)
}

func (base *Identify) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}
