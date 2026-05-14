package room

// Repository define comportamentos da sala
type Repository interface {
	Create(room *Room) error
	FindAll() ([]Room, error)
}