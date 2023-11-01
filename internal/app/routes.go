package app

import (
	"github.com/Marif226/melon/internal/handler"
	"github.com/go-chi/chi/v5"
)

func setRoutes(router chi.Router, handlers *handler.Provider) {
	// router.Route("/products", func(r chi.Router) {
		
	// })

	router.Post("/products", handlers.Create)
}