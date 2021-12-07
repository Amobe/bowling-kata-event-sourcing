package event

//go:generate mockgen -destination mocks/mock_repository.go -package mocks -source repository.go
type Repository interface {
	Get(id string) ([]Event, error)
	Append(id string, evs ...Event) error
}
