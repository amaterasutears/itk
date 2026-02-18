package transaction

import (
	"context"
	"time"

	"github.com/amaterasutears/itk/internal/model/transaction"
)

type WalletStorage interface {
	Update(ctx context.Context, t *transaction.Transaction, ua time.Time) error
}

type TransactionStorage interface {
	Create(ctx context.Context, t *transaction.Transaction) error
}
