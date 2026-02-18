package wallet

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	ID        uuid.UUID `db:"id"`
	Balance   int       `db:"balance"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func New(id uuid.UUID, ca, ua time.Time) *Wallet {
	return &Wallet{
		ID:        id,
		CreatedAt: ca,
		UpdatedAt: ua,
	}
}
