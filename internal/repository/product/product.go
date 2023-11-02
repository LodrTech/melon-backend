package product

import (
	"github.com/Marif226/melon/internal/model"
	"github.com/jmoiron/sqlx"
)

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) *productRepo {
	return &productRepo {
		db: db,
	}
}

func (r *productRepo) Create(request model.Product) (*model.Product, error) {
	return nil, nil
}