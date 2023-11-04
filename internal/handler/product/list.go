package product

import (
	"net/http"
	"github.com/Marif226/melon/internal/model"
	"github.com/google/jsonapi"
)

func (h *productHandler) List(w http.ResponseWriter, r *http.Request) {
	var request model.ProductListRequest

	// err := jsonapi.UnmarshalPayload(r.Body, &request)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	response, err := h.ProductService.List(r.Context(), request)
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