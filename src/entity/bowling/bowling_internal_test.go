package bowling

import (
	"testing"

	"github.com/amobe/bowling-kata-event-sourcing/src/event"
	"github.com/amobe/bowling-kata-event-sourcing/src/valueobject"
	"github.com/stretchr/testify/assert"
)

func Test_calculateGameHit(t *testing.T) {
	hit := uint32(3)
	game := valueobject.BowlingGame{
		FrameNumber: 2,
		Left:        10,
	}
	want := event.NewGameReplacedEvent(2, valueobject.BowlingGame{
		FrameNumber: 2,
		Left:        7,
		Score:       3,
		ThrowNumber: 1,
		Status:      valueobject.Open,
	})
	got := calculateGameHit(hit, game)
	assert.EqualValues(t, want, got)
}

func Test_calculateGameBonus(t *testing.T) {
	hit := uint32(3)
	game := valueobject.BowlingGame{
		FrameNumber: 2,
		Left:        10,
	}
	want := event.NewGameBonusedEvent(2, 0, 0)
	got := calculateGameBonus(hit, game)
	assert.EqualValues(t, want, got)
}

func Test_createNewGame(t *testing.T) {
	type args struct {
		frameNumber uint32
	}
	tests := []struct {
		name string
		args args
		want valueobject.BowlingGame
	}{
		{
			name: "game with extra frame",
			args: args{
				frameNumber: 10,
			},
			want: valueobject.BowlingGame{
				FrameNumber:       10,
				Left:              10,
				WithoutExtraBonus: true,
			},
		},
		{
			name: "game with no extra frame",
			args: args{
				frameNumber: 9,
			},
			want: valueobject.BowlingGame{
				FrameNumber:       9,
				Left:              10,
				WithoutExtraBonus: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, createNewGame(tt.args.frameNumber))
		})
	}
}

func Test_calculateScore(t *testing.T) {
	games := map[uint32]valueobject.BowlingGame{
		0: {
			Score: 1,
		},
		1: {
			Score: 2,
		},
	}
	want := uint32(3)
	got := calculateScore(games)
	assert.Equal(t, want, got)
}

func Test_hasExtraFrame(t *testing.T) {
	type args struct {
		frameNumber uint32
		game        valueobject.BowlingGame
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "don't have extea frame - open end",
			args: args{
				frameNumber: 10,
				game: valueobject.BowlingGame{
					Status: valueobject.Open,
				},
			},
			want: false,
		},
		{
			name: "don't have extea frame - strike twice",
			args: args{
				frameNumber: 12,
				game: valueobject.BowlingGame{
					Status: valueobject.Strike,
				},
			},
			want: false,
		},
		{
			name: "have extra frame - spare on 11th frame",
			args: args{
				frameNumber: 11,
				game: valueobject.BowlingGame{
					Status: valueobject.Spare,
				},
			},
			want: true,
		},
		{
			name: "have extra frame - strike on 10th frame",
			args: args{
				frameNumber: 10,
				game: valueobject.BowlingGame{
					Status: valueobject.Strike,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, hasExtraFrame(tt.args.frameNumber, tt.args.game))
		})
	}
}
