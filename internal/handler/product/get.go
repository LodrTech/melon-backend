package product

import (
	"net/http"
	"github.com/Marif226/melon/pkg/helper"
	"github.com/google/jsonapi"
)

func (h *productHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := helper.GetPathParamID(r)
	if err != nil {
		helper.JsonapiError(w, []*jsonapi.ErrorObject{{
			Title: "ID Error`",
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