package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type pgConfig interface {
	ConnectionString() string
}

func NewPostgresDB(ctx context.Context, cfg pgConfig) (*pgx.Conn, error) {
	connectionStr := cfg.ConnectionString()

	conn, err := pgx.Connect(ctx, connectionStr)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return conn, nil
}