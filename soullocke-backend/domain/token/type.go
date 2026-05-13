package token

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Token struct {
	ID        string
	Token     []byte
	LobbyId   string
	CreatedAt time.Time
	ExpiresAt time.Time
}

type TokenRepository interface {
	GetTokenByHash(ctx context.Context, hash []byte) (*Token, error)
	CreateToken(ctx context.Context, hash []byte, lobbyID uuid.UUID, expiresAt time.Time) (*Token, error)
	DeleteTokenByHash(ctx context.Context, hash []byte) error
}
