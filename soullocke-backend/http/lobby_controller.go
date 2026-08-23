package http

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"soullocke-backend/domain/lobby"
	"soullocke-backend/domain/token"
	"strings"

	"github.com/danielgtaylor/huma/v2"
)

type LobbyController struct {
	lr lobby.LobbyRepository
	tr token.TokenRepository
}

type LobbyCreationRequest struct {
	Name          string `json:"name" minLength:"10" maxLength:"255" example:"Weekend Lobby"`
	Password      string `json:"password" minLength:"8" maxLength:"255" example:"f00b4r1234"`
	GameEditionId uint16 `json:"gameEditionId" minimum:"1" example:"8" doc:"The game edition id of the pokemon version"`
}

type LobbyCreationResponse struct {
	LobbyId string `json:"lobbyId" example:"73beb67c-70c5-4c95-b99e-73e3c076f82f" doc:"Lobby ID as a UUID"`
}

type GetLobbyResponse struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	GameEditionId uint16 `json:"gameEditionId"`
}

type CreateLobbyInput struct {
	Body LobbyCreationRequest
}

type CreateLobbyOutput struct {
	Body LobbyCreationResponse
}

type GetLobbyOutput struct {
	Body GetLobbyResponse
}

type GetLobbyInput struct {
	ID string `path:"lobbyId" example:"73beb67c-70c5-4c95-b99e-73e3c076f82f" doc:"Lobby ID as a UUID" format:"uuid"`
}

func NewLobbyController(lr lobby.LobbyRepository, tr token.TokenRepository) *LobbyController {
	return &LobbyController{
		lr: lr,
		tr: tr,
	}
}

func (lc *LobbyController) RegisterLobbyRoutes(s *Server) {

	huma.Register(s.api, huma.Operation{
		OperationID:   "create-lobby",
		Method:        http.MethodPost,
		Path:          "/lobby",
		Summary:       "Creates a new lobby",
		Description:   "Creates a new lobby for players to join and manage their soullink run",
		Tags:          []string{"Lobby"},
		DefaultStatus: http.StatusCreated,
	}, func(ctx context.Context, i *CreateLobbyInput) (*CreateLobbyOutput, error) {

		resp := &CreateLobbyOutput{}
		l, err := lc.lr.CreateLobby(ctx, i.Body.Name, i.Body.Password, i.Body.GameEditionId)
		if err != nil {
			slog.Error("error creating lobby", "error", err)
			return resp, huma.Error500InternalServerError("creating lobby: Internal Server Error", err)
		}
		resp.Body = LobbyCreationResponse{l.ID.String()}
		return resp, nil
	})

	huma.Register(s.api, huma.Operation{
		OperationID:   "get-lobby",
		Method:        http.MethodGet,
		Path:          "/lobby/{lobbyId}",
		Summary:       "Retrieve a lobby",
		Description:   "Retrieve a lobby in which players manage their soullink run",
		Tags:          []string{"Lobby"},
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, i *GetLobbyInput) (*GetLobbyOutput, error) {
		l, err := lc.lr.GetLobby(ctx, i.ID)
		if err != nil {
			if errors.Is(err, lobby.ErrNotFound) {
				slog.Warn("lobby not found", "id", i.ID)
				return nil, huma.Error404NotFound("getting lobby: Lobby is not existing")
			}
			slog.Error("error getting lobby", "error", err)
			return nil, huma.Error500InternalServerError("getting lobby: Internal Server Error", err)
		}

		response := GetLobbyResponse{
			ID:            l.ID.String(),
			Name:          l.Name,
			GameEditionId: l.GameEditionID,
		}
		return &GetLobbyOutput{Body: response}, nil
	})
}

func (b *LobbyCreationRequest) Resolve(ctx huma.Context) []error {
	b.Name = strings.TrimSpace(b.Name)
	b.Password = strings.TrimSpace(b.Password)

	var errs []error

	if n := len(b.Name); n < 10 || n > 255 {
		errs = append(errs, fmt.Errorf("name length must be between 10 and 255"))
	}
	if n := len(b.Password); n < 8 || n > 255 {
		errs = append(errs, fmt.Errorf("password length must be between 8 and 255"))
	}
	if b.GameEditionId == 0 {
		errs = append(errs, fmt.Errorf("a game edition must be set"))
	}

	return errs
}
