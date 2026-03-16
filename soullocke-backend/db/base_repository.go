package db

import (
	"context"
	"fmt"
	"soullocke-backend/db/gen"

	"github.com/jackc/pgx/v5"
)

type contextKey struct{}

var txKey = contextKey{}

type BaseRepository struct {
	db *Database
}

func NewBaseRepository(db *Database) *BaseRepository {
	return &BaseRepository{
		db: db,
	}
}

func (r *BaseRepository) QueriesFromContext(ctx context.Context) dbgen.Queries {
	if tx, ok := ctx.Value(txKey).(pgx.Tx); ok {
		return *dbgen.New(tx)
	}

	return *dbgen.New(r.db.pool)
}

func (r *BaseRepository) Tx(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := r.db.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("db: failed to start transaction: %w", err)
	}
	txCtx := context.WithValue(ctx, txKey, tx)
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	err = fn(txCtx)
	if err != nil {
		return fmt.Errorf("db: transaction failed: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("db: failed to commit transaction: %w", err)
	}

	return nil
}
