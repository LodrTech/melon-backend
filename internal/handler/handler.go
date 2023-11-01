package handler

import "net/http"

type ProductHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
}