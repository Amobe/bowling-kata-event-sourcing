package storage

import (
	"github.com/amobe/bowling-kata-event-sourcing/src/event"
)

type storage struct {
	changes map[string][]event.Event
}

func NewInmemEventStorage() *storage {
	return &storage{
		changes: make(map[string][]event.Event),
	}
}

func (s *storage) Append(id string, ev event.Event) error {
	s.changes[id] = append(s.changes[id], ev)
	return nil
}
