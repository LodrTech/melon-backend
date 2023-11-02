package repository

import (
	"github.com/Marif226/melon/internal/model"
)

type ProductRepo interface {
	Create(request model.Product) (*model.Product, error)
}