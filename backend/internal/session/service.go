package session

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

// Create cria sessão
func (s *Service) Create(session *Session) error {

	if session.RoomID == 0 {
		return errors.New("room_id é obrigatório")
	}

	if session.KnowledgeArea == "" {
		return errors.New("área é obrigatória")
	}

	if session.PresentationType == "" {
		return errors.New("tipo de apresentação é obrigatório")
	}

	if session.MaxCapacity <= 0 {
		return errors.New("capacidade inválida")
	}

	return s.repo.Create(session)
}

// FindAll retorna sessões
func (s *Service) FindAll() ([]Session, error) {

	return s.repo.FindAll()
}