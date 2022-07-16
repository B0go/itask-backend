package config

import (
	"errors"

	"github.com/caarlos0/env"
)

type Config struct {
	DatabaseConnString string `env:"DATABASE_CONNECTION_STRING"`
}

func LoadConfig() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, errors.New("failed to load configuration from environment")
	}
	return &cfg, nil
}
