package app

import (
	"log/slog"
	"github.com/Marif226/melon/internal/handler"
	mw "github.com/Marif226/melon/internal/middleware"
	mwLogger "github.com/Marif226/melon/internal/middleware/logger"
	"github.com/go-chi/chi/v5"
)

func initRouter(log *slog.Logger, handlers *handler.Provider) chi.Router {
	router := chi.NewRouter()

	router.Use(mwLogger.New(log))
	router.Use(mw.JsonapiMediaTypeMiddleware)

	setRoutes(router, handlers)

	return router
}

func setRoutes(router chi.Router, handlers *handler.Provider) {
	router.Route("/products", func(r chi.Router) {
		r.Post("/", handlers.ProductHandler.Create)
		r.Get("/", handlers.ProductHandler.List)
		r.Get("/{id}", handlers.ProductHandler.Get)
		r.Patch("/", handlers.ProductHandler.Update)
	})
}
