package transaction

import (
	"context"
	"time"

	"github.com/amaterasutears/itk/internal/model/transaction"
)

type Service struct {
	ws WalletStorage
	ts TransactionStorage
	t  Transactor
}

func New(ws WalletStorage, ts TransactionStorage, t Transactor) *Service {
	return &Service{
		ws: ws,
		ts: ts,
		t:  t,
	}
}

func (s *Service) Create(ctx context.Context, t *transaction.Transaction) error {
	now := time.Now().UTC()

	t.CreatedAt = now

	return s.t.WithinTransaction(ctx, func(ctx context.Context) error {
		err := s.ws.Update(ctx, t, now)
		if err != nil {
			return err
		}

		return s.ts.Create(ctx, t)
	})
}
