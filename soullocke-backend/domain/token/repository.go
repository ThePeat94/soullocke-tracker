package token

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"soullocke-backend/db"
	dbgen "soullocke-backend/db/gen"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrNotFound = errors.New("token: not found")
)

type Repository struct {
	*db.BaseRepository
}

func NewRepository(database *db.Database) *Repository {
	return &Repository{
		BaseRepository: db.NewBaseRepository(database),
	}
}

func (r *Repository) GetTokenByHash(ctx context.Context, hash []byte) (*Token, error) {
	q := r.QueriesFromContext(ctx)
	token, err := q.FindTokenByHash(ctx, hash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("token: failed to get token by hash: %w", err)
	}

	return toDomainToken(token), nil
}

func (r *Repository) CreateToken(ctx context.Context, hash []byte, lobbyID uuid.UUID, expiresAt time.Time) (*Token, error) {
	var domainToken *Token
	err := r.Tx(ctx, func(ctx context.Context) error {
		createTokenParams := toCreateTokenParams(hash, lobbyID, expiresAt)
		q := r.QueriesFromContext(ctx)
		token, err := q.CreateToken(ctx, createTokenParams)
		if err != nil {
			return err
		}
		domainToken = toDomainToken(token)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("token: failed to create token: %w", err)
	}

	return domainToken, nil
}

func (r *Repository) DeleteTokenByHash(ctx context.Context, hash []byte) error {
	q := r.QueriesFromContext(ctx)
	err := q.DeleteTokenByHash(ctx, hash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFound
		}
		return fmt.Errorf("token: failed to delete token by hash: %w", err)
	}
	return nil
}

func toDomainToken(token dbgen.TokenAuth) *Token {
	return &Token{
		ID:        token.ID.String(),
		Token:     token.TokenHash,
		LobbyId:   token.LobbyID.String(),
		CreatedAt: token.CreatedAt.Time,
		ExpiresAt: token.ExpiresAt.Time,
	}
}

func toCreateTokenParams(hash []byte, lobbyID uuid.UUID, expiresAt time.Time) dbgen.CreateTokenParams {
	return dbgen.CreateTokenParams{
		TokenHash: hash,
		LobbyID:   lobbyID,
		ExpiresAt: pgtype.Timestamptz{Time: expiresAt, Valid: true},
	}
}
