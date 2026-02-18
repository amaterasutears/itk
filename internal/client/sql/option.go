package sql

import "time"

type options struct {
	maxOpenConns    *int
	maxIdleConns    *int
	connMaxLifetime *time.Duration
	connMaxIdleTime *time.Duration
}

type Option func(opts *options)

func WithMaxOpenConns(n int) Option {
	return func(opts *options) {
		opts.maxOpenConns = &n
	}
}

func WithMaxIdleConns(n int) Option {
	return func(opts *options) {
		opts.maxIdleConns = &n
	}
}

func WithConnMaxLifetime(d time.Duration) Option {
	return func(opts *options) {
		opts.connMaxLifetime = &d
	}
}

func WithConnMaxIdleTime(d time.Duration) Option {
	return func(opts *options) {
		opts.connMaxIdleTime = &d
	}
}
