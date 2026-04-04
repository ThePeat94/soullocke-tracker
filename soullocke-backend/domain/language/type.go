package language

import "context"

type Language struct {
	ID      uint16
	ISO639  string
	ISO3166 string
}

type LocalizedName struct {
	Lang string `json:"lang" example:"de"`
	Name string `json:"name" example:"Feuerrot"`
}

type LanguageRepository interface {
	GetLanguageByName(ctx context.Context, name string) (*Language, error)
	GetNamesForGameEdition(ctx context.Context, gameEditionID uint16, langIds []uint16) ([]*LocalizedName, error)
}
