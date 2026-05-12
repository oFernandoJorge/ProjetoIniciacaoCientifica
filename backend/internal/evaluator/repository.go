package evaluator

import "ProjetoIniciacaoCientifica/internal/config"

type repository struct{}

// NewRepository cria repository concreto
func NewRepository() Repository {
	return &repository{}
}

// Create cria avaliador
func (r *repository) Create(evaluator *Evaluator) error {

	return config.DB.Create(evaluator).Error
}

// FindAll retorna avaliadores
func (r *repository) FindAll() ([]Evaluator, error) {

	var evaluators []Evaluator

	err := config.DB.Find(&evaluators).Error

	return evaluators, err
}