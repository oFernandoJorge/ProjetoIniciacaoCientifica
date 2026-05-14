package submission

import "ProjetoIniciacaoCientifica/internal/config"

// repository implementa acesso ao banco
type repository struct{}

// NewRepository cria repository concreto
func NewRepository() Repository {
	return &repository{}
}

// Create cria submissão
func (r *repository) Create(submission *Submission) error {
	return config.DB.Create(submission).Error
}

// FindAll retorna submissões
func (r *repository) FindAll() ([]Submission, error) {

	var submissions []Submission

	err := config.DB.Find(&submissions).Error

	return submissions, err
}