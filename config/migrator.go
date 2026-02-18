package config

type Migrator struct {
	TimeoutSec int `env:"MIGRATOR_TIMEOUT_SEC" env-default:"10"`
}
