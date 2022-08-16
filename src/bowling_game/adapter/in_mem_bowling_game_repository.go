package adapter

import (
	"context"
	"fmt"

	"github.com/amobe/bowling-kata-event-sourcing/src/bowling_game/entity"
	"github.com/amobe/bowling-kata-event-sourcing/src/bowling_game/usecase/port/out"
	"github.com/amobe/bowling-kata-event-sourcing/src/core/usecase"
)

type inMemBowlingGameRepository struct {
	eventBus usecase.DomainEventBus
	store    map[string]*entity.BowlingGame
}

func NewInMemBowlingGameRepository(eventBus usecase.DomainEventBus) out.BowlingGameRepository {
	return &inMemBowlingGameRepository{
		eventBus: eventBus,
		store:    map[string]*entity.BowlingGame{},
	}
}

func (r *inMemBowlingGameRepository) Save(ctx context.Context, bowlingGame *entity.BowlingGame) error {
	r.store[bowlingGame.GameID()] = bowlingGame
	if err := r.eventBus.PostAll(ctx, bowlingGame.AggregateRoot); err != nil {
		return fmt.Errorf("event bus post all bowling game: %w", err)
	}
	return nil
}

func (r *inMemBowlingGameRepository) FindByID(ctx context.Context, gameID string) (*entity.BowlingGame, error) {
	g, ok := r.store[gameID]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return g, nil
}
