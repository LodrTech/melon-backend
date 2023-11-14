package app

import (
	"context"
	"log"
	"net/http"
	"time"
	"github.com/Marif226/melon/internal/config"
	"github.com/go-chi/chi/v5"
)

type App struct {
	httpServer 		*http.Server
	provider		*provider
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	return a.runHTTPerver()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error {
		a.initConfig,
		a.initServiceProvider,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load("./.env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.provider = newProvider()
	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	router := chi.NewRouter()

	setRoutes(router, a.provider.Handlers(ctx))

	a.httpServer = &http.Server{
		Addr:           a.provider.HTTPConfig().Address(),
		Handler:        router,
		ReadTimeout:    10 * time.Second, // 10 sec limit for reading request
		WriteTimeout:   10 * time.Second, // 10 sec limit for writing response
		MaxHeaderBytes: 1 << 20, // memory limit for response header
	}

	a.runHTTPerver()

	return nil
}

func (a *App) runHTTPerver() error {
	log.Printf("HTTP server is running on %s", a.provider.HTTPConfig().Address())

	err := a.httpServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}