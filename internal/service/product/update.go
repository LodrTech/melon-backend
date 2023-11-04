package product

import (
	"context"

	"github.com/Marif226/melon/internal/model"
)

func (s *productService) Update(ctx context.Context, request model.ProductUpdateRequest) (*model.Product, error) {
	response, err := s.productRepo.Update(ctx, request)
	if err != nil {
		return nil, err
	}
	return response, nil
}