package bowling

import (
	"testing"

	"github.com/amobe/bowling-kata-event-sourcing/src/valueobject"
	"github.com/stretchr/testify/assert"
)

func Test_bonus(t *testing.T) {
	type args struct {
		game valueobject.BowlingGame
		pins uint32
	}
	tests := []struct {
		name string
		args args
		want valueobject.BowlingGame
	}{
		{
			name: "game with no extra bonus",
			args: args{
				game: valueobject.BowlingGame{
					ExtraBonus: 0,
				},
				pins: 1,
			},
			want: valueobject.BowlingGame{
				ExtraBonus: 0,
			},
		},
		{
			name: "game with extra bonus",
			args: args{
				game: valueobject.BowlingGame{
					ExtraBonus: 1,
				},
				pins: 1,
			},
			want: valueobject.BowlingGame{
				ExtraBonus: 0,
				Score:      1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, bonus(tt.args.game, tt.args.pins))
		})
	}
}

func Test_isStrike(t *testing.T) {
	type args struct {
		game valueobject.BowlingGame
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "clean with 2 throws",
			args: args{
				game: valueobject.BowlingGame{
					ThrowNumber: 2,
					Left:        0,
				},
			},
			want: false,
		},
		{
			name: "did not clean the game with 1 throw",
			args: args{
				game: valueobject.BowlingGame{
					ThrowNumber: 1,
					Left:        1,
				},
			},
		},
		{
			name: "clean with 1 throw",
			args: args{
				game: valueobject.BowlingGame{
					ThrowNumber: 1,
					Left:        0,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isStrike(tt.args.game))
		})
	}
}

func Test_isSpare(t *testing.T) {
	type args struct {
		game valueobject.BowlingGame
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "clean with 2 throws",
			args: args{
				game: valueobject.BowlingGame{
					ThrowNumber: 2,
					Left:        0,
				},
			},
			want: true,
		},
		{
			name: "did not clean with 2 throws",
			args: args{
				game: valueobject.BowlingGame{
					ThrowNumber: 2,
					Left:        1,
				},
			},
			want: false,
		},
		{
			name: "clean with 1 throws",
			args: args{
				game: valueobject.BowlingGame{
					ThrowNumber: 1,
					Left:        0,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isSpare(tt.args.game))
		})
	}
}

func Test_gainExtraBonus(t *testing.T) {
	type args struct {
		game       valueobject.BowlingGame
		extraBonus uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "game without extra bonus",
			args: args{
				game: valueobject.BowlingGame{
					WithoutExtraBonus: false,
				},
				extraBonus: 1,
			},
			want: 1,
		},
		{
			name: "game with extra bonus",
			args: args{
				game: valueobject.BowlingGame{
					WithoutExtraBonus: true,
				},
				extraBonus: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, gainExtraBonus(tt.args.game, tt.args.extraBonus))
		})
	}
}

func Test_getStatusBonus(t *testing.T) {
	type args struct {
		game valueobject.BowlingGame
	}
	tests := []struct {
		name           string
		args           args
		wantS          valueobject.BowlingGameStatus
		wantExtraBonus uint32
	}{
		{
			name: "a spare game with extra bonus",
			args: args{
				game: valueobject.BowlingGame{
					ThrowNumber:       2,
					Left:              0,
					WithoutExtraBonus: false,
				},
			},
			wantS:          valueobject.Spare,
			wantExtraBonus: 1,
		},
		{
			name: "a spare game without extra bonus",
			args: args{
				game: valueobject.BowlingGame{
					ThrowNumber:       2,
					Left:              0,
					WithoutExtraBonus: true,
				},
			},
			wantS:          valueobject.Spare,
			wantExtraBonus: 0,
		},
		{
			name: "a strike game with extra bonus",
			args: args{
				game: valueobject.BowlingGame{
					ThrowNumber:       1,
					Left:              0,
					WithoutExtraBonus: false,
				},
			},
			wantS:          valueobject.Strike,
			wantExtraBonus: 2,
		},
		{
			name: "a strike game without extra bonus",
			args: args{
				game: valueobject.BowlingGame{
					ThrowNumber:       1,
					Left:              0,
					WithoutExtraBonus: true,
				},
			},
			wantS:          valueobject.Strike,
			wantExtraBonus: 0,
		},
		{
			name: "a normal game",
			args: args{
				game: valueobject.BowlingGame{
					ThrowNumber: 1,
					Left:        1,
				},
			},
			wantS:          valueobject.Open,
			wantExtraBonus: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, gotExtraBonus := getStatusBonus(tt.args.game)
			assert.Equal(t, tt.wantS, gotS)
			assert.Equal(t, tt.wantExtraBonus, gotExtraBonus)
		})
	}
}
