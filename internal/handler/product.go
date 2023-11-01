package handler

import (
	"net/http"

	"github.com/Marif226/melon/internal/model"
	"github.com/Marif226/melon/internal/service"
	"github.com/google/jsonapi"
)

type productHandlerImpl struct {
	service.ProductService
}

func NewProductHandler(productService service.ProductService) ProductHandler {
	return &productHandlerImpl{
		ProductService: productService,
	}
}

func (h *productHandlerImpl) Create(w http.ResponseWriter, r *http.Request) {
	request := &model.Product{}

	err := jsonapi.UnmarshalPayload(r.Body, request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := h.ProductService.Create(*request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(http.StatusOK)

	err = jsonapi.MarshalPayload(w, response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}