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
	Games       map[uint32]valueobject.BowlingGame

	repo    event.Repository
	version int
}

func NewBowling(id string, storage event.Repository) *Bowling {
	b := &Bowling{
		ID:          id,
		FrameNumber: 1,
		Status:      valueobject.GamePrepared,
		Games:       make(map[uint32]valueobject.BowlingGame, standardFrameNumber+maxExtraFrameNumber),
		repo:        storage,
	}
	return b
}

func (b *Bowling) Throw(hit uint32) {
	currentGame, ok := b.Games[b.FrameNumber]
	if !ok {
		currentGame = b.createNewGame(b.FrameNumber)
	}
	b.calculateGameHit(hit, currentGame)

	if b.FrameNumber > 1 {
		b.calculateGameBonus(hit, b.Games[b.FrameNumber-1])
	}
	if b.FrameNumber > 2 {
		b.calculateGameBonus(hit, b.Games[b.FrameNumber-2])
	}

	b.raise(&event.ThrownEvent{
		Status: valueobject.Thrown,
		Score:  b.calculateScore(b.Games),
	})

	if b.NoMoreHit(b.Games[b.FrameNumber]) || b.FrameNumber > 10 {
		b.Reload()
	}
}

func (b *Bowling) Reload() {
	var ev *event.ReloadedEvent
	if b.hasNoExtraFrame(b.FrameNumber, b.Games[b.FrameNumber]) {
		ev = &event.ReloadedEvent{
			Status:      valueobject.FrameFinished,
			FrameNumber: b.FrameNumber,
		}
	} else {
		ev = &event.ReloadedEvent{
			Status:      b.Status,
			FrameNumber: b.FrameNumber + 1,
		}
	}
	b.raise(ev)
}

func (b *Bowling) calculateGameHit(hit uint32, game valueobject.BowlingGame) {
	hitGame := b.Hit(game, hit)
	b.raise(&event.GameHitEvent{
		ID:          hitGame.FrameNumber,
		ThrowNumber: hitGame.ThrowNumber,
		Score:       hitGame.Score,
		Left:        hitGame.Left,
		Status:      hitGame.Status,
		ExtraBonus:  hitGame.ExtraBonus,
	})
}

func (b *Bowling) calculateGameBonus(hit uint32, game valueobject.BowlingGame) {
	bonusedGame := b.Bonus(game, hit)
	b.raise(&event.GameBonusedEvent{
		ID:         bonusedGame.FrameNumber,
		Score:      bonusedGame.Score,
		ExtraBonus: bonusedGame.ExtraBonus,
	})
}

func (b *Bowling) createNewGame(frameNumber uint32) valueobject.BowlingGame {
	if frameNumber > FrameWithExtraBonus {
		return b.NewBowlingGameWithoutExtraBonus(frameNumber)
	}
	return b.NewBowlingGame(frameNumber)
}

func (b *Bowling) calculateScore(games map[uint32]valueobject.BowlingGame) (score uint32) {
	for _, g := range games {
		score = score + g.Score
	}
	return
}

func (b *Bowling) hasNoExtraFrame(frameNumber uint32, game valueobject.BowlingGame) bool {
	openEnd := b.FrameNumber >= standardFrameNumber && b.Games[b.FrameNumber].Status == valueobject.Open
	strikeTwice := b.FrameNumber == standardFrameNumber+maxExtraFrameNumber && b.Games[b.FrameNumber].Status == valueobject.Strike
	return openEnd || strikeTwice
}

func (b *Bowling) raise(ev event.Event) {
	_ = b.repo.Append(b.ID, ev)
	b.On(ev, true)
}

func (b *Bowling) On(changed event.Event, isNew bool) {
	switch ev := changed.(type) {
	case *event.ThrownEvent:
		b.ApplyThrownEvent(ev)
	case *event.GameHitEvent:
		b.ApplyGameHitEvent(ev)
	case *event.GameBonusedEvent:
		b.ApplyGameBonusedEvent(ev)
	case *event.ReloadedEvent:
		b.ApplyReloadedEvent(ev)
	}
	if !isNew {
		b.version++
	}
}

func (b *Bowling) ApplyThrownEvent(ev *event.ThrownEvent) {
	switch b.Status {
	case valueobject.FrameFinished:
		return
	default:
		b.Status = ev.Status
		b.Score = ev.Score
	}
}

func (b *Bowling) ApplyGameHitEvent(ev *event.GameHitEvent) {
	g := b.createNewGame(ev.ID)
	g.FrameNumber = ev.ID
	g.ThrowNumber = ev.ThrowNumber
	g.Score = ev.Score
	g.Left = ev.Left
	g.Status = ev.Status
	g.ExtraBonus = ev.ExtraBonus
	b.Games[ev.ID] = g
}

func (b *Bowling) ApplyGameBonusedEvent(ev *event.GameBonusedEvent) {
	g := b.Games[ev.ID]
	g.Score = ev.Score
	g.ExtraBonus = ev.ExtraBonus
	b.Games[ev.ID] = g
}

func (b *Bowling) ApplyReloadedEvent(ev *event.ReloadedEvent) {
	b.Status = ev.Status
	b.FrameNumber = ev.FrameNumber
}
