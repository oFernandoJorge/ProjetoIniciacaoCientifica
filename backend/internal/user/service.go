package user

import (
	"ProjetoIniciacaoCientifica/internal/auth"
	"errors"
)

// Service contém regras de negócio
type Service struct {
	repo Repository
}

// NewService injeta dependência
func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// Create cria novo usuário
func (s *Service) Create(user *User) error {
	if user.Name == "" {
		return errors.New("nome obrigatório")
	}
	if user.Email == "" {
		return errors.New("email obrigatório")

	}
	validRoles := map[string]bool{
		"admin":       true,
		"aluno":       true,
		"avaliador":   true,
		"coordenador": true,
	}
	if !validRoles[user.Role] {
		return errors.New("role inválida")
	}
	// Gera hash seguro da senha
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return s.repo.Create(user)
}

// FindAll retorna todos usuários
func (s *Service) FindAll() ([]UserResponse, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	var response []UserResponse
	for _, u := range users {
		response = append(response, UserResponse{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
			Role:  u.Role,
		})
	}
	return response, nil
}
