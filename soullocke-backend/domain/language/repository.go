package language

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"soullocke-backend/db"
	dbgen "soullocke-backend/db/gen"
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

func (r *Repository) GetNamesForGameEdition(ctx context.Context, gameEditionID uint16, langIds []uint16) ([]*LocalizedName, error) {
	q := r.QueriesFromContext(ctx)
	int16Ids := make([]int16, len(langIds))
	for i, id := range langIds {
		int16Ids[i] = int16(id)
	}
	names, err := q.GetLocalizedNamesForGameEdition(ctx, dbgen.GetLocalizedNamesForGameEditionParams{
		GameEditionID: gameEditionID,
		LanguageIds:   int16Ids,
	})

	if err != nil {
		return nil, fmt.Errorf("language: failed to get localized names: %w", err)
	}

	localizedNames := make([]*LocalizedName, 0, len(names))
	for _, name := range names {
		localizedNames = append(localizedNames, toGameEditionLocalizedNames(name))
	}

	return localizedNames, nil
}

func toGameEditionLocalizedNames(name dbgen.GetLocalizedNamesForGameEditionRow) *LocalizedName {
	return &LocalizedName{
		Name: name.EditionName,
		Lang: name.LanguageName,
	}
}
