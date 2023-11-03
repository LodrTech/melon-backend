package repository

import (
	"github.com/Marif226/melon/internal/repository/product"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	ProductRepo
}

func New(conn *pgx.Conn) *Repository {
	return &Repository{
		ProductRepo: product.NewProductRepo(conn),
	}
}