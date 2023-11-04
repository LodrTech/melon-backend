package product

import (
	"context"

	"github.com/Marif226/melon/internal/model"
)

func (s *productService) Create(ctx context.Context, request model.Product) (*model.Product, error) {
	response, err := s.productRepo.Create(ctx, request)
	if err != nil {
		return nil, err
	}
	return &response, nil
}