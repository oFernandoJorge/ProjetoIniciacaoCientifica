package evaluator

import "errors"

// Service contém regras de negócio
type Service struct {
	repo Repository
}

// NewService injeta repository
func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// Create cria avaliador
func (s *Service) Create(evaluator *Evaluator) error {

	if evaluator.Name == "" {
		return errors.New("o nome do avaliador é obrigatório")
	}

	if evaluator.Email == "" {
		return errors.New("o email do avaliador é obrigatório")
	}

	if evaluator.Course == "" {
		return errors.New("o curso do avaliador é obrigatório")
	}

	if evaluator.KnowledgeArea == "" {
		return errors.New("a área do avaliador é obrigatória")
	}

	if evaluator.MaxPresentations <= 0 {
		return errors.New("a quantidade máxima de apresentações deve ser maior que zero")
	}

	validPresentationTypes := map[string]bool{
		"ORAL": true,
		"E-POSTER": true,
		"BOTH": true,
	}

	if !validPresentationTypes[evaluator.AcceptedPresentationType] {
		return errors.New("tipo de apresentação inválido")
	}

	return s.repo.Create(evaluator)
}

// FindAll retorna avaliadores
func (s *Service) FindAll() ([]Evaluator, error) {

	return s.repo.FindAll()
}