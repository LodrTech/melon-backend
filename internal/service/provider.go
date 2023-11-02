package service

import (
	"github.com/Marif226/melon/internal/repository"
	"github.com/Marif226/melon/internal/service/product"
)

type Provider struct {
	ProductService
}

func NewProvider(repos *repository.Repository) *Provider {
	return &Provider{
		ProductService: product.NewProductService(repos.ProductRepo),
	}
}