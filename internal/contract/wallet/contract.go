package wallet

import (
	"time"

	"github.com/google/uuid"
)

type CreateWalletResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

func NewCreateWalletResponse(id uuid.UUID, ca time.Time) *CreateWalletResponse {
	return &CreateWalletResponse{
		ID:        id,
		CreatedAt: ca,
	}
}
