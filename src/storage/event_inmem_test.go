package storage

import (
	"testing"

	"github.com/amobe/bowling-kata-event-sourcing/src/event"
	"github.com/amobe/bowling-kata-event-sourcing/src/valueobject"
	"github.com/stretchr/testify/assert"
)

func TestNewInmemEventStorage(t *testing.T) {
	s := NewInmemEventStorage()
	assert.NotNil(t, s)
	assert.NotNil(t, s.changes)
}

func Test_storage_Get(t *testing.T) {
	type fields struct {
		changes map[string][]event.Event
	}
	type args struct {
		id string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      []event.Event
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "get",
			fields: fields{
				changes: map[string][]event.Event{
					"1": {
						event.NewThrownEvent(valueobject.Thrown, 1),
					},
				},
			},
			args: args{
				"1",
			},
			want: []event.Event{
				event.NewThrownEvent(valueobject.Thrown, 1),
			},
			assertion: assert.NoError,
		},
		{
			name: "record not found",
			fields: fields{
				changes: map[string][]event.Event{},
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
		changes map[string][]event.Event
	}
	type args struct {
		id  string
		evs []event.Event
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      map[string][]event.Event
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "append",
			fields: fields{
				changes: map[string][]event.Event{},
			},
			args: args{
				id: "1",
				evs: []event.Event{
					event.NewThrownEvent(valueobject.Thrown, 2),
				},
			},
			want: map[string][]event.Event{
				"1": {event.NewThrownEvent(valueobject.Thrown, 2)},
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
