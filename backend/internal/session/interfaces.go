package session

// Repository define comportamentos
type Repository interface {
	Create(session *Session) error
	FindAll() ([]Session, error)
}