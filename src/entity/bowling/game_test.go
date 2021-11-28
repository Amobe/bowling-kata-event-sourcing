package bowling_test

import (
	"testing"

	"github.com/amobe/bowling-kata-event-sourcing/src/entity/bowling"
	"github.com/amobe/bowling-kata-event-sourcing/src/storage"
	"github.com/amobe/bowling-kata-event-sourcing/src/valueobject"
	"github.com/stretchr/testify/suite"
)

type GameSuite struct {
	suite.Suite
}

func TestGameSuite(t *testing.T) {
	suite.Run(t, new(GameSuite))
}

func (s *GameSuite) TestSpareHit() {
	b := bowling.NewBowling("0", storage.NewInmemEventStorage())
	g := b.NewBowlingGame(0)

	firstGame := b.Hit(g, 1)
	got := b.Hit(firstGame, 9)

	s.Equal(uint32(1), got.ExtraBonus)
	s.Equal(uint32(0), got.Left)
	s.Equal(uint32(10), got.Score)
	s.Equal(valueobject.Spare, got.Status)
	s.Equal(uint32(2), got.ThrowNumber)
}

func (s *GameSuite) TestStrikeHit() {
	b := bowling.NewBowling("0", storage.NewInmemEventStorage())
	g := b.NewBowlingGame(0)

	got := b.Hit(g, 10)

	s.Equal(uint32(2), got.ExtraBonus)
	s.Equal(uint32(0), got.Left)
	s.Equal(uint32(10), got.Score)
	s.Equal(valueobject.Strike, got.Status)
	s.Equal(uint32(1), got.ThrowNumber)
}

func (s *GameSuite) TestOpenGame() {
	b := bowling.NewBowling("0", storage.NewInmemEventStorage())
	g := b.NewBowlingGame(0)

	firstGame := b.Hit(g, 1)
	got := b.Hit(firstGame, 1)

	s.Equal(uint32(0), got.ExtraBonus)
	s.Equal(uint32(8), got.Left)
	s.Equal(uint32(2), got.Score)
	s.Equal(valueobject.Open, got.Status)
	s.Equal(uint32(2), got.ThrowNumber)
}

func (s *GameSuite) TestGameWithoutExtraBonus() {
	b := bowling.NewBowling("0", storage.NewInmemEventStorage())
	g := b.NewBowlingGameWithoutExtraBonus(0)

	got := b.Hit(g, 10)

	s.Equal(uint32(0), got.ExtraBonus)
	s.Equal(uint32(0), got.Left)
	s.Equal(uint32(10), got.Score)
	s.Equal(valueobject.Strike, got.Status)
	s.Equal(uint32(1), got.ThrowNumber)
}

func (s *GameSuite) TestNoMoreHit() {
	twoHitGame := valueobject.BowlingGame{ThrowNumber: 2}
	s.True(bowling.NoMoreHit(twoHitGame))
	strikeGame := valueobject.BowlingGame{Status: valueobject.Strike}
	s.True(bowling.NoMoreHit(strikeGame))
}

func (s *GameSuite) TestHitAfterNoMoreHit() {
	b := bowling.NewBowling("0", storage.NewInmemEventStorage())
	g := b.NewBowlingGame(0)

	firstGame := b.Hit(g, 10)
	got := b.Hit(firstGame, 1)

	s.Equal(uint32(2), got.ExtraBonus)
	s.Equal(uint32(0), got.Left)
	s.Equal(uint32(10), got.Score)
	s.Equal(valueobject.Strike, got.Status)
	s.Equal(uint32(1), got.ThrowNumber)
}
