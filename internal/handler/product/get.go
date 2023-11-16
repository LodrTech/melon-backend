package product

import (
	"fmt"
	"net/http"

	"github.com/Marif226/melon/pkg/helper"
	"github.com/google/jsonapi"
)

func (h *productHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := helper.GetPathParamID(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
			Title: "Invalid Request",
			Detail: err.Error(),
			Status: fmt.Sprint(http.StatusBadRequest),
		}})
		return
	}

	response, err := h.ProductService.Get(r.Context(), id)
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