package db

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	dsn  string
	pool *pgxpool.Pool
}

func NewDatabase(ctx context.Context, dsn string) (*Database, error) {
	pool, err := pgxpool.New(ctx, dsn)

	if err != nil {
		slog.Error("failed to connect to database", "err", err)
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		slog.Error("failed to ping database", "err", err)
		return nil, err
	}

	return &Database{dsn: dsn, pool: pool}, nil
}
