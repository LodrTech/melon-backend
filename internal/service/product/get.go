package product

import (
	"context"

	"github.com/Marif226/melon/internal/model"
)

func (s *productService) Get(ctx context.Context, id int) (*model.Product, error) {
	response, err := s.productRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return response, nil
}