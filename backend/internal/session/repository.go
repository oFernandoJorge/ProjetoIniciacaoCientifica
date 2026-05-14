package session

import "ProjetoIniciacaoCientifica/internal/config"

type repository struct{}

// NewRepository cria repository concreto
func NewRepository() Repository {
	return &repository{}
}

// Create cria sessão
func (r *repository) Create(session *Session) error {

	return config.DB.Create(session).Error
}

// FindAll retorna sessões
func (r *repository) FindAll() ([]Session, error) {

	var sessions []Session

	err := config.DB.Find(&sessions).Error

	return sessions, err
}