package user

// Repository define comportamentos esperados
// da camada de persistência
type Repository interface {
	Create(user *User) error
	FindAll() ([]User, error)
	FindByID(id uint) (*User, error)
}
