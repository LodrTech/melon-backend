package product

import (
	"context"
	"github.com/Marif226/melon/internal/lib/querybuilder"
	"github.com/Marif226/melon/internal/model"
)

func (r *productRepo) Update(ctx context.Context, request model.ProductUpdateRequest) (*model.Product, error) {
	query, args, err := querybuilder.ProductUpdate(request)
	if err != nil {
		return nil, err
	}

	product := &model.Product{}
	err = r.db.QueryRow(ctx, query, args...).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Weight)
	if err != nil {
		return nil, err
	}

	return product, nil
}