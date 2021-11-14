package event

import (
	"github.com/amobe/bowling-kata-event-sourcing/src/valueobject"
)

type GameReplacedEvent struct {
	baseEvent
	FrameNumber uint32
	Game        valueobject.BowlingGame
}

func NewGameReplacedEvent(frameNumber uint32, game valueobject.BowlingGame) *GameReplacedEvent {
	return &GameReplacedEvent{
		baseEvent:   baseEvent{name: "GameReplacedEvent"},
		FrameNumber: frameNumber,
		Game:        game,
	}
}

type GameBonusedEvent struct {
	baseEvent
	FrameNumber uint32
	Score       uint32
	ExtraBonus  uint32
}

func NewGameBonusedEvent(frameNumber uint32, score uint32, extraBonus uint32) *GameBonusedEvent {
	return &GameBonusedEvent{
		baseEvent:   baseEvent{name: "GameBonusedEvent"},
		FrameNumber: frameNumber,
		Score:       score,
		ExtraBonus:  extraBonus,
	}
}

type ThrownEvent struct {
	baseEvent
	Status valueobject.BowlingStatus
	Score  uint32
}

func NewThrownEvent(status valueobject.BowlingStatus, score uint32) *ThrownEvent {
	return &ThrownEvent{
		baseEvent: baseEvent{name: "ThrownEvent"},
		Status:    status,
		Score:     score,
	}
}

type ReloadedEvent struct {
	baseEvent
	Status      valueobject.BowlingStatus
	FrameNumber uint32
}

func NewReloadedEvent(status valueobject.BowlingStatus, frameNumber uint32) *ReloadedEvent {
	return &ReloadedEvent{
		baseEvent:   baseEvent{name: "ReloadedEvent"},
		Status:      status,
		FrameNumber: frameNumber,
	}
}
