package repository

import (
	"github.com/Marif226/melon/internal/repository/product"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	ProductRepo
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		ProductRepo: product.NewProductRepo(db),
	}
}