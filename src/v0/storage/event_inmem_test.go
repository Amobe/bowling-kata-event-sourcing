package storage

import (
	"testing"

	event2 "github.com/amobe/bowling-kata-event-sourcing/src/v0/event"
	"github.com/amobe/bowling-kata-event-sourcing/src/v0/valueobject"
	"github.com/stretchr/testify/assert"
)

func TestNewInmemEventStorage(t *testing.T) {
	s := NewInmemEventStorage()
	assert.NotNil(t, s)
	assert.NotNil(t, s.changes)
}

func Test_storage_Get(t *testing.T) {
	type fields struct {
		changes map[string][]event2.Event
	}
	type args struct {
		id string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      []event2.Event
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "get",
			fields: fields{
				changes: map[string][]event2.Event{
					"1": {
						event2.NewThrownEvent(valueobject.Thrown, 1),
					},
				},
			},
			args: args{
				"1",
			},
			want: []event2.Event{
				event2.NewThrownEvent(valueobject.Thrown, 1),
			},
			assertion: assert.NoError,
		},
		{
			name: "record not found",
			fields: fields{
				changes: map[string][]event2.Event{},
			},
			args: args{
				"1",
			},
			want:      nil,
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &storage{
				changes: tt.fields.changes,
			}
			got, err := s.Get(tt.args.id)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_storage_Append(t *testing.T) {
	type fields struct {
		changes map[string][]event2.Event
	}
	type args struct {
		id  string
		evs []event2.Event
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      map[string][]event2.Event
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "append",
			fields: fields{
				changes: map[string][]event2.Event{},
			},
			args: args{
				id: "1",
				evs: []event2.Event{
					event2.NewThrownEvent(valueobject.Thrown, 2),
				},
			},
			want: map[string][]event2.Event{
				"1": {event2.NewThrownEvent(valueobject.Thrown, 2)},
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &storage{
				changes: tt.fields.changes,
			}
			tt.assertion(t, s.Append(tt.args.id, tt.args.evs...))
			assert.Equal(t, tt.want, s.changes)
		})
	}
}
