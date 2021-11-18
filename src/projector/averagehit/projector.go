package averagehit

import (
	"github.com/amobe/bowling-kata-event-sourcing/src/entity/bowling"
	"github.com/amobe/bowling-kata-event-sourcing/src/event"
)

type projector struct {
	count      uint64
	averageHit uint64
}

func NewPojector() *projector {
	return &projector{}
}

func (p *projector) Project(ev event.Event) {
	switch ev.Data().(type) {
	case *event.GameReplacedEvent:
		data := ev.Data().(*event.GameReplacedEvent)
		hit := bowling.StandardPins - data.Game.Left

		p.count++
		p.averageHit = (p.averageHit*(p.count-1) + uint64(hit)) / p.count
	}
}
