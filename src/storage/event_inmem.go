package storage

import (
	"fmt"

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

func (s *storage) Get(id string) ([]event.Event, error) {
	evs, ok := s.changes[id]
	if !ok {
		return nil, fmt.Errorf("record not found")
	}
	return evs, nil
}

func (s *storage) Append(id string, evs ...event.Event) error {
	s.changes[id] = append(s.changes[id], evs...)
	return nil
}
