package db

import (
	"embed"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"os"

	"go.uber.org/zap"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func SetupPostgres(pool *pgxpool.Pool, logger *zap.Logger) {
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		logger.Error("unable to set postgres dialect", zap.Error(err))
		os.Exit(-1)
	}

	db := stdlib.OpenDBFromPool(pool)
	if err := goose.Up(db, "migrations"); err != nil {
		logger.Error("unable to up migrations", zap.Error(err))
		os.Exit(-1)
	}
}
