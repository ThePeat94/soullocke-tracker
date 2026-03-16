package lobby

import "context"

type Lobby struct {
	ID            string
	Name          string
	GameEditionID string
}

type LobbyWithCredentials struct {
	Lobby
	Password string
}

type LobbyRepository interface {
	CreateLobby(ctx context.Context, name string, password string, gameEditionID string) (*Lobby, error)
	GetLobbies(ctx context.Context) ([]*Lobby, error)
	GetLobby(ctx context.Context, id string) (*Lobby, error)
	UpdateLobby(ctx context.Context, id string, name, password string) (*Lobby, error)
}
