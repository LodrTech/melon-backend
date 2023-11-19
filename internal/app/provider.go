package app

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"github.com/Marif226/melon/internal/config"
	"github.com/Marif226/melon/internal/handler"
	"github.com/Marif226/melon/internal/lib/logger/sl"
	"github.com/Marif226/melon/internal/repository"
	"github.com/Marif226/melon/internal/service"
	"github.com/Marif226/melon/pkg/client/postgres"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
)

type provider struct {
	config		*config.Config
	log 		*slog.Logger
	postgres	*pgx.Conn
	repos		*repository.Repository
	services	*service.Provider
	handlers	*handler.Provider
}

func newProvider() *provider {
	p := &provider{}
	p.setupLogger()

	return p
}

func (p *provider) setupLogger() {
	switch p.Config().Env {
	case envLocal:
		p.log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		p.log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		p.log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
}

func (p *provider) Config() *config.Config {
	if p.config == nil {
		config, err := config.NewConfig()
		if err != nil {
			p.log.Error("failed to get config", sl.Err(err))
			os.Exit(1)
		}

		p.config = config
	}

	return p.config
}

func (p *provider) Postgres(ctx context.Context) *pgx.Conn {
	if p.postgres == nil {
		postgres, err := postgres.NewPostgresDB(ctx, p.Config().PGConfig)
		p.log.Debug("init config",
			slog.String("connection string",  p.Config().ConnectionString()),
		)
		if err != nil {
			p.log.Error("failed to get postgres", sl.Err(err))
			os.Exit(1)
		}

		p.postgres = postgres
	}

	err := p.runDBMigrations(p.Config().MigrationURL, p.Config().ConnectionString())
	if err != nil {
		p.log.Error("failed to migrate", sl.Err(err))
		os.Exit(1)
	}

	return p.postgres
}

func (p *provider) runDBMigrations(migrationURL string, postgresURL string) error {
	m, err := migrate.New(migrationURL, postgresURL)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			p.log.Debug("migrate up", sl.Err(err))
			return nil
		}
		return err
	}

	p.log.Info("db migrated successfully")

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
		p.handlers = handler.NewProvider(p.log, p.Services(ctx))
	}

	return p.handlers
}