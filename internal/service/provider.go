package service

import "github.com/Marif226/melon/internal/repository"

type Provider struct {
	ProductService
}

func NewProvider(repos *repository.Repository) *Provider {
	return &Provider{
		ProductService: NewProductService(repos.ProductRepo),
	}
}