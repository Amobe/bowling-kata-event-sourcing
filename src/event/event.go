package event

type Event interface {
	ID() string
	Name() string
	Data() interface{}
}

type event struct {
	id   string
	name string
	data interface{}
}

func (e event) ID() string {
	return e.id
}

func (e event) Name() string {
	return e.name
}

func (e event) Data() interface{} {
	return e.data
}
