package sql

import (
	"context"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

const driverName string = "pgx"

type Client struct {
	db *sqlx.DB
}

func New(dataSourceName string, opts ...Option) (*Client, error) {
	db, err := sqlx.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	var opt options

	for _, o := range opts {
		o(&opt)
	}

	if opt.maxOpenConns != nil {
		db.SetMaxOpenConns(*opt.maxOpenConns)
	}

	if opt.maxIdleConns != nil {
		db.SetMaxIdleConns(*opt.maxIdleConns)
	}

	if opt.connMaxLifetime != nil {
		db.SetConnMaxLifetime(*opt.connMaxLifetime)
	}

	if opt.connMaxIdleTime != nil {
		db.SetConnMaxIdleTime(*opt.connMaxIdleTime)
	}

	return &Client{
		db: db,
	}, nil
}

func (c *Client) Ping(ctx context.Context) error {
	return c.db.PingContext(ctx)
}

func (c *Client) DB() *sqlx.DB {
	return c.db
}

func (c *Client) Close() error {
	return c.db.Close()
}
