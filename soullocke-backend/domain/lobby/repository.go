package lobby

import (
	"context"
	"errors"
	"fmt"
	db2 "soullocke-backend/db"
	db "soullocke-backend/db/gen"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type Repository struct {
	db2.CommonRepository
}

func NewRepository(db db2.Database) *Repository {
	return &Repository{
		CommonRepository: *db2.NewCommonRepository(db),
	}
}

func (r *Repository) GetLobbies(ctx context.Context) ([]*Lobby, error) {
	q := r.QueriesFromContext(ctx)
	storedLobbies, err := q.GetLobbies(ctx)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("lobby: failed to get lobbies: %w", err)
		}
		return []*Lobby{}, err
	}

	lobbies := make([]*Lobby, len(storedLobbies))
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
	err := r.Tx(ctx, func() error {
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
	args, err := toUpdateLoggyArgs(id, name, password)
	if err != nil {
		return nil, fmt.Errorf("lobby: failed to create args for lobby update: %w", err)
	}
	var domainLobby *Lobby
	err = r.Tx(ctx, func() error {
		q := r.QueriesFromContext(ctx)
		lobby, uErr := q.UpdateLobby(ctx, *args)
		if uErr != nil {
			return fmt.Errorf("lobby: failed to update lobby: %w", err)
		}
		domainLobby = toDomainLobby(lobby)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("lobby: failed to update lobby: %w", err)
	}

	return domainLobby, nil
}

func toDomainLobby(l db.Lobby) *Lobby {
	return &Lobby{
		ID:            l.ID.String(),
		Name:          l.Name,
		Password:      l.Password,
		GameEditionID: l.GameEditionID.String,
	}
}

func toCreateLobbyArgs(name string, password string, gameEditionID string) *db.CreateLobbyParams {
	return &db.CreateLobbyParams{
		Name:          name,
		Password:      password,
		GameEditionID: pgtype.Text{String: gameEditionID},
	}
}

func toUpdateLoggyArgs(id string, name string, password string) (*db.UpdateLobbyParams, error) {
	idUuid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("lobby: failed to parse uuid: %w", err)
	}

	return &db.UpdateLobbyParams{
		ID:       idUuid,
		Name:     name,
		Password: password,
	}, nil
}
