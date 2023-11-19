package middleware

import (
	"net/http"
	"github.com/google/jsonapi"
)

func JsonapiMediaTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", jsonapi.MediaType)
		next.ServeHTTP(w, r)
	})
}