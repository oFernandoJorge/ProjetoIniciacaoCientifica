package user

import (
	"errors"
)

// Service concentra as regras de negócio relacionadas aos usuários
//
// Camada responsável por fazr ponte entre:
//
// handlers < - > repositories
type Service struct {
	repo Repository
}

// NewService cria uma nova instância do UserService
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// Create raliza validações de negócio antes de criar um usuário
func (s *Service) Create(user *User) error {

	if user.Name == "" {
		return errors.New("O nome do usuário é obrigatório")
	}

	if user.Email == "" {
		return errors.New("O email do usuário é obrigatório")
	}

	if user.Password == "" {
		return errors.New("A senha do usuário é obrigatória")
	}

	validRoles := map[string]bool{
		"admin":       true,
		"coordenador": true,
		"avaliador":   true,
		"aluno":       true,
	}

	if !validRoles[user.Role] {
		return errors.New("O papel do usuário é inválido")
	}

	//Chama repository para persistir usuário
	return s.repo.Create(user)
}

// FindAll retorna todos os usuários cadastrados
func (s *Service) FindAll() ([]UserResponse, error) {

	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []UserResponse
	for _, u := range users {
		responses = append(responses, UserResponse{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
			Role:  u.Role,
		})
	}
	return responses, nil
}

// FindByID busca usuário por ID
func (s *Service) FindByID(id uint) (*UserResponse, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}
