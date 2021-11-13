package event

import (
	"github.com/amobe/bowling-kata-event-sourcing/src/valueobject"
)

type GameReplacedEvent struct {
	baseEvent
	ID   uint32 `json:"id"`
	Game valueobject.BowlingGame
}

type GameBonusedEvent struct {
	baseEvent
	ID         uint32
	Score      uint32
	ExtraBonus uint32
}

type ThrownEvent struct {
	baseEvent
	Status valueobject.BowlingStatus
	Score  uint32
}

type ReloadedEvent struct {
	baseEvent
	Status      valueobject.BowlingStatus
	FrameNumber uint32
}
