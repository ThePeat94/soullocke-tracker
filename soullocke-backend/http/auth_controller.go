package http

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"log/slog"
	"net/http"
	"soullocke-backend/domain/lobby"
	"soullocke-backend/domain/token"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"golang.org/x/crypto/bcrypt"
)

const lobbyAccessCookieName = "soullocker_lobby_access"

type AuthController struct {
	tr token.TokenRepository
	lr lobby.LobbyRepository
}

type AuthInput struct {
	Body AuthRequest
}

type AuthRequest struct {
	Password string `json:"password" example:"123456" doc:"The password of the lobby to login to"`
	LobbyID  string `json:"lobbyId" example:"73beb67c-70c5-4c95-b99e-73e3c076f82f" doc:"Lobby ID as a UUID"`
}

type AuthOutput struct {
	SetCookie http.Cookie `header:"Set-Cookie"`

	Body AuthResponse
}

type AuthResponse struct {
	ExpiresAt time.Time `json:"expiresAt" example:"2022-01-01T00:00:00+00:00"`
}

type AccessInput struct {
	AccessToken string `cookie:"soullocker_lobby_access"`

	LobbyID string `path:"lobbyId" example:"73beb67c-70c5-4c95-b99e-73e3c076f82f" doc:"Lobby ID as a UUID" format:"uuid"`
}

type AccessRequest struct {
	LobbyID string `json:"lobbyId" example:"73beb67c-70c5-4c95-b99e-73e3c076f82f" doc:"Lobby ID as a UUID"`
}

type AccessOutput struct {
	Body AccessResponse
}

type AccessResponse struct {
	CanWrite bool `json:"canWrite" example:"true" doc:"Whether or not the user can modify the lobby"`
}

func NewAuthController(tr token.TokenRepository, lr lobby.LobbyRepository) *AuthController {
	return &AuthController{
		tr: tr,
		lr: lr,
	}
}

func (lc *AuthController) RegisterAuthRoutes(s *Server) {
	huma.Register(s.api, huma.Operation{
		OperationID:   "login",
		Method:        http.MethodPost,
		Path:          "/login",
		Summary:       "Login to a lobby",
		Description:   "Login to a lobby with the password",
		Tags:          []string{"Login"},
		DefaultStatus: http.StatusOK,
	}, lc.handleLogin())

	huma.Register(s.api, huma.Operation{
		OperationID:   "check-access",
		Method:        http.MethodGet,
		Path:          "/access/{lobbyId}",
		Summary:       "Cookie access check",
		Description:   "Check if the access cookie is valid for a given lobby",
		Tags:          []string{"Login"},
		DefaultStatus: http.StatusOK,
	}, lc.handleAccessCheck())
}

func (lc *AuthController) handleLogin() func(ctx context.Context, i *AuthInput) (*AuthOutput, error) {
	return func(ctx context.Context, i *AuthInput) (*AuthOutput, error) {
		foundLobby, err := lc.lr.GetLobby(ctx, i.Body.LobbyID)
		if err != nil {
			if errors.Is(err, lobby.ErrNotFound) {
				return nil, huma.Error404NotFound("lobby not found")
			}
			slog.Error("error loading lobby", "error", err)
			return nil, huma.Error500InternalServerError("failed to load lobby")
		}

		err = bcrypt.CompareHashAndPassword(
			[]byte(foundLobby.Password),
			[]byte(i.Body.Password),
		)

		if err != nil {
			return nil, huma.Error401Unauthorized("unauthorized")
		}

		rawToken, err := generateAccessToken()
		if err != nil {
			return nil, huma.Error500InternalServerError("failed to create access token")
		}

		tokenHash := hashAccessToken(rawToken)
		expiresAt := time.Now().Add(6 * time.Hour)

		_, err = lc.tr.CreateToken(ctx, tokenHash, foundLobby.ID, expiresAt)
		if err != nil {
			return nil, huma.Error500InternalServerError("failed to create token")
		}

		out := &AuthOutput{
			SetCookie: http.Cookie{
				Name:     lobbyAccessCookieName,
				Value:    rawToken,
				Path:     "/",
				HttpOnly: true,
				Secure:   false,
				SameSite: http.SameSiteLaxMode,
				Expires:  expiresAt,
				MaxAge:   int(time.Until(expiresAt).Seconds()),
			},
			Body: AuthResponse{
				ExpiresAt: expiresAt,
			},
		}

		return out, nil
	}
}

func (lc *AuthController) handleAccessCheck() func(ctx context.Context, i *AccessInput) (*AccessOutput, error) {
	return func(ctx context.Context, i *AccessInput) (*AccessOutput, error) {
		_, err := lc.lr.GetLobby(ctx, i.LobbyID)
		if err != nil {
			if errors.Is(err, lobby.ErrNotFound) {
				return nil, huma.Error404NotFound("lobby not found")
			}
			slog.Error("error loading lobby", "error", err)
			return nil, huma.Error500InternalServerError("failed to load lobby")
		}

		tokenHash := hashAccessToken(i.AccessToken)
		foundToken, err := lc.tr.GetTokenByHash(ctx, tokenHash)
		if err != nil {
			return &AccessOutput{Body: AccessResponse{CanWrite: false}}, huma.Error401Unauthorized("unauthorized")
		}

		if foundToken.LobbyId != i.LobbyID {
			return &AccessOutput{Body: AccessResponse{CanWrite: false}}, huma.Error401Unauthorized("unauthorized")
		}

		out := &AccessOutput{
			Body: AccessResponse{
				CanWrite: true,
			},
		}

		return out, nil
	}
}

func generateAccessToken() (string, error) {
	bytes := make([]byte, 32)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(bytes), nil
}

func hashAccessToken(rawToken string) []byte {
	sum := sha256.Sum256([]byte(rawToken))
	return sum[:]
}
