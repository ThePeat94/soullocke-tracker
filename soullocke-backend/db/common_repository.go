package db

import (
	"context"
	"fmt"
	queries "soullocke-backend/db/gen"

	"github.com/jackc/pgx/v5"
)

type CommonRepository struct {
	db Database
}

func NewCommonRepository(db Database) *CommonRepository {
	return &CommonRepository{
		db: db,
	}
}

func (r *CommonRepository) QueriesFromContext(ctx context.Context) queries.Queries {
	if tx, ok := ctx.Value("tx").(pgx.Tx); ok {
		return *queries.New(tx)
	}

	return *queries.New(r.db.pool)
}

func (r *CommonRepository) Tx(ctx context.Context, fn func() error) error {
	tx, err := r.db.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("db: failed to start transaction: %w", err)
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	err = fn()
	if err != nil {
		return fmt.Errorf("db: transaction failed: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("db: failed to commit transaction: %w", err)
	}

	return nil
}
