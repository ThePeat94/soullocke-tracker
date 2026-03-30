package game_edition

import (
	"context"
)

type GameEdition struct {
	ID         uint16
	Name       string
	Generation uint64
}

type GameEditionRepository interface {
	GetGameEditions(ctx context.Context) ([]*GameEdition, error)
}
