package http

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"soullocke-backend/domain/lobby"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"gopkg.in/yaml.v3"
)

type Server struct {
	port uint16
	mux  *http.ServeMux
	api  huma.API
	lr   lobby.LobbyRepository
}

type LobbyCreationRequest struct {
	Name          string `json:"name" minLength:"1" maxLength:"255" example:"Weekend Lobby"`
	Password      string `json:"password" minLength:"1" maxLength:"255" example:"f00b4r"`
	GameEditionId string `json:"gameEditionId" minLength:"1" maxLength:"255" example:"firered" doc:"The game edition id of the pokemon version"`
}

type LobbyCreationResponse struct {
	LobbyId string `json:"lobbyId" example:"73beb67c-70c5-4c95-b99e-73e3c076f82f" doc:"Lobby ID as a UUID"`
}

type GetLobbyResponse struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	GameEditionId string `json:"gameEditionId"`
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
	ID string `path:"id" example:"73beb67c-70c5-4c95-b99e-73e3c076f82f" doc:"Lobby ID as a UUID" format:"uuid"`
}

func NewServer(port uint16, lr lobby.LobbyRepository) *Server {
	mux := http.NewServeMux()
	api := humago.New(mux, huma.DefaultConfig("SoulLocker API", "0.0.1"))
	return &Server{
		port: port,
		mux:  mux,
		api:  api,
		lr:   lr,
	}
}

func (s *Server) Setup() {
	s.registerLobbyRoutes()
}

func (s *Server) Serve(ctx context.Context) error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: corsMiddleware(s.mux),
	}

	// Start server in background
	go func() {
		<-ctx.Done()
		slog.Info("shutting down http server")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_ = srv.Shutdown(shutdownCtx)
	}()

	err := srv.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}

func (s *Server) ExportOpenAPISpec(path string) error {
	spec, err := yaml.Marshal(s.api.OpenAPI())
	if err != nil {
		return fmt.Errorf("marshalling openapi spec: %w", err)
	}
	return os.WriteFile(path, spec, 0644)
}

func (s *Server) registerLobbyRoutes() {

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
		l, err := s.lr.CreateLobby(ctx, i.Body.Name, i.Body.Password, i.Body.GameEditionId)
		if err != nil {
			return resp, huma.Error500InternalServerError("creating lobby: Internal Server Error", err)
		}
		resp.Body = LobbyCreationResponse{l.ID}
		return resp, nil
	})

	huma.Register(s.api, huma.Operation{
		OperationID:   "get-lobby",
		Method:        http.MethodGet,
		Path:          "/lobby/{id}",
		Summary:       "Retrieve a lobby",
		Description:   "Retrieve a lobby in which players manage their soullink run",
		Tags:          []string{"Lobby"},
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, i *GetLobbyInput) (*GetLobbyOutput, error) {
		l, err := s.lr.GetLobby(ctx, i.ID)
		if err != nil {
			if errors.Is(err, lobby.ErrNotFound) {
				return nil, huma.Error404NotFound("getting lobby: Lobby is not existing")
			}
			return nil, huma.Error500InternalServerError("getting lobby: Internal Server Error", err)
		}

		response := GetLobbyResponse{
			ID:            l.ID,
			Name:          l.Name,
			GameEditionId: l.GameEditionID,
		}
		return &GetLobbyOutput{Body: response}, nil
	})
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
