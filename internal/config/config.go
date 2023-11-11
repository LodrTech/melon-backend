package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	MigrationURL() string
}

type config struct {
	migrationURL	string
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}

func NewConfig() (Config, error) {
	migrationURL := os.Getenv("MIGRATION_URL")
	if len(migrationURL) == 0 {
		return nil, errors.New("migration url not found")
	}

	return &config{
		migrationURL: migrationURL,
	}, nil
}

func (c *config) MigrationURL() string {
	return c.migrationURL
}