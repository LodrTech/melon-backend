package app

import (
	"log"

	"github.com/Marif226/melon/internal/config"
	"github.com/Marif226/melon/internal/handler"
	"github.com/Marif226/melon/internal/service"
)

type provider struct {
	httpConfig config.HTTPConfig

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

func (p *provider) Services() *service.Provider {
	if p.services == nil {
		p.services = service.NewProvider()
	}

	return p.services
}

func (p *provider) Handlers() *handler.Provider{
	if p.handlers == nil {
		p.handlers = handler.NewProvider(p.Services())
	}

	return p.handlers
}