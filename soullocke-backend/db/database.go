package db

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	dsn  string
	pool *pgxpool.Pool
}

func NewDatabase(ctx context.Context, dsn string) (*Database, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("db: failed to connect to database: %w", err)
	}

	db := &Database{dsn: dsn, pool: pool}
	err = db.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("db: failed to ping database: %w", err)
	}

	return db, nil
}

func (db *Database) Pool() *pgxpool.Pool {
	return db.pool
}

func (db *Database) Close() {
	db.pool.Close()
}

func (db *Database) Ping(ctx context.Context) error {
	return db.pool.Ping(ctx)
}

func (db *Database) Migrate(ctx context.Context) error {
	m, err := migrate.New("file://db/migrations", db.dsn)
	defer m.Close()

	if err != nil {
		return fmt.Errorf("db: failed to connect to migrations: %w", err)
	}

	err = m.Up()
	if err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("db: failed to apply migrations: %w", err)
		}
		slog.Info("db: migrations not applied, since no change has been detected")
	}
	return nil
}
