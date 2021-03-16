package bowling_test

import (
	"testing"

	"github.com/amobe/bowling-kata-event-sourcing/src/entity/bowling"
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
	g := bowling.NewGame(0)

	got := g.Hit(1).Hit(9)

	s.Equal(uint32(1), got.ExtraBonus)
	s.Equal(uint32(0), got.Left)
	s.Equal(uint32(10), got.Score)
	s.Equal(valueobject.Spare, got.Status)
	s.Equal(uint32(2), got.ThrowNumber)
}

func (s *GameSuite) TestStrikeHit() {
	g := bowling.NewGame(0)

	got := g.Hit(10)

	s.Equal(uint32(2), got.ExtraBonus)
	s.Equal(uint32(0), got.Left)
	s.Equal(uint32(10), got.Score)
	s.Equal(valueobject.Strike, got.Status)
	s.Equal(uint32(1), got.ThrowNumber)
}

func (s *GameSuite) TestOpenGame() {
	g := bowling.NewGame(0)

	got := g.Hit(1).Hit(1)

	s.Equal(uint32(0), got.ExtraBonus)
	s.Equal(uint32(8), got.Left)
	s.Equal(uint32(2), got.Score)
	s.Equal(valueobject.Open, got.Status)
	s.Equal(uint32(2), got.ThrowNumber)
}

func (s *GameSuite) TestSpareAndBonus() {
	g := bowling.NewGame(0)

	got := g.Hit(1).Hit(9).Bonus(1)

	s.Equal(uint32(0), got.ExtraBonus)
	s.Equal(uint32(0), got.Left)
	s.Equal(uint32(11), got.Score)
	s.Equal(valueobject.Spare, got.Status)
	s.Equal(uint32(2), got.ThrowNumber)
}

func (s *GameSuite) TestStrikeAndBonus() {
	g := bowling.NewGame(0)

	got := g.Hit(10).Bonus(1).Bonus(1)

	s.Equal(uint32(0), got.ExtraBonus)
	s.Equal(uint32(0), got.Left)
	s.Equal(uint32(12), got.Score)
	s.Equal(valueobject.Strike, got.Status)
	s.Equal(uint32(1), got.ThrowNumber)
}

func (s *GameSuite) TestGameWithoutExtraBonus() {
	g := bowling.NewGameWithoutExtraBonus(0)

	got := g.Hit(10)

	s.Equal(uint32(0), got.ExtraBonus)
	s.Equal(uint32(0), got.Left)
	s.Equal(uint32(10), got.Score)
	s.Equal(valueobject.Strike, got.Status)
	s.Equal(uint32(1), got.ThrowNumber)
}

func (s *GameSuite) TestNoMoreHit() {
	twoHitGame := bowling.Game{ThrowNumber: 2}
	s.True(twoHitGame.NoMoreHit())
	strikeGame := bowling.Game{Status: valueobject.Strike}
	s.True(strikeGame.NoMoreHit())
}

func (s *GameSuite) TestHitAfterNoMoreHit() {
	g := bowling.NewGame(0)

	got := g.Hit(10).Hit(1)

	s.Equal(uint32(2), got.ExtraBonus)
	s.Equal(uint32(0), got.Left)
	s.Equal(uint32(10), got.Score)
	s.Equal(valueobject.Strike, got.Status)
	s.Equal(uint32(1), got.ThrowNumber)
}
