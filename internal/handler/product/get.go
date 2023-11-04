package product

import (
	"net/http"
	"strconv"
	validate "github.com/Marif226/melon/internal/lib/validator/product"
	"github.com/Marif226/melon/pkg/helper"
	"github.com/go-chi/chi/v5"
	"github.com/google/jsonapi"
)

func (h *productHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helper.JsonapiError(w, []*jsonapi.ErrorObject{{
			Title: "ID Error`",
			Detail: err.Error(),
			Status: "400",
		}})
		return
	}

	err = validate.ValidateID(id) 
	if err != nil {
		helper.JsonapiError(w, []*jsonapi.ErrorObject{{
			Title: "ID error`",
			Detail: err.Error(),
			Status: "400",
		}})
		return
	}

	response, err := h.ProductService.Get(r.Context(), id)
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