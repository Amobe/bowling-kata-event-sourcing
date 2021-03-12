package bowling_test

import (
	"testing"

	"github.com/amobe/bowling-kata-event-sourcing/src/bowling"
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
	s.Equal(bowling.Thrown, b.Status)
	s.Equal(uint32(1), b.FrameNumber)
	s.Equal(uint32(0), b.ExtraFrame)
}

func (s *BowlingSuite) TestThrowOpenGame() {
	b := getPreparedBowlingGame()

	b.Throw(1)
	b.Throw(1)

	s.Equal(uint32(2), b.Score)
	s.Equal(bowling.Thrown, b.Status)
	s.Equal(uint32(2), b.FrameNumber)
	s.Equal(uint32(0), b.ExtraFrame)
}

func (s *BowlingSuite) TestThrowSpareGame() {
	b := getPreparedBowlingGame()

	b.Throw(1)
	b.Throw(9)

	s.Equal(uint32(10), b.Score)
	s.Equal(bowling.Thrown, b.Status)
	s.Equal(uint32(2), b.FrameNumber)
	s.Equal(uint32(0), b.ExtraFrame)
}

func (s *BowlingSuite) TestThrowStrikeGame() {
	b := getPreparedBowlingGame()

	b.Throw(10)

	s.Equal(uint32(10), b.Score)
	s.Equal(bowling.Thrown, b.Status)
	s.Equal(uint32(2), b.FrameNumber)
	s.Equal(uint32(0), b.ExtraFrame)
}

func (s *BowlingSuite) TestFinishedWithTenOpenGame() {
	b := getPreparedBowlingGame()

	for _, hit := range getOneHitArray() {
		b.Throw(hit)
	}

	s.Equal(uint32(20), b.Score)
	s.Equal(bowling.FrameFinished, b.Status)
	s.Equal(uint32(10), b.FrameNumber)
	s.Equal(uint32(0), b.ExtraFrame)
}

func (s *BowlingSuite) TestFinishedWithSpareGame() {
	b := getPreparedBowlingGame()

	for _, hit := range getSpareArray() {
		b.Throw(hit)
	}

	s.Equal(uint32(110), b.Score)
	s.Equal(bowling.FrameFinished, b.Status)
	s.Equal(uint32(11), b.FrameNumber)
	s.Equal(uint32(1), b.ExtraFrame)
}

func (s *BowlingSuite) TestFinishedWithPerfectGame() {
	b := getPreparedBowlingGame()

	for _, hit := range getStrikeArray() {
		b.Throw(hit)
	}

	s.Equal(uint32(300), b.Score)
	s.Equal(bowling.FrameFinished, b.Status)
	s.Equal(uint32(12), b.FrameNumber)
	s.Equal(uint32(2), b.ExtraFrame)
}

func getPreparedBowlingGame() *bowling.Bowling {
	return bowling.NewBowling("0")
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
