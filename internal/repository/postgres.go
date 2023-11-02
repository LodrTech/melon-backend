package repository

import (
	"github.com/Marif226/melon/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB(cfg config.PGConfig) (*sqlx.DB, error) {
	connectionStr := cfg.ConnectionString()
	
	db, err := sqlx.Open("postgres", connectionStr)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}