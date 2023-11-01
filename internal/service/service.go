package service

import "github.com/Marif226/melon/internal/model"

type ProductService interface {
	Create(request model.Product) (*model.Product, error)
}