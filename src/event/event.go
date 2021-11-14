package event

type Event interface {
	EventName() string
}

type baseEvent struct {
	name string
}

func (b baseEvent) EventName() string {
	return b.name
}
