package handler

import (
	"log/slog"
	"github.com/Marif226/melon/internal/handler/product"
	"github.com/Marif226/melon/internal/service"
)

type Provider struct {
	ProductHandler
}

func NewProvider(log *slog.Logger, services *service.Provider) *Provider {
	return &Provider{
		ProductHandler: product.NewProductHandler(log, services.ProductService),
	}
}