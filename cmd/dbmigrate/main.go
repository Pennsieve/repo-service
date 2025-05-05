package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/pennsieve/dbmigrate-go/pkg/dbmigrate"
	"github.com/pennsieve/dbmigrate-go/pkg/shared/config"
	"log/slog"
	"os"
)

var logger = slog.Default()

//go:embed migrations/*.sql
var migrationsFS embed.FS

func main() {
	ctx := context.Background()
	defaultSettings := config.NewDefaultSettings()
	defaultSettings["POSTGRES_SCHEMA"] = "repositories"

	migrateConfig, err := config.LoadConfig(defaultSettings)
	if err != nil {
		logger.Error("error loading config", slog.Any("error", err))
		os.Exit(1)
	}
	if migrateConfig.PostgresDB.Password == nil {
		logger.Error("password must be specified; cannot currently use RDS proxy for migrates since no Postgres role with the appropriate grants has credentials in the proxy")
		os.Exit(1)
	}
	logger.
		With(slog.Bool("verboseLogging", migrateConfig.VerboseLogging),
			slog.Group("postgres",
				slog.String("host", migrateConfig.PostgresDB.Host),
				slog.Int("port", migrateConfig.PostgresDB.Port),
				slog.String("username", migrateConfig.PostgresDB.User),
				slog.String("database", migrateConfig.PostgresDB.Database),
				slog.String("schema", migrateConfig.PostgresDB.Schema),
			)).
		Info("DB schema migration started")

	migrationsSource, err := iofs.New(migrationsFS, "migrations")
	if err != nil {
		logger.Error(fmt.Errorf("error creating migration iofs source.Driver: %w", err).Error())
		os.Exit(1)
	}

	m, err := dbmigrate.NewLocalMigrator(ctx, migrateConfig, migrationsSource)
	if err != nil {
		logger.Error("error creating DatabaseMigrator", slog.Any("error", err))
		os.Exit(1)
	}
	defer m.CloseAndLogError()

	if err := m.Up(); err != nil {
		logger.Error("error running 'up' migrations", slog.Any("error", err))
		os.Exit(1)
	}

	logger.Info("Database schema migration complete")
}
