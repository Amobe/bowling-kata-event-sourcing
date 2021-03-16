package event

import (
	"github.com/amobe/bowling-kata-event-sourcing/src/valueobject"
)

type GameHitEvent struct {
	ID          uint32
	ThrowNumber uint32
	Score       uint32
	Left        uint32
	Status      valueobject.BowlingGameStatus
	ExtraBonus  uint32
}

type GameBonusedEvent struct {
	ID         uint32
	Score      uint32
	ExtraBonus uint32
}

type ThrownEvent struct {
	Status valueobject.BowlingStatus
	Score  uint32
	// ExtraFrame uint32
}

type ReloadedEvent struct {
	Status      valueobject.BowlingStatus
	FrameNumber uint32
}
