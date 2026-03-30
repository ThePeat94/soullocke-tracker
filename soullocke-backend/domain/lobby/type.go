package lobby

import "context"

type Lobby struct {
	ID            string
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
