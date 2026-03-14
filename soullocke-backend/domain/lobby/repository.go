package lobby

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"soullocke-backend/db"
	"soullocke-backend/db/gen"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Repository struct {
	db.BaseRepository
}

func NewRepository(database db.Database) *Repository {
	return &Repository{
		BaseRepository: *db.NewBaseRepository(database),
	}
}

func (r *Repository) GetLobbies(ctx context.Context) ([]*Lobby, error) {
	q := r.QueriesFromContext(ctx)
	storedLobbies, err := q.GetLobbies(ctx)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("lobby: failed to get lobbies: %w", err)
		}
		return []*Lobby{}, nil
	}

	var lobbies []*Lobby
	for _, l := range storedLobbies {
		lobbies = append(lobbies, toDomainLobby(l))
	}

	return lobbies, nil

}

func (r *Repository) GetLobby(ctx context.Context, id string) (*Lobby, error) {
	parsedUuid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("lobby: failed to parse lobby uuid: %w", err)
	}
	q := r.QueriesFromContext(ctx)
	storedLobby, err := q.GetLobby(ctx, parsedUuid)
	if err != nil {
		return nil, fmt.Errorf("lobby: failed to get lobby: %w", err)
	}

	return toDomainLobby(storedLobby), nil
}

func (r *Repository) CreateLobby(ctx context.Context, name string, password string, gameEditionID string) (*Lobby, error) {
	var domainLobby *Lobby
	err := r.Tx(ctx, func(ctx context.Context) error {
		creationArgs := toCreateLobbyArgs(name, password, gameEditionID)
		q := r.QueriesFromContext(ctx)
		lobby, err := q.CreateLobby(ctx, *creationArgs)
		if err != nil {
			return fmt.Errorf("lobby: failed to create lobby: %w", err)
		}
		domainLobby = toDomainLobby(lobby)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("lobby: failed to create lobby: %w", err)
	}

	return domainLobby, nil
}

func (r *Repository) UpdateLobby(ctx context.Context, id string, name, password string) (*Lobby, error) {
	args, err := toUpdateLobbyArgs(id, name, password)
	if err != nil {
		return nil, fmt.Errorf("lobby: failed to create args for lobby update: %w", err)
	}
	var domainLobby *Lobby
	err = r.Tx(ctx, func(ctx context.Context) error {
		q := r.QueriesFromContext(ctx)
		lobby, uErr := q.UpdateLobby(ctx, *args)
		if uErr != nil {
			return fmt.Errorf("lobby: failed to update lobby: %w", uErr)
		}
		domainLobby = toDomainLobby(lobby)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("lobby: failed to update lobby: %w", err)
	}

	return domainLobby, nil
}

func toDomainLobby(l dbgen.Lobby) *Lobby {
	return &Lobby{
		ID:            l.ID.String(),
		Name:          l.Name,
		Password:      l.Password,
		GameEditionID: l.GameEditionID.String,
	}
}

func toCreateLobbyArgs(name string, password string, gameEditionID string) *dbgen.CreateLobbyParams {
	return &dbgen.CreateLobbyParams{
		Name:          name,
		Password:      password,
		GameEditionID: pgtype.Text{String: gameEditionID, Valid: true},
	}
}

func toUpdateLobbyArgs(id string, name string, password string) (*dbgen.UpdateLobbyParams, error) {
	idUuid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("lobby: failed to parse uuid: %w", err)
	}

	return &dbgen.UpdateLobbyParams{
		ID:       idUuid,
		Name:     name,
		Password: password,
	}, nil
}
