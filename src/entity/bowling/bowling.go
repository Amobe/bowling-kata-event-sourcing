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

type XBowling interface {
	Throw(hit uint32) []event.Event
}

type Bowling struct {
	ID          string
	FrameNumber uint32
	Status      valueobject.BowlingStatus
	Score       uint32
	Games       map[uint32]valueobject.BowlingGame

	version int
}

func NewBowling(id string) *Bowling {
	b := &Bowling{
		ID:          id,
		FrameNumber: 1,
		Status:      valueobject.GamePrepared,
		Games:       make(map[uint32]valueobject.BowlingGame, standardFrameNumber+maxExtraFrameNumber),
	}
	return b
}

func (b *Bowling) Throw(hit uint32) (evs []event.Event) {
	currentGame, ok := b.Games[b.FrameNumber]
	if !ok {
		currentGame = createNewGame(b.FrameNumber)
	}
	evs = append(evs, b.raise(calculateGameHit(hit, currentGame)))

	if b.FrameNumber > 1 {
		evs = append(evs,
			b.raise(calculateGameBonus(hit, b.Games[b.FrameNumber-1])))
	}
	if b.FrameNumber > 2 {
		evs = append(evs,
			b.raise(calculateGameBonus(hit, b.Games[b.FrameNumber-2])))
	}

	if b.Status == valueobject.FrameFinished {
		return
	}
	evs = append(evs,
		b.raise(event.NewThrownEvent(valueobject.Thrown, calculateScore(b.Games))))

	if NoMoreHit(b.Games[b.FrameNumber]) || b.FrameNumber > 10 {
		evs = append(evs, b.raise(b.Reload()))
	}
	return
}

func (b *Bowling) Reload() event.Event {
	status := valueobject.FrameFinished
	frameNumber := b.FrameNumber
	if b.hasExtraFrame(b.FrameNumber, b.Games[b.FrameNumber]) {
		status = b.Status
		frameNumber = b.FrameNumber + 1
	}
	return event.NewReloadedEvent(status, frameNumber)
}

func calculateGameHit(hit uint32, game valueobject.BowlingGame) event.Event {
	hitGame := gameHit(game, hit)
	return event.NewGameReplacedEvent(hitGame.FrameNumber, hitGame)
}

func calculateGameBonus(hit uint32, game valueobject.BowlingGame) event.Event {
	bonusedGame := bonus(game, hit)
	return event.NewGameBonusedEvent(
		bonusedGame.FrameNumber, bonusedGame.Score, bonusedGame.ExtraBonus)
}

func createNewGame(frameNumber uint32) valueobject.BowlingGame {
	if frameNumber > FrameWithExtraBonus {
		return newGameWithoutExtraBonus(frameNumber)
	}
	return newGame(frameNumber)
}

func calculateScore(games map[uint32]valueobject.BowlingGame) (score uint32) {
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

func (b *Bowling) raise(ev event.Event) event.Event {
	on(ev, b)
	b.version++
	return ev
}
