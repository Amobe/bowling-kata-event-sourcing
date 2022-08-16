package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/amobe/bowling-kata-event-sourcing/src/bowling_game/adapter"
	"github.com/amobe/bowling-kata-event-sourcing/src/bowling_game/usecase/port/in"
	"github.com/amobe/bowling-kata-event-sourcing/src/core/entity"
	"github.com/amobe/bowling-kata-event-sourcing/src/utils/pubsub/inmem"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateBowlingGameUseCase(t *testing.T) {
	ctx := context.Background()

	eventBus := inmem.NewDomainEventBus[entity.DomainEvent]()
	listener := &fakeCreateBowlingGameListener{}
	subscriber := inmem.NewSubscriber("event-listener", listener.Notifier)
	eventBus.Register(ctx, subscriber)

	repository := adapter.NewInMemBowlingGameRepository(eventBus)
	createBowlingGameUseCase := NewCreateBowlingGameService(repository)
	input := in.CreateBowlingGameInput{
		GameID: uuid.New().String(),
	}

	output, err := createBowlingGameUseCase.execute(ctx, input)
	assert.NoError(t, err)
	_, err = repository.FindByID(ctx, output)
	assert.NoError(t, err)

	assert.Eventually(t, func() bool {
		return assert.Equal(t, 1, listener.count)
	}, 5*time.Millisecond, 1*time.Millisecond)
}

type fakeCreateBowlingGameListener struct {
	count int
}

func (l *fakeCreateBowlingGameListener) Notifier(event entity.DomainEvent) {
	l.count++
}
