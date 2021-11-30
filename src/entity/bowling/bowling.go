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

	if b.Status == valueobject.FrameFinished {
		return
	}
	b.raise(event.NewThrownEvent(valueobject.Thrown, b.calculateScore(b.Games)))

	if NoMoreHit(b.Games[b.FrameNumber]) || b.FrameNumber > 10 {
		b.Reload()
	}
}

func (b *Bowling) Reload() {
	status := valueobject.FrameFinished
	frameNumber := b.FrameNumber
	if b.hasExtraFrame(b.FrameNumber, b.Games[b.FrameNumber]) {
		status = b.Status
		frameNumber = b.FrameNumber + 1
	}
	b.raise(event.NewReloadedEvent(status, frameNumber))
}

func (b *Bowling) calculateGameHit(hit uint32, game valueobject.BowlingGame) {
	hitGame := gameHit(game, hit)
	b.raise(event.NewGameReplacedEvent(hitGame.FrameNumber, hitGame))
}

func (b *Bowling) calculateGameBonus(hit uint32, game valueobject.BowlingGame) {
	bonusedGame := bonus(game, hit)
	b.raise(event.NewGameBonusedEvent(
		bonusedGame.FrameNumber, bonusedGame.Score, bonusedGame.ExtraBonus))
}

func (b *Bowling) createNewGame(frameNumber uint32) valueobject.BowlingGame {
	if frameNumber > FrameWithExtraBonus {
		return newGameWithoutExtraBonus(frameNumber)
	}
	return newGame(frameNumber)
}

func (b *Bowling) calculateScore(games map[uint32]valueobject.BowlingGame) (score uint32) {
	for _, g := range games {
		score = score + g.Score
	}
	return
}

func (b *Bowling) hasExtraFrame(frameNumber uint32, game valueobject.BowlingGame) bool {
	openEnd := b.FrameNumber >= standardFrameNumber && b.Games[b.FrameNumber].Status == valueobject.Open
	strikeTwice := b.FrameNumber == standardFrameNumber+maxExtraFrameNumber && b.Games[b.FrameNumber].Status == valueobject.Strike
	return !(openEnd || strikeTwice)
}

func (b *Bowling) raise(ev event.Event) {
	_ = b.repo.Append(b.ID, ev)
	on(ev, b)
	b.version++
}
