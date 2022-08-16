package entity

import (
	"testing"

	core "github.com/amobe/bowling-kata-event-sourcing/src/core/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateBowlingGameGenerateABowlingGameCreatedEvent(t *testing.T) {
	bowlingGame := NewBowlingGame("")
	assert.Equal(t, 1, len(bowlingGame.DomainEvents()))
	assert.IsType(t, BowlingGameCreated{}, bowlingGame.DomainEvents()[0])
}

func TestReplayBowlingGameFromBowlingGameCreatedEvent(t *testing.T) {
	event := NewBowlingGameCreated("abc")
	bowlingGame := ReplayBowlingGame([]core.DomainEvent{event})
	assert.Equal(t, event.gameID, bowlingGame.GameID())
	assert.Equal(t, 0, len(bowlingGame.DomainEvents()))
}

func TestThrowCommandGenerateBowlingGameThrownEvent(t *testing.T) {
	event := NewBowlingGameCreated("abc")
	bowlingGame := ReplayBowlingGame([]core.DomainEvent{event})
	bowlingGame.RollABall(1)
	assert.Equal(t, 1, len(bowlingGame.DomainEvents()))
	assert.IsType(t, BowlingGameRolledABall{}, bowlingGame.DomainEvents()[0])
}
