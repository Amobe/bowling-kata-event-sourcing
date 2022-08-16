package out

import (
	"context"

	"github.com/amobe/bowling-kata-event-sourcing/src/bowling_game/entity"
)

type BowlingGameRepository interface {
	Save(ctx context.Context, bowlingGame *entity.BowlingGame) error
	FindByID(ctx context.Context, gameID string) (*entity.BowlingGame, error)
}
