package bowling

import (
	"github.com/amobe/bowling-kata-event-sourcing/src/event"
	"github.com/amobe/bowling-kata-event-sourcing/src/valueobject"
)

const (
	standardFrameNumber = 10
	maxExtraFrameNumber = 2
	FrameWithExtraBonus = 9
)

type Bowling struct {
	ID          string
	FrameNumber uint32
	Status      valueobject.BowlingStatus
	Score       uint32
	Games       map[uint32]Game
}

func NewBowling(id string) *Bowling {
	b := &Bowling{
		ID:          id,
		FrameNumber: 1,
		Status:      valueobject.GamePrepared,
		Games:       make(map[uint32]Game, standardFrameNumber+maxExtraFrameNumber),
	}
	return b
}

func (b *Bowling) Throw(hit uint32) event.ThrownEvent {
	currentGame, ok := b.Games[b.FrameNumber]
	if !ok {
		currentGame = b.createNewGame(b.FrameNumber)
	}
	b.ApplyGameHitEvent(b.calculateGameHit(hit, currentGame))

	if b.FrameNumber > 1 {
		b.ApplyGameBonusedEvent(b.calculateGameBonus(hit, b.Games[b.FrameNumber-1]))
	}
	if b.FrameNumber > 2 {
		b.ApplyGameBonusedEvent(b.calculateGameBonus(hit, b.Games[b.FrameNumber-2]))
	}

	ev := event.ThrownEvent{
		Status: valueobject.Thrown,
		Score:  b.calculateScore(b.Games),
	}
	b.ApplyThrownEvent(ev)

	if b.Games[b.FrameNumber].NoMoreHit() || b.FrameNumber > 10 {
		b.ApplyReloadedEvent(b.Reload())
	}

	return ev
}

func (b *Bowling) calculateGameHit(hit uint32, game Game) event.GameHitEvent {
	hitGame := game.Hit(hit)
	return event.GameHitEvent{
		ID:          hitGame.ID,
		ThrowNumber: hitGame.ThrowNumber,
		Score:       hitGame.Score,
		Left:        hitGame.Left,
		Status:      hitGame.Status,
		ExtraBonus:  hitGame.ExtraBonus,
	}
}

func (b *Bowling) ApplyGameHitEvent(ev event.GameHitEvent) {
	g := b.createNewGame(ev.ID)
	g.ID = ev.ID
	g.ThrowNumber = ev.ThrowNumber
	g.Score = ev.Score
	g.Left = ev.Left
	g.Status = ev.Status
	g.ExtraBonus = ev.ExtraBonus
	b.Games[ev.ID] = g
}

func (b *Bowling) calculateGameBonus(hit uint32, game Game) event.GameBonusedEvent {
	bonusedGame := game.Bonus(hit)
	return event.GameBonusedEvent{
		ID:         bonusedGame.ID,
		Score:      bonusedGame.Score,
		ExtraBonus: bonusedGame.ExtraBonus,
	}
}

func (b *Bowling) ApplyGameBonusedEvent(ev event.GameBonusedEvent) {
	g := b.Games[ev.ID]
	g.Score = ev.Score
	g.ExtraBonus = ev.ExtraBonus
	b.Games[ev.ID] = g
}

func (b *Bowling) ApplyThrownEvent(ev event.ThrownEvent) {
	switch b.Status {
	case valueobject.FrameFinished:
		return
	default:
		b.Status = ev.Status
		b.Score = ev.Score
	}
}

func (b *Bowling) Reload() event.ReloadedEvent {
	if b.hasNoExtraFrame(b.FrameNumber, b.Games[b.FrameNumber]) {
		return event.ReloadedEvent{
			Status:      valueobject.FrameFinished,
			FrameNumber: b.FrameNumber,
		}
	}
	return event.ReloadedEvent{
		Status:      b.Status,
		FrameNumber: b.FrameNumber + 1,
	}
}

func (b *Bowling) ApplyReloadedEvent(ev event.ReloadedEvent) {
	b.Status = ev.Status
	b.FrameNumber = ev.FrameNumber
}

func (b *Bowling) createNewGame(frameNumber uint32) Game {
	if frameNumber > FrameWithExtraBonus {
		return NewGameWithoutExtraBonus(frameNumber)
	}
	return NewGame(frameNumber)
}

func (b *Bowling) calculateScore(games map[uint32]Game) (score uint32) {
	for _, g := range games {
		score = score + g.Score
	}
	return
}

func (b *Bowling) hasNoExtraFrame(frameNumber uint32, game Game) bool {
	openEnd := b.FrameNumber >= standardFrameNumber && b.Games[b.FrameNumber].Status == valueobject.Open
	strikeTwice := b.FrameNumber == standardFrameNumber+maxExtraFrameNumber && b.Games[b.FrameNumber].Status == valueobject.Strike
	return openEnd || strikeTwice
}
