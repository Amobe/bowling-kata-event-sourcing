package usecase

import (
	"context"
	"fmt"

	"github.com/amobe/bowling-kata-event-sourcing/src/bowling_game/entity"
	"github.com/amobe/bowling-kata-event-sourcing/src/bowling_game/usecase/port/in"
	"github.com/amobe/bowling-kata-event-sourcing/src/bowling_game/usecase/port/out"
)

type CreateBowlingGameService struct {
	repository out.BowlingGameRepository
}

func NewCreateBowlingGameService(repository out.BowlingGameRepository) *CreateBowlingGameService {
	return &CreateBowlingGameService{
		repository: repository,
	}
}

func (s *CreateBowlingGameService) execute(ctx context.Context, input in.CreateBowlingGameInput) (string, error) {
	bowlingGame := entity.NewBowlingGame(input.GameID)
	if err := s.repository.Save(ctx, bowlingGame); err != nil {
		return "", fmt.Errorf("repository save bowling game: %w", err)
	}
	return bowlingGame.GameID(), nil
}
