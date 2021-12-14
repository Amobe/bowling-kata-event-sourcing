package bowling

import "github.com/amobe/bowling-kata-event-sourcing/src/event"

func Projector(id string, evs []event.Event) *Bowling {
	b := NewBowling(id)
	for _, ev := range evs {
		on(ev, b)
	}
	return b
}

func on(ev event.Event, b *Bowling) {
	switch data := ev.Data().(type) {
	case *event.ThrownEvent:
		applyThrownEvent(data, b)
	case *event.GameReplacedEvent:
		applyGameReplacedEvent(data, b)
	case *event.GameBonusedEvent:
		applyGameBonusedEvent(data, b)
	case *event.ReloadedEvent:
		applyReloadedEvent(data, b)
	}
}

func applyThrownEvent(ev *event.ThrownEvent, b *Bowling) {
	b.Status = ev.Status
	b.Score = ev.Score
}

func applyGameReplacedEvent(ev *event.GameReplacedEvent, b *Bowling) {
	b.Games[ev.FrameNumber] = ev.Game
}

func applyGameBonusedEvent(ev *event.GameBonusedEvent, b *Bowling) {
	g := b.Games[ev.FrameNumber]
	g.Score = ev.Score
	g.ExtraBonus = ev.ExtraBonus
	b.Games[ev.FrameNumber] = g
}

func applyReloadedEvent(ev *event.ReloadedEvent, b *Bowling) {
	b.Status = ev.Status
	b.FrameNumber = ev.FrameNumber
}
