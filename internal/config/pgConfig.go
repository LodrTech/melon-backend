package config

import (
	"errors"
	"fmt"
	"os"
)

type PGConfig interface {
	ConnectionString() string
}

const (
	pgHostEnvName = "POSTGRES_HOST"
	pgPortEnvName = "POSTGRES_PORT"
	pgUserEnvName = "POSTGRES_USER"
	pgPasswordEnvName = "POSTGRES_PASSWORD"
	pgDBEnvName = "POSTGRES_DB"
	pgSSLModeEnvName = "POSTGRES_SSL_MODE"
)

type pgConfig struct {
	host     	string
	port     	string
	username 	string
	password 	string
	dbName   	string
	sslMode  	string
}

func NewPGConfig() (PGConfig, error) {
	host := os.Getenv(pgHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("postgres host not found")
	}

	port := os.Getenv(pgPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("postgres port not found")
	}

	username := os.Getenv(pgUserEnvName)
	if len(username) == 0 {
		return nil, errors.New("postgres username not found")
	}

	password := os.Getenv(pgPasswordEnvName)
	if len(password) == 0 {
		return nil, errors.New("postgres password not found")
	}

	dbName := os.Getenv(pgDBEnvName)
	if len(dbName) == 0 {
		return nil, errors.New("postgres db name not found")
	}

	sslMode := os.Getenv(pgSSLModeEnvName)
	if len(sslMode) == 0 {
		return nil, errors.New("postgres ssl mode not found")
	}

	return &pgConfig{
		host: host,
		port: port,
		username: username,
		password: password,
		dbName: dbName,
		sslMode: sslMode,
	}, nil
}

func (cfg *pgConfig) ConnectionString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.username, cfg.password, cfg.host, cfg.port, cfg.dbName, cfg.sslMode)
}