package store

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
	"go.uber.org/zap"

	"github.com/shiroemons/touhou_arrangement_chronicle/go/pkg/config"
)

// PostgresProvider Fx Provider
func PostgresProvider(cfg config.Config, logger *zap.Logger) *bun.DB {
	pgConfig, err := pgx.ParseConfig(cfg.ConnectURL)
	if err != nil {
		log.Fatal("postgres provider error", err)
	}
	conn := bun.NewDB(stdlib.OpenDB(*pgConfig), pgdialect.New())

	conn.AddQueryHook(bundebug.NewQueryHook(
		// disable the hook
		bundebug.WithEnabled(false),

		// BUNDEBUG=1 logs failed queries
		// BUNDEBUG=2 logs all queries
		bundebug.FromEnv("BUNDEBUG"),
	))

	var v string
	if err = conn.NewSelect().ColumnExpr("version()").Scan(context.Background(), &v); err != nil {
		log.Fatal(err)
	}
	logger.Info(v)

	return conn
}
