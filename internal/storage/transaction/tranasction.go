package transaction

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/amaterasutears/itk/internal/model/transaction"
	"github.com/amaterasutears/itk/internal/storage/transactor"
	"github.com/jmoiron/sqlx"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

type Storage struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) Create(ctx context.Context, t *transaction.Transaction) error {
	insertb := psql.Insert("transactions").Columns(
		"wallet_id",
		"operation_type",
		"amount",
		"created_at",
	).Values(
		t.WalletID,
		t.OperationType,
		t.Amount,
		t.CreatedAt,
	)

	query, args, err := insertb.ToSql()
	if err != nil {
		return err
	}

	_, err = transactor.Executor(ctx, s.db).ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
