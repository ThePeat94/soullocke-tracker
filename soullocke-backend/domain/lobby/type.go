package lobby

import (
	"context"

	"github.com/google/uuid"
)

type Lobby struct {
	ID            uuid.UUID
	Name          string
	GameEditionID uint16
}

type LobbyWithCredentials struct {
	Lobby
	Password string
}

type LobbyRepository interface {
	CreateLobby(ctx context.Context, name string, password string, gameEditionID uint16) (*Lobby, error)
	GetLobby(ctx context.Context, id string) (*Lobby, error)
}
