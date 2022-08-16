package entity

import (
	"fmt"

	core "github.com/amobe/bowling-kata-event-sourcing/src/core/entity"
)

type BowlingGame struct {
	*core.AggregateRoot
	gameID             string
	score              int
	leavingPins        int
	bonusRemainHit     int
	bonusChance        []int
	finishedFrameCount int
	bonusFrame         int
}

func newBowlingGame() *BowlingGame {
	bowlingGame := &BowlingGame{}
	bowlingGame.AggregateRoot = core.NewAggregateRoot(bowlingGame)
	return bowlingGame
}

func NewBowlingGame(gameID string) *BowlingGame {
	bowlingGame := newBowlingGame()
	bowlingGame.Apply(NewBowlingGameCreated(gameID))
	return bowlingGame
}

func ReplayBowlingGame(domainEvents []core.DomainEvent) *BowlingGame {
	bowlingGame := newBowlingGame()
	for _, event := range domainEvents {
		bowlingGame.Apply(event)
		bowlingGame.ClearDomainEvents()
	}
	return bowlingGame
}

func (b BowlingGame) GameID() string {
	return b.gameID
}

func (b BowlingGame) Score() int {
	return b.score
}

func (b BowlingGame) LeavingPins() int {
	return b.leavingPins
}

func (b BowlingGame) BonusRemainHit() int {
	return b.bonusRemainHit
}

func (b *BowlingGame) RollABall(hit int) error {
	if (b.finishedFrameCount-b.bonusFrame > 9) ||
		(b.finishedFrameCount > 9 && (b.bonusFrame == 1 && b.bonusRemainHit == 1)) {
		return fmt.Errorf("unavailable to roll a ball after game finished")
	}
	if hit < 0 {
		return fmt.Errorf("hit number should not less than 0")
	}
	b.Apply(NewBowlingGameRollABall(hit))
	return nil
}

func (b *BowlingGame) When(domainEvent core.DomainEvent) {
	switch event := interface{}(domainEvent).(type) {
	case BowlingGameCreated:
		b.gameID = event.gameID
		b.leavingPins = 10
		b.bonusRemainHit = 2
	case BowlingGameRolledABall:
		b.score += event.hit
		b.leavingPins -= event.hit
		b.bonusRemainHit--
		for i, chance := range b.bonusChance {
			if chance > 0 {
				b.score += event.hit
				b.bonusChance[i] = chance - 1
			}
		}
		if frameHasBonusChance(b.finishedFrameCount) {
			if strike(b.leavingPins, b.bonusRemainHit) {
				b.bonusChance = append(b.bonusChance, 2)
			}
			if spare(b.leavingPins, b.bonusRemainHit) {
				b.bonusChance = append(b.bonusChance, 1)
			}
		}
		if spare(b.leavingPins, b.bonusRemainHit) && b.finishedFrameCount == 9 {
			b.bonusFrame++
		}
		if strike(b.leavingPins, b.bonusRemainHit) && b.finishedFrameCount >= 9 {
			b.bonusFrame++
		}
		if b.leavingPins == 0 || b.bonusRemainHit == 0 {
			b.finishedFrameCount++
			b.leavingPins = 10
			b.bonusRemainHit = 2
		}
	}
}

func frameHasBonusChance(finishedFrameCount int) bool {
	return finishedFrameCount < 9
}

func strike(leavingPins, bonusRemainHit int) bool {
	return leavingPins == 0 && bonusRemainHit == 1
}

func spare(leavingPins, bonusRemainHit int) bool {
	return leavingPins == 0 && bonusRemainHit == 0
}
