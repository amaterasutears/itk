package transaction

import (
	"context"

	"github.com/amaterasutears/itk/internal/model/transaction"
)

type TransactionService interface {
	Create(ctx context.Context, t *transaction.Transaction) error
}
