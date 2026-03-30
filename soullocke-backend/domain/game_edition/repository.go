package game_edition

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"soullocke-backend/db"
	dbgen "soullocke-backend/db/gen"
)

type Repository struct {
	*db.BaseRepository
}

func NewRepository(database *db.Database) *Repository {
	return &Repository{
		BaseRepository: db.NewBaseRepository(database),
	}
}

func (r *Repository) GetGameEditions(ctx context.Context) ([]*GameEdition, error) {
	q := r.QueriesFromContext(ctx)
	editions, err := q.GetGameEditions(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Warn("game edition: no game editions found")
			return []*GameEdition{}, nil
		}
		return nil, fmt.Errorf("game edition: getting game editions: %w", err)
	}

	domainEditions := make([]*GameEdition, 0, len(editions))
	for _, edition := range editions {
		domainEditions = append(domainEditions, toDomainGameEdition(edition))
	}

	return domainEditions, nil
}

func toDomainGameEdition(edition dbgen.GameEdition) *GameEdition {
	return &GameEdition{
		ID:         uint16(edition.ID),
		Name:       edition.CodeName,
		Generation: uint64(edition.Generation),
	}
}
