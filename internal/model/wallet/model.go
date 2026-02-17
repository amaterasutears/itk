package wallet

import (
	"time"

	"github.com/google/uuid"
)

const (
	tableName           string = "wallets"
	idColumnName        string = "id"
	createdAtColumnName string = "created_at"
)

type Wallet struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"created_at"`
}

func New(id uuid.UUID, ca time.Time) *Wallet {
	return &Wallet{
		ID:        id,
		CreatedAt: ca,
	}
}

func (w *Wallet) TableName() string {
	return tableName
}

func (w *Wallet) IDColumnName() string {
	return idColumnName
}

func (w *Wallet) CreatedAtColumnName() string {
	return createdAtColumnName
}
