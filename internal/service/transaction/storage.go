package transaction

import (
	"context"

	"github.com/amaterasutears/itk/internal/model/transaction"
)

type WalletStorage interface {
	Update(ctx context.Context, t *transaction.Transaction) error
}

type TransactionStorage interface {
	Create(ctx context.Context, t *transaction.Transaction) error
}
