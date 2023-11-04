package helper

import (
	"net/http"
	"strconv"
	validate "github.com/Marif226/melon/internal/lib/validator/product"
	"github.com/go-chi/chi/v5"
)

func GetPathParamID(r *http.Request) (int, error) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return 0, err
	}

	err = validate.ValidateID(id) 
	if err != nil {
		return 0, err
	}

	return id, nil
}