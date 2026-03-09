package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"soullocke-backend/domain/lobby"
	"strconv"
)

type Server struct {
	port uint16
	mux  *http.ServeMux
	lr   lobby.LobbyRepository
}

func NewServer(port uint16, lr lobby.LobbyRepository) *Server {
	mux := http.NewServeMux()
	return &Server{
		port: port,
		mux:  mux,
		lr:   lr,
	}
}

func (s *Server) Start() error {
	err := s.registerLobbyRoutes()

	if err != nil {
		return fmt.Errorf("http: failed to register lobby routes: %w", err)
	}

	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.mux)
}

type LobbyCreationRequest struct {
	Name          string `json:"name"`
	Password      string `json:"password"`
	GameEditionId string `json:"gameEditionId"`
}

type LobbyCreationResponse struct {
	LobbyId string `json:"lobbyId"`
}

func (s *Server) registerLobbyRoutes() error {

	s.mux.HandleFunc("POST /lobby", func(writer http.ResponseWriter, request *http.Request) {
		var creationRequest LobbyCreationRequest
		decoder := json.NewDecoder(request.Body)
		err := decoder.Decode(&creationRequest)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		l, err := s.lr.CreateLobby(context.Background(), creationRequest.Name, creationRequest.Password, creationRequest.GameEditionId)

		creationResponse := LobbyCreationResponse{
			LobbyId: l.ID,
		}
		encoded, err := json.Marshal(creationResponse)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.Header().Set("Content-Length", strconv.Itoa(len(encoded)))
		writer.WriteHeader(http.StatusCreated)
		writer.Write(encoded)
	})

	return nil
}
