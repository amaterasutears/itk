package config

import (
	"errors"
	"fmt"
)

var (
	ErrNilSSLMode     error = errors.New("ssl mode cannot be nil")
	ErrInvalidSSLMode error = errors.New("invalid ssl mode")
)

type SSLMode string

const (
	DisableSSLMode    SSLMode = "disable"
	AllowSSLMode      SSLMode = "allow"
	PreferSSLMode     SSLMode = "prefer"
	RequireSSLMode    SSLMode = "require"
	VerifyCASSLMode   SSLMode = "verify-ca"
	VerifyFullSSLMode SSLMode = "verify-full"
)

func (s *SSLMode) UnmarshalText(text []byte) error {
	if s == nil {
		return ErrNilSSLMode
	}

	switch string(text) {
	case string(DisableSSLMode):
		*s = DisableSSLMode
	case string(AllowSSLMode):
		*s = AllowSSLMode
	case string(PreferSSLMode):
		*s = PreferSSLMode
	case string(RequireSSLMode):
		*s = RequireSSLMode
	case string(VerifyCASSLMode):
		*s = VerifyCASSLMode
	case string(VerifyFullSSLMode):
		*s = VerifyFullSSLMode
	default:
		return ErrInvalidSSLMode
	}

	return nil
}

type Postgres struct {
	User        string  `env:"POSTGRES_USER" env-required`
	Password    string  `env:"POSTGRES_PASSWORD" env-required`
	DB          string  `env:"POSTGRES_DB" env-required`
	Host        string  `env:"POSTGRES_HOST" env-required`
	Port        int     `env:"POSTGRES_PORT" env-required`
	SSLMode     SSLMode `env:"POSTGRES_SSL_MODE" env-required`
	PingTimeout int     `env:"POSTGRES_PING_TIMEOUT_SEC" env-default:"5"`
}

func (p *Postgres) DataSourceName() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		p.User,
		p.Password,
		p.Host,
		p.Port,
		p.DB,
		p.SSLMode,
	)
}
