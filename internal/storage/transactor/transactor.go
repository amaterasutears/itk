package transactor

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type txKey struct{}

type Transactor struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Transactor {
	return &Transactor{
		db: db,
	}
}

func (t *Transactor) WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := t.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	err = fn(injectTx(ctx, tx))
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func Executor(ctx context.Context, db *sqlx.DB) sqlx.ExtContext {
	tx := extractTx(ctx)
	if tx != nil {
		return tx
	}

	return db
}

func injectTx(ctx context.Context, tx *sqlx.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

func extractTx(ctx context.Context) *sqlx.Tx {
	tx, ok := ctx.Value(txKey{}).(*sqlx.Tx)
	if ok {
		return tx
	}

	return nil
}
