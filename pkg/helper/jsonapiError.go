package helper

import (
	"net/http"

	"github.com/google/jsonapi"
)

// JsonapiError replies to the request with JSON:api response using given error
func JsonapiError(w http.ResponseWriter, errors []*jsonapi.ErrorObject) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusBadRequest)
	jsonapi.MarshalErrors(w, errors)
}