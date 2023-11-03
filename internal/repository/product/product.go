package product

import (
	"github.com/Marif226/melon/internal/model"
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

func (r *productRepo) Create(request model.Product) (*model.Product, error) {


	return nil, nil
}