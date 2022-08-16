package entity

import core "github.com/amobe/bowling-kata-event-sourcing/src/core/entity"

type BowlingGameEvents struct {
	core.DomainEvent
	name string
}

func (e BowlingGameEvents) String() string {
	return e.name
}

type BowlingGameCreated struct {
	BowlingGameEvents
	gameID string
}

func NewBowlingGameCreated(gameID string) BowlingGameCreated {
	return BowlingGameCreated{
		BowlingGameEvents: BowlingGameEvents{
			name: "NewBowlingGameCreated",
		},
		gameID: gameID,
	}
}
