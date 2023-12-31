package app

import (
	"context"
	"log/slog"
	"net/http"
	"time"
	"github.com/Marif226/melon/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type App struct {
	httpServer 		*http.Server
	provider		*provider
	log				*slog.Logger
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
	router := initRouter(a.provider.log, a.provider.Handlers(ctx))

	a.httpServer = &http.Server{
		Addr:           a.provider.Config().Address(),
		Handler:        router,
		ReadTimeout:    10 * time.Second, // 10 sec limit for reading request
		WriteTimeout:   10 * time.Second, // 10 sec limit for writing response
		MaxHeaderBytes: 1 << 20, // memory limit for response header
	}

	return nil
}

func (a *App) runHTTPerver() error {
	a.provider.log.Info("HTTP server is running", 
		slog.String("address", a.provider.Config().Address()),
	)

	err := a.httpServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}