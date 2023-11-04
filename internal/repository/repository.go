package repository

import (
	"context"

	"github.com/Marif226/melon/internal/model"
)

type ProductRepo interface {
	Create(ctx context.Context, request model.Product) (model.Product, error)
	List(ctx context.Context, request model.ProductListRequest) ([]*model.Product, error)
	Get(ctx context.Context, id int) (*model.Product, error)
}