package product

import (
	"fmt"
	"net/http"
	validate "github.com/Marif226/melon/internal/lib/validator/product"
	"github.com/Marif226/melon/internal/model"
	"github.com/google/jsonapi"
)

func (h *productHandler) Create(w http.ResponseWriter, r *http.Request) {
	var request model.Product

	err := jsonapi.UnmarshalPayload(r.Body, &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
			Title: "Invalid Request Body",
			Detail: err.Error(),
			Status: fmt.Sprint(http.StatusBadRequest),
		}})
		return
	}

	err = validate.Create(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
			Title: "Validation Error",
			Detail: err.Error(),
			Status: fmt.Sprint(http.StatusBadRequest),
		}})
		return
	}

	response, err := h.ProductService.Create(r.Context(), request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
			Title: "Server Error",
			Detail: err.Error(),
			Status: fmt.Sprint(http.StatusInternalServerError),
		}})
		return
	}
	w.WriteHeader(http.StatusCreated)

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