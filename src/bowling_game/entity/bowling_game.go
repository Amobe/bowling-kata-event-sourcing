package entity

import (
	core "github.com/amobe/bowling-kata-event-sourcing/src/core/entity"
)

type BowlingGame struct {
	*core.AggregateRoot
	gameID        string
	score         int
	throwingCount int
}

func newBowlingGame() *BowlingGame {
	bowlingGame := &BowlingGame{}
	bowlingGame.AggregateRoot = core.NewAggregateRoot(bowlingGame)
	return bowlingGame
}

func NewBowlingGame(gameID string) *BowlingGame {
	bowlingGame := newBowlingGame()
	bowlingGame.Apply(NewBowlingGameCreated(gameID))
	return bowlingGame
}

func ReplayBowlingGame(domainEvents []core.DomainEvent) *BowlingGame {
	bowlingGame := newBowlingGame()
	for _, event := range domainEvents {
		bowlingGame.Apply(event)
		bowlingGame.ClearDomainEvents()
	}
	return bowlingGame
}

func (b BowlingGame) GameID() string {
	return b.gameID
}

func (b BowlingGame) Score() int {
	return b.score
}

func (b BowlingGame) ThrowingCount() int {
	return b.throwingCount
}

func (b *BowlingGame) When(domainEvent core.DomainEvent) {
	switch event := interface{}(domainEvent).(type) {
	case BowlingGameCreated:
		b.gameID = event.gameID
	}
}
