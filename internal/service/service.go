package service

import (
	"context"
	"github.com/Marif226/melon/internal/model"
)

type ProductService interface {
	Create(ctx context.Context, request model.Product) (*model.Product, error)
	List(ctx context.Context, request model.ProductListRequest) ([]*model.Product, error)
}