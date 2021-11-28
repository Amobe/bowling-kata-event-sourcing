package handler

import (
	"fmt"

	"github.com/amobe/bowling-kata-event-sourcing/src/entity/bowling"
	"github.com/amobe/bowling-kata-event-sourcing/src/event"
)

type Handler interface {
	Roll(hit uint32) error
}

func NewHandler(r event.Repository) Handler {
	return &handler{
		game: bowling.NewBowling("0", r),
	}
}

var _ Handler = &handler{}

type handler struct {
	game *bowling.Bowling
}

func (h *handler) Roll(hit uint32) error {
	if hit > 10 {
		return fmt.Errorf("roll handler: %d is out of range", hit)
	}
	fmt.Println(hit)
	h.game.Throw(hit)
	fmt.Println(h.game.Score)
	return nil
}
