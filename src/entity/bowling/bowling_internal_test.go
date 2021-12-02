package bowling

import (
	"testing"

	"github.com/amobe/bowling-kata-event-sourcing/src/event"
	"github.com/amobe/bowling-kata-event-sourcing/src/valueobject"
	"github.com/stretchr/testify/assert"
)

func Test_calculateGameHit(t *testing.T) {
	hit := uint32(3)
	game := valueobject.BowlingGame{
		FrameNumber: 2,
		Left:        10,
	}
	want := event.NewGameReplacedEvent(2, valueobject.BowlingGame{
		FrameNumber: 2,
		Left:        7,
		Score:       3,
		ThrowNumber: 1,
		Status:      valueobject.Open,
	})
	got := calculateGameHit(hit, game)
	assert.EqualValues(t, want, got)
}

func Test_calculateGameBonus(t *testing.T) {
	hit := uint32(3)
	game := valueobject.BowlingGame{
		FrameNumber: 2,
		Left:        10,
	}
	want := event.NewGameBonusedEvent(2, 0, 0)
	got := calculateGameBonus(hit, game)
	assert.EqualValues(t, want, got)
}
