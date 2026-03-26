package http

import (
	"context"
	"log/slog"
	"net/http"
	"soullocke-backend/domain/game_edition"

	"github.com/danielgtaylor/huma/v2"
)

type GameEditionsController struct {
	ger game_edition.GameEditionRepository
}

type GameEditionDto struct {
	ID         string `json:"id" example:"firered" doc:"Game Edition ID" format:"string"`
	Name       string `json:"name" example:"Firered" doc:"Game Edition Name" format:"string"`
	Generation uint64 `json:"generation" example:"3" doc:"Game Edition Generation"`
}

type GetGameEditionsOutput struct {
	Body []GameEditionDto
}

func (controller *GameEditionsController) RegisterRoutes(s *Server) {
	huma.Register(s.api, huma.Operation{
		OperationID:   "get-editions",
		Method:        http.MethodGet,
		Path:          "/editions",
		Summary:       "Retrieve all available Pokémon editions",
		Description:   "Retrieve all available Pokémon editions which the player can choose from to start a soullink run in",
		Tags:          []string{"GameEdition"},
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, i *struct{}) (*GetGameEditionsOutput, error) {
		gameEditions, err := controller.ger.GetGameEditions(ctx)
		if err != nil {
			slog.Error("error retrieving game editions", "error", err)
			return nil, huma.Error500InternalServerError("getting game editions: Internal Server Error", err)
		}

		output := &GetGameEditionsOutput{}
		outputEditions := make([]GameEditionDto, 0, len(gameEditions))
		for _, gameEdition := range gameEditions {
			outputEditions = append(outputEditions, GameEditionDto{
				ID:         gameEdition.ID,
				Name:       gameEdition.Name,
				Generation: gameEdition.Generation,
			})
		}
		output.Body = outputEditions
		return output, nil
	})
}
