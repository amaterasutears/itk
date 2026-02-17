package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Postgres
}

func Load() (*Config, error) {
	var c Config

	err := cleanenv.ReadConfig("./.env", &c)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
