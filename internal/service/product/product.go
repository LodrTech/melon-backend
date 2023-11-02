package product

import (
	"log"
	"github.com/Marif226/melon/internal/model"
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

func (s *productService) Create(request model.Product) (*model.Product, error) {
	log.Println("Product.service.Create: ", request)
	return &request, nil
}