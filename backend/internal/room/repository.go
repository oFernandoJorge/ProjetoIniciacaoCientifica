package room

import "ProjetoIniciacaoCientifica/internal/config"

type repository struct{}

// NewRepository cria repository concreto
func NewRepository() Repository {
	return &repository{}
}

// Create cria sala
func (r *repository) Create(room *Room) error {
	return config.DB.Create(room).Error
}

// FindAll retorna salas
func (r *repository) FindAll() ([]Room, error) {

	var rooms []Room

	err := config.DB.Find(&rooms).Error

	return rooms, err
}