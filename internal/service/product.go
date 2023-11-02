package service

import (
	"log"
	"github.com/Marif226/melon/internal/model"
	"github.com/Marif226/melon/internal/repository"
)

type productServiceImpl struct {
	productRepo repository.ProductRepo
}

func NewProductService(productRepo repository.ProductRepo) ProductService {
	return &productServiceImpl{
		productRepo: productRepo,
	}
}

func (s *productServiceImpl) Create(request model.Product) (*model.Product, error) {
	log.Println("Product.service.Create: ", request)
	return &request, nil
}