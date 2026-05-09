package user

import (
	"errors"
)

// UserService concentra as regras de negócio relacionadas aos usuários
//
// Camada responsável por fazr ponte entre:
//
// handlers < - > repositories
type Service struct {
	Repo Repository
}


// NewUserService cria uma nova instância do UserService
func NewService(repo Repository) *Service {
	return &Service{Repo: repo}
}

// CreateUser raliza validações de negócio antes de criar um usuário
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
	return s.Repo.Create(user)
}

// GetAllUsers retorna todos os usuários cadastrados
func (s *Service) FindAll() ([]User, error) {
	return s.Repo.FindAll()
}

// GetUserByID busca usuário por ID
func (s *Service) FindByID(id uint) (*User, error) {
	return s.Repo.FindByID(id)
}
