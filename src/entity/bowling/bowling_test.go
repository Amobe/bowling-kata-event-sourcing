package bowling_test

import (
	"testing"

	"github.com/amobe/bowling-kata-event-sourcing/src/entity/bowling"
	"github.com/amobe/bowling-kata-event-sourcing/src/storage"
	"github.com/amobe/bowling-kata-event-sourcing/src/valueobject"
	"github.com/stretchr/testify/suite"
)

type BowlingSuite struct {
	suite.Suite
}

func TestBowlingSuite(t *testing.T) {
	suite.Run(t, new(BowlingSuite))
}

func (s *BowlingSuite) TestThrowOneBall() {
	b := getPreparedBowlingGame()

	b.Throw(1)

	s.Equal(uint32(1), b.Score)
	s.Equal(valueobject.Thrown, b.Status)
	s.Equal(uint32(1), b.FrameNumber)
}

func (s *BowlingSuite) TestThrowOpenGame() {
	b := getPreparedBowlingGame()

	b.Throw(1)
	b.Throw(1)

	s.Equal(uint32(2), b.Score)
	s.Equal(valueobject.Thrown, b.Status)
	s.Equal(uint32(2), b.FrameNumber)
}

func (s *BowlingSuite) TestThrowSpareGame() {
	b := getPreparedBowlingGame()

	b.Throw(1)
	b.Throw(9)

	s.Equal(uint32(10), b.Score)
	s.Equal(valueobject.Thrown, b.Status)
	s.Equal(uint32(2), b.FrameNumber)
}

func (s *BowlingSuite) TestThrowSpareGameWithBonus() {
	b := getPreparedBowlingGame()

	b.Throw(1)
	b.Throw(9)
	b.Throw(3)
	b.Throw(3)

	s.Equal(uint32(19), b.Score)
}

func (s *BowlingSuite) TestThrowStrikeGame() {
	b := getPreparedBowlingGame()

	b.Throw(10)

	s.Equal(uint32(10), b.Score)
	s.Equal(valueobject.Thrown, b.Status)
	s.Equal(uint32(2), b.FrameNumber)
}

func (s *BowlingSuite) TestThrowStrikeGameWithBonus() {
	b := getPreparedBowlingGame()

	b.Throw(10)
	b.Throw(3)
	b.Throw(3)

	s.Equal(uint32(22), b.Score)
}

func (s *BowlingSuite) TestFinishedWithTenOpenGame() {
	b := getPreparedBowlingGame()

	for _, hit := range getOneHitArray() {
		b.Throw(hit)
	}

	s.Equal(uint32(20), b.Score)
	s.Equal(valueobject.FrameFinished, b.Status)
	s.Equal(uint32(10), b.FrameNumber)
}

func (s *BowlingSuite) TestFinishedWithSpareGame() {
	b := getPreparedBowlingGame()

	for _, hit := range getSpareArray() {
		b.Throw(hit)
	}

	s.Equal(uint32(110), b.Score)
	s.Equal(valueobject.FrameFinished, b.Status)
	s.Equal(uint32(11), b.FrameNumber)
}

func (s *BowlingSuite) TestFinishedWithPerfectGame() {
	b := getPreparedBowlingGame()

	for _, hit := range getStrikeArray() {
		b.Throw(hit)
	}

	s.Equal(uint32(300), b.Score)
	s.Equal(valueobject.FrameFinished, b.Status)
	s.Equal(uint32(12), b.FrameNumber)
}

func getPreparedBowlingGame() *bowling.Bowling {
	return bowling.NewBowling("0", storage.NewInmemEventStorage())
}

func getOneHitArray() []uint32 {
	var ball []uint32
	for i := 0; i < 20; i++ {
		ball = append(ball, 1)
	}
	return ball
}

func getStrikeArray() []uint32 {
	var ball []uint32
	for i := 0; i < 12; i++ {
		ball = append(ball, 10)
	}
	return ball
}

func getSpareArray() []uint32 {
	var ball []uint32
	for i := 0; i < 10; i++ {
		ball = append(ball, 1)
		ball = append(ball, 9)
	}
	ball = append(ball, 1)
	return ball
}
