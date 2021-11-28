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
