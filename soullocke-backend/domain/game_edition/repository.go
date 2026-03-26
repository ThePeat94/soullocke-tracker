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

var (
	ErrNotFound = errors.New("game edition: not found")
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

func (r *Repository) CreateGameEdition(ctx context.Context, id string, name string, generation uint64) (*GameEdition, error) {
	var domainEdition *GameEdition
	err := r.Tx(ctx, func(ctx context.Context) error {
		params := toCreateGameEditionParams(id, name, generation)
		q := r.QueriesFromContext(ctx)
		edition, err := q.CreateGameEdition(ctx, params)
		if err != nil {
			return err
		}
		domainEdition = toDomainGameEdition(edition)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("game edition: creating game edition: %w", err)
	}

	return domainEdition, nil
}

func toDomainGameEdition(edition dbgen.GameEdition) *GameEdition {
	return &GameEdition{
		ID:         edition.ID,
		Name:       edition.Name,
		Generation: uint64(edition.Generation),
	}
}

func toCreateGameEditionParams(id string, name string, generation uint64) dbgen.CreateGameEditionParams {
	return dbgen.CreateGameEditionParams{
		ID:         id,
		Name:       name,
		Generation: int16(generation),
	}
}
