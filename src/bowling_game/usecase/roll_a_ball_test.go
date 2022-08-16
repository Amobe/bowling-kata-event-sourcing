package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/amobe/bowling-kata-event-sourcing/src/bowling_game/adapter"
	"github.com/amobe/bowling-kata-event-sourcing/src/bowling_game/entity"
	"github.com/amobe/bowling-kata-event-sourcing/src/bowling_game/usecase/port/in"
	"github.com/amobe/bowling-kata-event-sourcing/src/bowling_game/usecase/port/out"
	coreentity "github.com/amobe/bowling-kata-event-sourcing/src/core/entity"
	coreusecase "github.com/amobe/bowling-kata-event-sourcing/src/core/usecase"
	"github.com/amobe/bowling-kata-event-sourcing/src/utils/pubsub/inmem"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type RollABallUseCaseTestSuite struct {
	suite.Suite

	eventBus   coreusecase.DomainEventBus
	listener   *fakeRollABallListener
	repository out.BowlingGameRepository
}

func TestRollABallUseCaseTestSuite(t *testing.T) {
	suite.Run(t, &RollABallUseCaseTestSuite{})
}

func (s *RollABallUseCaseTestSuite) SetupTest() {
	ctx := context.Background()

	s.listener = &fakeRollABallListener{}

	eventBus := inmem.NewDomainEventBus[coreentity.DomainEvent]()
	subscriber := inmem.NewSubscriber("event-listener", s.listener.Notifier)
	eventBus.Register(ctx, subscriber)
	s.eventBus = eventBus

	s.repository = adapter.NewInMemBowlingGameRepository(eventBus)
}

func (s *RollABallUseCaseTestSuite) TestRollABallUseCase() {
	gameID := s.createBowlingGame()

	ctx := context.Background()
	rollABallUseCase := NewRollABallServiceService(s.repository)
	input := in.RollABallInput{
		GameID: gameID,
		Hit:    1,
	}

	gotGameID, err := rollABallUseCase.execute(ctx, input)
	s.NoError(err)
	s.Equal(gameID, gotGameID)
	_, err = s.repository.FindByID(ctx, gotGameID)
	s.NoError(err)

	s.Eventually(func() bool {
		return s.Equal(1, s.listener.count)
	}, 5*time.Millisecond, 1*time.Millisecond)
}

func (s *RollABallUseCaseTestSuite) TestRollTwentyBallHasTwentyScore() {
	gameID := s.createBowlingGame()

	ctx := context.Background()
	rollABallUseCase := NewRollABallServiceService(s.repository)
	input := in.RollABallInput{
		GameID: gameID,
		Hit:    1,
	}

	for i := 0; i < 20; i++ {
		_, _ = rollABallUseCase.execute(ctx, input)
	}
	output, err := s.repository.FindByID(ctx, gameID)
	s.NoError(err)
	s.Equal(20, output.ThrowingCount())
	s.Equal(20, output.Score())
}

func (s *RollABallUseCaseTestSuite) TestRollTwentyOneBallHasTwentyScore() {
	gameID := s.createBowlingGame()

	ctx := context.Background()
	rollABallUseCase := NewRollABallServiceService(s.repository)
	input := in.RollABallInput{
		GameID: gameID,
		Hit:    1,
	}

	for i := 0; i < 21; i++ {
		_, _ = rollABallUseCase.execute(ctx, input)
	}
	output, err := s.repository.FindByID(ctx, gameID)
	s.NoError(err)
	s.Equal(20, output.ThrowingCount())
	s.Equal(20, output.Score())
}

func (s *RollABallUseCaseTestSuite) TestRollABallShouldNotHasNegativeHitValue() {
	gameID := s.createBowlingGame()

	ctx := context.Background()
	rollABallUseCase := NewRollABallServiceService(s.repository)
	input := in.RollABallInput{
		GameID: gameID,
		Hit:    -1,
	}

	_, _ = rollABallUseCase.execute(ctx, input)
	output, err := s.repository.FindByID(ctx, gameID)
	s.NoError(err)
	s.Equal(0, output.ThrowingCount())
	s.Equal(0, output.Score())
}

func (s *RollABallUseCaseTestSuite) TestRollABallWithSpareHasNoLeavingPinsAndNotBonusRemainHit() {
	gameID := s.createBowlingGame()
	s.hasSpare(gameID)

	ctx := context.Background()
	output, err := s.repository.FindByID(ctx, gameID)
	s.NoError(err)
	s.Equal(0, output.LeavingPins())
	s.Equal(0, output.BonusRemainHit())
}

func (s *RollABallUseCaseTestSuite) TestRollABallAfterSpareHasBonusValueAndLeavingPinsIsRestored() {
	gameID := s.createBowlingGame()
	s.hasSpare(gameID)

	ctx := context.Background()
	rollABallUseCase := NewRollABallServiceService(s.repository)
	_, _ = rollABallUseCase.execute(ctx, in.RollABallInput{
		GameID: gameID,
		Hit:    1,
	})
	output, err := s.repository.FindByID(ctx, gameID)
	s.NoError(err)
	s.Equal(12, output.Score())
	s.Equal(9, output.LeavingPins())
}

func (s *RollABallUseCaseTestSuite) createBowlingGame() string {
	createBowlingGameUseCase := NewCreateBowlingGameService(s.repository)
	gameID := uuid.New().String()
	input := in.CreateBowlingGameInput{
		GameID: gameID,
	}

	_, _ = createBowlingGameUseCase.execute(context.Background(), input)
	return gameID
}

func (s *RollABallUseCaseTestSuite) hasSpare(gameID string) {
	ctx := context.Background()
	rollABallUseCase := NewRollABallServiceService(s.repository)
	_, _ = rollABallUseCase.execute(ctx, in.RollABallInput{
		GameID: gameID,
		Hit:    1,
	})
	_, _ = rollABallUseCase.execute(ctx, in.RollABallInput{
		GameID: gameID,
		Hit:    9,
	})
}

type fakeRollABallListener struct {
	count int
}

func (l *fakeRollABallListener) Notifier(domainEvent coreentity.DomainEvent) {
	switch (domainEvent).(type) {
	case entity.BowlingGameRolledABall:
		l.count++
	}
}
