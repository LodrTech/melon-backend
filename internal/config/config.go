package config

import (
	"errors"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	MigrationURL	string
	HTTPConfig
	PGConfig
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}

func NewConfig() (*Config, error) {
	migrationURL := os.Getenv("MIGRATION_URL")
	if len(migrationURL) == 0 {
		return nil, errors.New("migration url not found")
	}

	httpConfig, err := NewHTTPConfig()
	if err != nil {
		return nil, err
	}

	pgConfig, err := NewPGConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		MigrationURL: migrationURL,
		HTTPConfig: httpConfig,
		PGConfig: pgConfig,
	}, nil
}