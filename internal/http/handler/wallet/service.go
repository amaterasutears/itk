package wallet

import (
	"context"

	"github.com/amaterasutears/itk/internal/model/wallet"
	"github.com/google/uuid"
)

type WalletService interface {
	Create(ctx context.Context) (*wallet.Wallet, error)
	Balance(ctx context.Context, wid uuid.UUID) (int, error)
}
