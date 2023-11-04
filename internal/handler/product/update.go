package product

import (
	"net/http"
	validate "github.com/Marif226/melon/internal/lib/validator/product"
	"github.com/Marif226/melon/internal/model"
	"github.com/Marif226/melon/pkg/helper"
	"github.com/google/jsonapi"
)

func (h *productHandler) Update(w http.ResponseWriter, r *http.Request) {
	var request model.ProductUpdateRequest

	err := jsonapi.UnmarshalPayload(r.Body, &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validate.Update(request)
	if err != nil {
		helper.JsonapiError(w, []*jsonapi.ErrorObject{{
			Title:  "Validation Error",
			Detail: err.Error(),
			Status: "400",
		}})
		return
	}

	response, err := h.ProductService.Update(r.Context(), request)
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