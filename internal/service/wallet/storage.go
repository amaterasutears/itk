package wallet

import (
	"context"

	"github.com/amaterasutears/itk/internal/model/wallet"
	"github.com/google/uuid"
)

type WalletStorage interface {
	Insert(ctx context.Context, w *wallet.Wallet) (*wallet.Wallet, error)
	Balance(ctx context.Context, wid uuid.UUID) (int, error)
}
