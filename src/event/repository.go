package event

type Repository interface {
	Append(id string, ev Event) error
}
