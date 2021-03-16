package bowling

import "github.com/amobe/bowling-kata-event-sourcing/src/valueobject"

const standardPins = 10

type Game struct {
	ID                uint32
	ThrowNumber       uint32
	Score             uint32
	Left              uint32
	Status            valueobject.BowlingGameStatus
	WithoutExtraBonus bool
	ExtraBonus        uint32
}

func NewGame(id uint32) Game {
	return Game{
		ID:   id,
		Left: standardPins,
	}
}

func NewGameWithoutExtraBonus(id uint32) Game {
	g := NewGame(id)
	g.WithoutExtraBonus = true
	return g
}

func (g Game) Hit(pins uint32) Game {
	if g.NoMoreHit() {
		return g
	}
	g.ThrowNumber = g.ThrowNumber + 1
	g.Left = g.Left - pins
	g.Score = g.Score + pins
	g.Status, g.ExtraBonus = g.getStatusBonus()
	if g.WithoutExtraBonus {
		g.ExtraBonus = 0
	}
	return g
}

func (g Game) Bonus(pins uint32) Game {
	if g.ExtraBonus > 0 {
		g.Score = g.Score + pins
		g.ExtraBonus = g.ExtraBonus - 1
	}
	return g
}

func (g Game) NoMoreHit() bool {
	return g.ThrowNumber == 2 || g.Status == valueobject.Strike
}

func (g Game) getStatusBonus() (s valueobject.BowlingGameStatus, extraBonus uint32) {
	if g.isSpare() {
		return valueobject.Spare, g.gainExtraBonus(1)
	} else if g.isStrike() {
		return valueobject.Strike, g.gainExtraBonus(2)
	}
	return valueobject.Open, 0
}

func (g Game) gainExtraBonus(extraBonus uint32) uint32 {
	if g.WithoutExtraBonus {
		return 0
	}
	return extraBonus
}

func (g Game) isStrike() bool {
	return g.ThrowNumber == 1 && g.Left == 0
}

func (g Game) isSpare() bool {
	return g.ThrowNumber == 2 && g.Left == 0
}
