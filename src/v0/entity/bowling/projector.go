package bowling

import (
	event2 "github.com/amobe/bowling-kata-event-sourcing/src/v0/event"
)

func Projector(id string, evs []event2.Event) *Bowling {
	b := NewBowling(id)
	for _, ev := range evs {
		on(ev, b)
	}
	return b
}

func on(ev event2.Event, b *Bowling) {
	switch data := ev.Data().(type) {
	case *event2.ThrownEvent:
		applyThrownEvent(data, b)
	case *event2.GameReplacedEvent:
		applyGameReplacedEvent(data, b)
	case *event2.GameBonusedEvent:
		applyGameBonusedEvent(data, b)
	case *event2.ReloadedEvent:
		applyReloadedEvent(data, b)
	}
}

func applyThrownEvent(ev *event2.ThrownEvent, b *Bowling) {
	b.Status = ev.Status
	b.Score = ev.Score
}

func applyGameReplacedEvent(ev *event2.GameReplacedEvent, b *Bowling) {
	b.Games[ev.FrameNumber] = ev.Game
}

func applyGameBonusedEvent(ev *event2.GameBonusedEvent, b *Bowling) {
	g := b.Games[ev.FrameNumber]
	g.Score = ev.Score
	g.ExtraBonus = ev.ExtraBonus
	b.Games[ev.FrameNumber] = g
}

func applyReloadedEvent(ev *event2.ReloadedEvent, b *Bowling) {
	b.Status = ev.Status
	b.FrameNumber = ev.FrameNumber
}
