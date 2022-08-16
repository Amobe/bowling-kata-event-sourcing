package event

import (
	"github.com/amobe/bowling-kata-event-sourcing/src/v0/valueobject"
)

type GameReplacedEvent struct {
	FrameNumber uint32
	Game        valueobject.BowlingGame
}

func NewGameReplacedEvent(frameNumber uint32, game valueobject.BowlingGame) Event {
	return event{
		name: "GameReplacedEvent",
		data: &GameReplacedEvent{
			FrameNumber: frameNumber,
			Game:        game,
		},
	}
}

type GameBonusedEvent struct {
	FrameNumber uint32
	Score       uint32
	ExtraBonus  uint32
}

func NewGameBonusedEvent(frameNumber uint32, score uint32, extraBonus uint32) Event {
	return event{
		name: "GameBonusedEvent",
		data: &GameBonusedEvent{
			FrameNumber: frameNumber,
			Score:       score,
			ExtraBonus:  extraBonus,
		},
	}
}

type ThrownEvent struct {
	Status valueobject.BowlingStatus
	Score  uint32
}

func NewThrownEvent(status valueobject.BowlingStatus, score uint32) Event {
	return event{
		name: "ThrownEvent",
		data: &ThrownEvent{
			Status: status,
			Score:  score,
		},
	}
}

type ReloadedEvent struct {
	Status      valueobject.BowlingStatus
	FrameNumber uint32
}

func NewReloadedEvent(status valueobject.BowlingStatus, frameNumber uint32) Event {
	return event{
		name: "ReloadedEvent",
		data: &ReloadedEvent{
			Status:      status,
			FrameNumber: frameNumber,
		},
	}
}
