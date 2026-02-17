package migrator

import (
	"context"
	"embed"
	"io/fs"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
)

const (
	migrationsDirectory string = "migrations"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

type Migrator struct {
	provider *goose.Provider
}

func New(db *sqlx.DB) (*Migrator, error) {
	migrationFS, err := fs.Sub(embedMigrations, migrationsDirectory)
	if err != nil {
		return nil, err
	}

	provider, err := goose.NewProvider(goose.DialectPostgres, db.DB, migrationFS)
	if err != nil {
		return nil, err
	}

	return &Migrator{
		provider: provider,
	}, nil
}

func (m *Migrator) Up(ctx context.Context) error {
	_, err := m.provider.Up(ctx)
	if err != nil {
		return err
	}

	return nil
}
