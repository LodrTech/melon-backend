package handler

import "github.com/Marif226/melon/internal/service"

type Provider struct {
	ProductHandler
}

func NewProvider(services *service.Provider) *Provider {
	return &Provider{
		ProductHandler: NewProductHandler(services.ProductService),
	}
}