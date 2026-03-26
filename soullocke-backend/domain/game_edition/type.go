package game_edition

import (
	"context"
)

type GameEdition struct {
	ID         string
	Name       string
	Generation uint64
}

type GameEditionRepository interface {
	GetGameEditions(ctx context.Context) ([]*GameEdition, error)
	CreateGameEdition(ctx context.Context, id string, name string, generation uint64) (*GameEdition, error)
}
