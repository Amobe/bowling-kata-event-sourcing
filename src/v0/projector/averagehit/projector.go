package averagehit

import (
	"github.com/amobe/bowling-kata-event-sourcing/src/v0/entity/bowling"
	event2 "github.com/amobe/bowling-kata-event-sourcing/src/v0/event"
)

type projector struct {
	count      uint64
	averageHit uint64
}

func NewPojector() *projector {
	return &projector{}
}

func (p *projector) Project(ev event2.Event) {
	switch ev.Data().(type) {
	case *event2.GameReplacedEvent:
		data := ev.Data().(*event2.GameReplacedEvent)
		hit := bowling.StandardPins - data.Game.Left

		p.count++
		p.averageHit = (p.averageHit*(p.count-1) + uint64(hit)) / p.count
	}
}
