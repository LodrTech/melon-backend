package product

import (
	"github.com/jackc/pgx/v5"
)

type productRepo struct {
	db *pgx.Conn
}

func NewProductRepo(conn *pgx.Conn) *productRepo {
	return &productRepo {
		db: conn,
	}
}