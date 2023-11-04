package product

import (
	"github.com/Marif226/melon/internal/repository"
)

type productService struct {
	productRepo repository.ProductRepo
}

func NewProductService(productRepo repository.ProductRepo) *productService {
	return &productService {
		productRepo: productRepo,
	}
}