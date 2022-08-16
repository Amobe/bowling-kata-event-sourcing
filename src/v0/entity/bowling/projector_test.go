package bowling

import (
	"testing"

	event2 "github.com/amobe/bowling-kata-event-sourcing/src/v0/event"
	"github.com/amobe/bowling-kata-event-sourcing/src/v0/valueobject"
	"github.com/stretchr/testify/assert"
)

func TestProjector(t *testing.T) {
	evs := []event2.Event{
		event2.NewThrownEvent(valueobject.Thrown, 3),
		event2.NewReloadedEvent(valueobject.FrameFinished, 4),
	}
	expected := &Bowling{
		ID:          "2",
		Status:      valueobject.FrameFinished,
		Score:       3,
		FrameNumber: 4,
		Games:       make(map[uint32]valueobject.BowlingGame),
	}
	actual := Projector("2", evs)
	assert.EqualValues(t, expected, actual)
}

func Test_applyThrownEvent(t *testing.T) {
	ev := &event2.ThrownEvent{
		Status: valueobject.Thrown,
		Score:  3,
	}
	actual := &Bowling{}
	expected := &Bowling{
		Status: valueobject.Thrown,
		Score:  3,
	}

	applyThrownEvent(ev, actual)
	assert.EqualValues(t, expected, actual)
}

func Test_applyGameReplacedEvent(t *testing.T) {
	ev := &event2.GameReplacedEvent{
		FrameNumber: 2,
		Game: valueobject.BowlingGame{
			FrameNumber: 2,
		},
	}
	actual := &Bowling{
		Games: make(map[uint32]valueobject.BowlingGame),
	}
	expected := &Bowling{
		Games: map[uint32]valueobject.BowlingGame{
			2: {
				FrameNumber: 2,
			},
		},
	}
	applyGameReplacedEvent(ev, actual)
	assert.EqualValues(t, expected, actual)
}

func Test_applyGameBonusedEvent(t *testing.T) {
	ev := &event2.GameBonusedEvent{
		FrameNumber: 2,
		Score:       3,
		ExtraBonus:  4,
	}
	actual := &Bowling{
		Games: map[uint32]valueobject.BowlingGame{
			2: {
				FrameNumber: 2,
			},
		},
	}
	expected := &Bowling{
		Games: map[uint32]valueobject.BowlingGame{
			2: {
				FrameNumber: 2,
				Score:       3,
				ExtraBonus:  4,
			},
		},
	}
	applyGameBonusedEvent(ev, actual)
	assert.EqualValues(t, expected, actual)
}

func Test_applyReloadedEvent(t *testing.T) {
	ev := &event2.ReloadedEvent{
		Status:      valueobject.FrameFinished,
		FrameNumber: 3,
	}
	actual := &Bowling{}
	expected := &Bowling{
		Status:      valueobject.FrameFinished,
		FrameNumber: 3,
	}
	applyReloadedEvent(ev, actual)
	assert.EqualValues(t, expected, actual)
}

func Test_on(t *testing.T) {
	type args struct {
		ev event2.Event
		b  *Bowling
	}
	tests := []struct {
		name string
		args args
		want *Bowling
	}{
		{
			name: "apply thrown event",
			args: args{
				ev: event2.NewThrownEvent(valueobject.Thrown, 3),
				b:  &Bowling{},
			},
			want: &Bowling{
				Status: valueobject.Thrown,
				Score:  3,
			},
		},
		{
			name: "apply game replaced event",
			args: args{
				ev: event2.NewGameReplacedEvent(2, valueobject.BowlingGame{FrameNumber: 2}),
				b: &Bowling{
					Games: make(map[uint32]valueobject.BowlingGame),
				},
			},
			want: &Bowling{
				Games: map[uint32]valueobject.BowlingGame{
					2: {
						FrameNumber: 2,
					},
				},
			},
		},
		{
			name: "apply game bonused event",
			args: args{
				ev: event2.NewGameBonusedEvent(2, 3, 4),
				b: &Bowling{
					Games: map[uint32]valueobject.BowlingGame{
						2: {
							FrameNumber: 2,
						},
					},
				},
			},
			want: &Bowling{
				Games: map[uint32]valueobject.BowlingGame{
					2: {
						FrameNumber: 2,
						Score:       3,
						ExtraBonus:  4,
					},
				},
			},
		},
		{
			name: "apply reloaded event",
			args: args{
				ev: event2.NewReloadedEvent(valueobject.FrameFinished, 3),
				b:  &Bowling{},
			},
			want: &Bowling{
				Status:      valueobject.FrameFinished,
				FrameNumber: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			on(tt.args.ev, tt.args.b)
			assert.EqualValues(t, tt.want, tt.args.b)
		})
	}
}
