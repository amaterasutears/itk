package wallet

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/amaterasutears/itk/internal/model/wallet"
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

func (s *Storage) Insert(ctx context.Context, w *wallet.Wallet) (*wallet.Wallet, error) {
	insertb := psql.Insert(w.TableName()).Columns(w.IDColumnName(), w.CreatedAtColumnName()).
		Values(w.ID, w.CreatedAt).Suffix(fmt.Sprintf("RETURNING %s, %s", w.IDColumnName(), w.CreatedAtColumnName()))

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
