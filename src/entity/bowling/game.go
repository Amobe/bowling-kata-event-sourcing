package bowling

import "github.com/amobe/bowling-kata-event-sourcing/src/valueobject"

const StandardPins = standardPins
const standardPins = 10

func (b *Bowling) NewBowlingGame(frameNubmer uint32) valueobject.BowlingGame {
	return valueobject.BowlingGame{
		FrameNumber: frameNubmer,
		Left:        standardPins,
	}
}

func (b *Bowling) NewBowlingGameWithoutExtraBonus(frameNubmer uint32) valueobject.BowlingGame {
	return valueobject.BowlingGame{
		FrameNumber:       frameNubmer,
		Left:              standardPins,
		WithoutExtraBonus: true,
	}
}

func (b *Bowling) Hit(game valueobject.BowlingGame, pins uint32) valueobject.BowlingGame {
	if b.NoMoreHit(game) {
		return game
	}
	game.ThrowNumber = game.ThrowNumber + 1
	game.Left = game.Left - pins
	game.Score = game.Score + pins
	game.Status, game.ExtraBonus = b.getStatusBonus(game)
	if game.WithoutExtraBonus {
		game.ExtraBonus = 0
	}
	return game
}

func (b *Bowling) Bonus(game valueobject.BowlingGame, pins uint32) valueobject.BowlingGame {
	if game.ExtraBonus > 0 {
		game.Score = game.Score + pins
		game.ExtraBonus = game.ExtraBonus - 1
	}
	return game
}

func (b *Bowling) isStrike(game valueobject.BowlingGame) bool {
	return game.ThrowNumber == 1 && game.Left == 0
}

func (b *Bowling) isSpare(game valueobject.BowlingGame) bool {
	return game.ThrowNumber == 2 && game.Left == 0
}

func (b *Bowling) gainExtraBonus(game valueobject.BowlingGame, extraBonus uint32) uint32 {
	if game.WithoutExtraBonus {
		return 0
	}
	return extraBonus
}

func (b *Bowling) NoMoreHit(game valueobject.BowlingGame) bool {
	return NoMoreHit(game.ThrowNumber, game.Status)
}

func NoMoreHit(throwNumber uint32, status valueobject.BowlingGameStatus) bool {
	return throwNumber == 2 || status == valueobject.Strike
}

func (b *Bowling) getStatusBonus(game valueobject.BowlingGame) (s valueobject.BowlingGameStatus, extraBonus uint32) {
	if b.isSpare(game) {
		return valueobject.Spare, b.gainExtraBonus(game, 1)
	} else if b.isStrike(game) {
		return valueobject.Strike, b.gainExtraBonus(game, 2)
	}
	return valueobject.Open, 0
}
