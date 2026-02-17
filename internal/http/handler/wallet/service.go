package wallet

import (
	"context"

	"github.com/amaterasutears/itk/internal/model/wallet"
)

type WalletService interface {
	Create(ctx context.Context) (*wallet.Wallet, error)
}
