package user

import "ProjetoIniciacaoCientifica/internal/config"

// repository é implementação concreta usando GORM
type repository struct{}

// NewRepository cria repository concreto
func NewRepository() Repository {
	return &repository{}
}
func (r *repository) Create(user *User) error {
	return config.DB.Create(user).Error
}
func (r *repository) FindAll() ([]User, error) {
	var users []User
	err := config.DB.Find(&users).Error
	return users, err
}
func (r *repository) FindByID(id uint) (*User, error) {
	var user User
	err := config.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
