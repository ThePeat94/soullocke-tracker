package http

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"soullocke-backend/domain/game_edition"
	"soullocke-backend/domain/lobby"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"gopkg.in/yaml.v3"
)

type Server struct {
	port                  uint16
	allowedOrigins        []string
	mux                   *http.ServeMux
	api                   huma.API
	lobbyController       *LobbyController
	gameEditionController *GameEditionsController
}

func NewServer(port uint16, allowedOrigins []string, lr lobby.LobbyRepository, ger game_edition.GameEditionRepository) *Server {
	mux := http.NewServeMux()
	api := humago.New(mux, huma.DefaultConfig("SoulLocker API", "0.0.1"))
	return &Server{
		port:           port,
		mux:            mux,
		api:            api,
		allowedOrigins: allowedOrigins,
		lobbyController: &LobbyController{
			lr: lr,
		},
		gameEditionController: &GameEditionsController{
			ger: ger,
		},
	}
}

func (s *Server) Setup() {
	s.lobbyController.RegisterLobbyRoutes(s)
	s.gameEditionController.RegisterRoutes(s)
}

func (s *Server) Serve(ctx context.Context) error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: corsMiddleware(s.mux, s.allowedOrigins),
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

func corsMiddleware(next http.Handler, allowedOrigins []string) http.Handler {
	allowed := make(map[string]bool, len(allowedOrigins))
	for _, o := range allowedOrigins {
		allowed[o] = true
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if allowed[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
