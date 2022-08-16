package bowling

import (
	event2 "github.com/amobe/bowling-kata-event-sourcing/src/v0/event"
	"github.com/amobe/bowling-kata-event-sourcing/src/v0/valueobject"
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

func (b *Bowling) Throw(hit uint32) (evs []event2.Event) {
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
		b.raise(event2.NewThrownEvent(valueobject.Thrown, calculateScore(b.Games))))

	if NoMoreHit(b.Games[b.FrameNumber]) || b.FrameNumber > 10 {
		evs = append(evs, b.raise(b.Reload()))
	}
	return
}

func (b *Bowling) Reload() event2.Event {
	status := valueobject.FrameFinished
	frameNumber := b.FrameNumber
	if hasExtraFrame(b.FrameNumber, b.Games[b.FrameNumber]) {
		status = b.Status
		frameNumber = b.FrameNumber + 1
	}
	return event2.NewReloadedEvent(status, frameNumber)
}

func hasExtraFrame(frameNumber uint32, game valueobject.BowlingGame) bool {
	openEnd := frameNumber >= standardFrameNumber && game.Status == valueobject.Open
	strikeTwice := frameNumber == standardFrameNumber+maxExtraFrameNumber && game.Status == valueobject.Strike
	return !(openEnd || strikeTwice)
}

func (b *Bowling) raise(ev event2.Event) event2.Event {
	on(ev, b)
	b.version++
	return ev
}

func calculateGameHit(hit uint32, game valueobject.BowlingGame) event2.Event {
	hitGame := gameHit(game, hit)
	return event2.NewGameReplacedEvent(hitGame.FrameNumber, hitGame)
}

func calculateGameBonus(hit uint32, game valueobject.BowlingGame) event2.Event {
	bonusedGame := bonus(game, hit)
	return event2.NewGameBonusedEvent(
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
