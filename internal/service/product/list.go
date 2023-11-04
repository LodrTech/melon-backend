package product

import (
	"context"

	"github.com/Marif226/melon/internal/model"
)

func (s *productService) List(ctx context.Context, request model.ProductListRequest) ([]*model.Product, error) {
	response, err := s.productRepo.List(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}