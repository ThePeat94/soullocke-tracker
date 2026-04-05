package http

import (
	"context"
	"log/slog"
	"net/http"
	"soullocke-backend/domain/game_edition"
	"soullocke-backend/domain/language"

	"github.com/danielgtaylor/huma/v2"
)

type GameEditionsController struct {
	ger                game_edition.GameEditionRepository
	lr                 language.LanguageRepository
	supportedLanguages []language.Language
}

type GameEditionDto struct {
	ID           uint16                    `json:"id" example:"1" doc:"Game Edition ID"`
	Generation   uint64                    `json:"generation" example:"3" doc:"Game Edition Generation"`
	FallbackName string                    `json:"fallbackName" example:"firered" doc:"Fallback Name (Code Name) for the Game Edition"`
	Names        []*language.LocalizedName `json:"names" doc:"Game Edition Names"`
}

type GetGameEditionsOutput struct {
	Body []GameEditionDto
}

func NewGameEditionsController(ger game_edition.GameEditionRepository, supportedLanguages []language.Language, lr language.LanguageRepository) *GameEditionsController {
	return &GameEditionsController{
		ger:                ger,
		supportedLanguages: supportedLanguages,
		lr:                 lr,
	}
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
		editionNames, err := controller.lr.GetNamesForGameEditions(ctx, language.GetLanguageIds(controller.supportedLanguages))
		if err != nil || editionNames == nil {
			slog.Warn("error retrieving game edition names", "error", err)
		}
		for _, gameEdition := range gameEditions {
			genNames := editionNames[gameEdition.ID]
			if genNames == nil {
				genNames = []*language.LocalizedName{}
			}
			outputEditions = append(outputEditions, GameEditionDto{
				ID:           gameEdition.ID,
				Generation:   gameEdition.Generation,
				FallbackName: gameEdition.Name,
				Names:        genNames,
			})
		}
		output.Body = outputEditions
		return output, nil
	})
}
