package app

import (
	"context"
	"errors"
	"log"
	"github.com/Marif226/melon/internal/config"
	"github.com/Marif226/melon/internal/handler"
	"github.com/Marif226/melon/internal/repository"
	"github.com/Marif226/melon/internal/service"
	"github.com/Marif226/melon/pkg/client/postgres"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
)

type provider struct {
	config		config.Config
	httpConfig 	config.HTTPConfig
	pgConfig	config.PGConfig
	postgres	*pgx.Conn
	repos		*repository.Repository
	services	*service.Provider
	handlers	*handler.Provider
}

func newProvider() *provider {
	return &provider{}
}

func (p *provider) Config() config.Config {
	if p.config == nil {
		config, err := config.NewConfig()
		if err != nil {
			log.Fatalf("failed to get config: %s", err.Error())
		}

		p.config = config
	}

	return p.config
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

func (p *provider) Postgres(ctx context.Context) *pgx.Conn {
	if p.postgres == nil {
		postgres, err := postgres.NewPostgresDB(ctx, p.PGConfig())
		if err != nil {
			log.Fatalf("failed to get postgres: %s", err.Error())
		}

		p.postgres = postgres
	}

	err := runDBMigrations(p.Config().MigrationURL(), p.PGConfig().ConnectionString())
	if err != nil {
		log.Fatalf("failed to migrate: %s", err.Error())
	}

	return p.postgres
}

func runDBMigrations(migrationURL string, postgresURL string) error {
	m, err := migrate.New(migrationURL, postgresURL)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("migrate up:", err)
			return nil
		}
		return err
	}

	log.Println("db migrated successfully")

	return nil
}

func (p *provider) Repos(ctx context.Context) *repository.Repository {
	if p.repos == nil {
		p.repos = repository.New(p.Postgres(ctx))
	}

	return p.repos
}

func (p *provider) Services(ctx context.Context) *service.Provider {
	if p.services == nil {
		p.services = service.NewProvider(p.Repos(ctx))
	}

	return p.services
}

func (p *provider) Handlers(ctx context.Context) *handler.Provider{
	if p.handlers == nil {
		p.handlers = handler.NewProvider(p.Services(ctx))
	}

	return p.handlers
}