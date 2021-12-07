package service

import (
	"fmt"

	"github.com/amobe/bowling-kata-event-sourcing/src/entity/bowling"
	"github.com/amobe/bowling-kata-event-sourcing/src/event"
)

type Bowling interface {
	Throw(id string, hit uint32) error
}

var _ Bowling = &bowlingService{}

type bowlingService struct {
	repo event.Repository
}

func (s *bowlingService) Throw(id string, hit uint32) error {
	evs, err := s.repo.Get(id)
	if err != nil {
		return fmt.Errorf("failed to get from repo: %w", err)
	}
	newEvs := bowling.Projector(id, evs).Throw(hit)
	if err := s.repo.Append(id, newEvs...); err != nil {
		return fmt.Errorf("failed to append events: %w", err)
	}
	return nil
}
