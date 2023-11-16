package product

import (
	"fmt"
	"net/http"

	"github.com/Marif226/melon/internal/model"
	"github.com/google/jsonapi"
)

func (h *productHandler) List(w http.ResponseWriter, r *http.Request) {
	var request model.ProductListRequest

	response, err := h.ProductService.List(r.Context(), request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
			Title: "Server Error",
			Detail: err.Error(),
			Status: fmt.Sprint(http.StatusInternalServerError),
		}})
		return
	}

	w.WriteHeader(http.StatusOK)

	err = jsonapi.MarshalPayload(w, response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
			Title: "Server Error",
			Detail: err.Error(),
			Status: fmt.Sprint(http.StatusInternalServerError),
		}})
		return
	}
}