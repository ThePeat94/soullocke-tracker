package language

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"soullocke-backend/db"
)

var (
	ErrNotFound = errors.New("language: not found")
)

type Repository struct {
	*db.BaseRepository
}

func NewRepository(database *db.Database) *Repository {
	return &Repository{
		BaseRepository: db.NewBaseRepository(database),
	}
}

func (r *Repository) GetLanguageByName(ctx context.Context, name string) (*Language, error) {
	q := r.QueriesFromContext(ctx)
	language, err := q.FindLanguageByName(ctx, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("language: failed to find language by id: %w", err)
	}
	return &Language{
		ID:      uint16(language.ID),
		ISO639:  language.Iso639,
		ISO3166: language.Iso3166,
	}, nil
}

func (r *Repository) GetNamesForGameEditions(ctx context.Context, langIds []uint16) (map[uint16][]*LocalizedName, error) {
	q := r.QueriesFromContext(ctx)
	int16Ids := make([]int16, len(langIds))
	for i, id := range langIds {
		int16Ids[i] = int16(id)
	}
	rows, err := q.GetLocalizedNamesForGameEditions(ctx, int16Ids)
	if err != nil {
		return nil, fmt.Errorf("language: failed to get localized names: %w", err)
	}

	grouped := make(map[uint16][]*LocalizedName)

	for _, row := range rows {
		grouped[row.GameEditionID] = append(grouped[row.GameEditionID], &LocalizedName{
			Name: row.EditionName,
			Lang: row.LanguageName,
		})
	}

	return grouped, nil
}
