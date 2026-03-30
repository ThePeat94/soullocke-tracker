package lobby

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"soullocke-backend/db"
	"soullocke-backend/db/gen"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrNotFound = errors.New("lobby: not found")
)

type Repository struct {
	*db.BaseRepository
}

func NewRepository(database *db.Database) *Repository {
	return &Repository{
		BaseRepository: db.NewBaseRepository(database),
	}
}

func (r *Repository) GetLobby(ctx context.Context, id string) (*Lobby, error) {
	parsedUuid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("lobby: failed to parse lobby uuid: %w", err)
	}
	q := r.QueriesFromContext(ctx)
	storedLobby, err := q.GetLobby(ctx, parsedUuid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("lobby: failed to get lobby: %w", err)
	}

	return toDomainLobby(storedLobby), nil
}

func (r *Repository) CreateLobby(ctx context.Context, name string, password string, gameEditionID uint16) (*Lobby, error) {
	var domainLobby *Lobby
	err := r.Tx(ctx, func(ctx context.Context) error {
		creationArgs, cErr := toCreateLobbyArgs(name, password, gameEditionID)
		if cErr != nil {
			return cErr
		}
		q := r.QueriesFromContext(ctx)
		lobby, cErr := q.CreateLobby(ctx, *creationArgs)
		if cErr != nil {
			return cErr
		}
		domainLobby = toDomainLobbyFromCreation(lobby)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("lobby: failed to create lobby: %w", err)
	}

	return domainLobby, nil
}

func toDomainLobby(l dbgen.GetLobbyRow) *Lobby {
	return &Lobby{
		ID:            l.ID.String(),
		Name:          l.Name,
		GameEditionID: uint16(l.GameEditionID),
	}
}

func toDomainLobbyFromCreation(l dbgen.CreateLobbyRow) *Lobby {
	return &Lobby{
		ID:            l.ID.String(),
		Name:          l.Name,
		GameEditionID: uint16(l.GameEditionID),
	}
}

func toCreateLobbyArgs(name string, password string, gameEditionID uint16) (*dbgen.CreateLobbyParams, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, fmt.Errorf("lobby: failed to hash password: %w", err)
	}
	return &dbgen.CreateLobbyParams{
		Name:          name,
		Password:      string(hashedPassword),
		GameEditionID: int32(gameEditionID),
	}, nil
}
