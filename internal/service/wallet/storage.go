package wallet

import (
	"context"

	"github.com/amaterasutears/itk/internal/model/wallet"
)

type WalletStorage interface {
	Insert(ctx context.Context, w *wallet.Wallet) (*wallet.Wallet, error)
}
