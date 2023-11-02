package app

import (
	"log"
	"github.com/Marif226/melon/internal/config"
	"github.com/Marif226/melon/internal/handler"
	"github.com/Marif226/melon/internal/repository"
	"github.com/Marif226/melon/internal/service"
	"github.com/jmoiron/sqlx"
)

type provider struct {
	httpConfig 	config.HTTPConfig
	pgConfig	config.PGConfig
	postgres	*sqlx.DB
	repos		*repository.Repository
	services	*service.Provider
	handlers	*handler.Provider
}

func newProvider() *provider {
	return &provider{}
}

func (p *provider) HTTPConfig() config.HTTPConfig {
	if p.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		p.httpConfig = cfg
	}

	return p.httpConfig
}

func (p *provider) PGConfig() config.PGConfig {
	if p.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get postgres config: %s", err.Error())
		}

		p.pgConfig = cfg
	}

	return p.pgConfig
}

func (p *provider) Postgres() *sqlx.DB {
	if p.postgres == nil {
		postgres, err := repository.NewPostgresDB(p.PGConfig())
		if err != nil {
			log.Fatalf("failed to get postgres: %s", err.Error())
		}

		p.postgres = postgres
	}

	return p.postgres
}

func (p *provider) Repos() *repository.Repository {
	if p.repos == nil {
		p.repos = repository.New(p.Postgres())
	}

	return p.repos
}

func (p *provider) Services() *service.Provider {
	if p.services == nil {
		p.services = service.NewProvider(p.Repos())
	}

	return p.services
}

func (p *provider) Handlers() *handler.Provider{
	if p.handlers == nil {
		p.handlers = handler.NewProvider(p.Services())
	}

	return p.handlers
}