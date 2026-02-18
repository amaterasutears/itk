package wallet

import (
	"time"

	"github.com/amaterasutears/itk/internal/model/wallet"
	"github.com/google/uuid"
)

type CreateWalletResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

func NewCreateWalletResponse(w *wallet.Wallet) *CreateWalletResponse {
	return &CreateWalletResponse{
		ID:        w.ID,
		CreatedAt: w.CreatedAt,
	}
}
