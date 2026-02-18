package wallet

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/amaterasutears/itk/internal/model/transaction"
	"github.com/amaterasutears/itk/internal/model/wallet"
	"github.com/amaterasutears/itk/internal/storage/transactor"
	"github.com/jmoiron/sqlx"
)

const table string = "wallets"

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

type Storage struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) Insert(ctx context.Context, w *wallet.Wallet) (*wallet.Wallet, error) {
	insertb := psql.Insert(table).Columns("id", "balance", "created_at", "updated_at").
		Values(w.ID, w.Balance, w.CreatedAt, w.UpdatedAt).Suffix("RETURNING id, created_at")

	query, args, err := insertb.ToSql()
	if err != nil {
		return nil, err
	}

	var cw wallet.Wallet

	err = s.db.QueryRowxContext(ctx, query, args...).StructScan(&cw)
	if err != nil {
		return nil, err
	}

	return &cw, nil
}

func (s *Storage) Update(ctx context.Context, t *transaction.Transaction, ua time.Time) error {
	updateb := psql.Update(table)

	switch t.OperationType {
	case transaction.DepositOperationType:
		updateb = updateb.Set("balance", squirrel.Expr("balance + ?", t.Amount)).
			Set("updated_at", ua).Where(squirrel.Eq{"id": t.WalletID})
	case transaction.WithdrawOperationType:
		updateb = updateb.Set("balance", squirrel.Expr("balance - ?", t.Amount)).
			Set("updated_at", ua).Where(
			squirrel.And{
				squirrel.Eq{"id": t.WalletID},
				squirrel.GtOrEq{"balance": t.Amount},
			},
		)
	default:
		return transaction.ErrInvalidOperationType
	}

	query, args, err := updateb.ToSql()
	if err != nil {
		return err
	}

	_, err = transactor.Executor(ctx, s.db).ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
