package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Postgres
	Server
	Migrator
}

func Load() (*Config, error) {
	var c Config

	err := cleanenv.ReadEnv(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
