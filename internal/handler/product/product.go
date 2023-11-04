package product

import (
	"github.com/Marif226/melon/internal/service"
)

type productHandler struct {
	service.ProductService
}

func NewProductHandler(productService service.ProductService) *productHandler {
	return &productHandler{
		ProductService: productService,
	}
}