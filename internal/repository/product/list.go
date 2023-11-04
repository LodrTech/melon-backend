package product

import (
	"context"
	"github.com/Marif226/melon/internal/lib/querybuilder"
	"github.com/Marif226/melon/internal/model"
)

func (r *productRepo) List(ctx context.Context, request model.ProductListRequest) ([]*model.Product, error) {
	query, args, err := querybuilder.ProductList(request)
	if err != nil {
		return nil, nil
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	products := make([]*model.Product, 0)
	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Weight)
		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return products, nil
}