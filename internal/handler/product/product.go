package product

import (
	"log/slog"
	"github.com/Marif226/melon/internal/service"
)

type productHandler struct {
	log	*slog.Logger
	service.ProductService
}

func NewProductHandler(log	*slog.Logger, productService service.ProductService) *productHandler {
	return &productHandler{
		log: log,
		ProductService: productService,
	}
}