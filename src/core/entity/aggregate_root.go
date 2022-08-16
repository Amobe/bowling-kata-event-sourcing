package entity

type IAggregateRoot interface {
	When(event DomainEvent)
}

type AggregateRoot struct {
	iAggregateRoot IAggregateRoot
	domainEvents   []DomainEvent
}

func NewAggregateRoot[T IAggregateRoot](t T) *AggregateRoot {
	return &AggregateRoot{
		iAggregateRoot: t,
	}
}

func (a *AggregateRoot) Apply(event DomainEvent) {
	a.iAggregateRoot.When(event)
	a.addDomainEvent(event)
}

func (a *AggregateRoot) addDomainEvent(event DomainEvent) {
	a.domainEvents = append(a.domainEvents, event)
}

func (a *AggregateRoot) DomainEvents() []DomainEvent {
	return a.domainEvents
}

func (a *AggregateRoot) ClearDomainEvents() {
	a.domainEvents = nil
}
