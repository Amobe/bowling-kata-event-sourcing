package event

type Event interface {
	isEvent()
}

type baseEvent struct{}

func (baseEvent) isEvent() {}
