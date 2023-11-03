package repository

import (
	"context"

	"github.com/Marif226/melon/internal/config"
	"github.com/jackc/pgx/v5"
)

func NewPostgresDB(ctx context.Context, cfg config.PGConfig) (*pgx.Conn, error) {
	connectionStr := cfg.ConnectionString()

	conn, err := pgx.Connect(ctx, connectionStr)
	if err != nil {
		return nil, err
	}

	defer conn.Close(ctx)

	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return conn, nil
}