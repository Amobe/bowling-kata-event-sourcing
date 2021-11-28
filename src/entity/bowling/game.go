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
	if NoMoreHit(game) {
		return game
	}
	game.ThrowNumber = game.ThrowNumber + 1
	game.Left = game.Left - pins
	game.Score = game.Score + pins
	game.Status, game.ExtraBonus = getStatusBonus(game)
	if game.WithoutExtraBonus {
		game.ExtraBonus = 0
	}
	return game
}

func bonus(game valueobject.BowlingGame, pins uint32) valueobject.BowlingGame {
	if game.ExtraBonus > 0 {
		game.Score = game.Score + pins
		game.ExtraBonus = game.ExtraBonus - 1
	}
	return game
}

func isStrike(game valueobject.BowlingGame) bool {
	return game.ThrowNumber == 1 && game.Left == 0
}

func isSpare(game valueobject.BowlingGame) bool {
	return game.ThrowNumber == 2 && game.Left == 0
}

func gainExtraBonus(game valueobject.BowlingGame, extraBonus uint32) uint32 {
	if game.WithoutExtraBonus {
		return 0
	}
	return extraBonus
}

func getStatusBonus(game valueobject.BowlingGame) (s valueobject.BowlingGameStatus, extraBonus uint32) {
	if isSpare(game) {
		return valueobject.Spare, gainExtraBonus(game, 1)
	} else if isStrike(game) {
		return valueobject.Strike, gainExtraBonus(game, 2)
	}
	return valueobject.Open, 0
}

func NoMoreHit(game valueobject.BowlingGame) bool {
	return game.ThrowNumber == 2 || game.Status == valueobject.Strike
}
