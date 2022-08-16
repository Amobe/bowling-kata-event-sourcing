package inmem

import (
	"context"
	"fmt"

	"github.com/amobe/bowling-kata-event-sourcing/src/core/entity"
	"github.com/amobe/bowling-kata-event-sourcing/src/core/usecase"
)

type domainEventBusAdapter[T entity.DomainEvent] struct {
	bus *hub[T]
}

func NewDomainEventBus[T entity.DomainEvent]() usecase.DomainEventBus {
	return &domainEventBusAdapter[T]{
		bus: NewPubSub[T](),
	}
}

func (d domainEventBusAdapter[T]) Post(ctx context.Context, domainEvent entity.DomainEvent) error {
	if err := d.bus.Publish(ctx, domainEvent); err != nil {
		return fmt.Errorf("bus publish event: %w", err)
	}
	return nil
}

func (d domainEventBusAdapter[T]) PostAll(ctx context.Context, aggregateRoot *entity.AggregateRoot) error {
	for _, e := range aggregateRoot.DomainEvents() {
		if err := d.bus.Publish(ctx, e); err != nil {
			return fmt.Errorf("bus publish event: %w", err)
		}
	}
	return nil
}

func (d domainEventBusAdapter[T]) Register(ctx context.Context, listener any) error {
	if err := d.bus.Subscribe(ctx, listener.(*subscriber[T])); err != nil {
		return fmt.Errorf("bus subscribe listener: %w", err)
	}
	return nil
}

func (d domainEventBusAdapter[T]) Unregister(ctx context.Context, listener any) error {
	if err := d.bus.Unsubscribe(ctx, listener.(*subscriber[T])); err != nil {
		return fmt.Errorf("bus unsubscribe listener: %w", err)
	}
	return nil
}
