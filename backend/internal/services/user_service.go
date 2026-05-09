package services

import (
	"errors"

	"ProjetoIniciacaoCientifica/internal/models"
	"ProjetoIniciacaoCientifica/internal/repositories"
)

//UserService concentra as regras de negócio relacionadas aos usuários
//
//Camada responsável por fazr ponte entre:
// handlers < - > repositories
type UserService struct{
	repository *repositories.UserRepository
}

//NewUserService cria uma nova instância do UserService
func NewUserService() *UserService{
	return &UserService{
		repository: repositories.NewUserRepository(),
	}
}

//CreateUser raliza validações de negócio antes de criar um usuário
func (s *UserService) CreateUser (user *models.User) error{

	if user.Name == ""{
		return errors.New("O nome do usuário é obrigatório")
	}

	if user.Email == ""{
		return errors.New("O email do usuário é obrigatório")
	}

	if user.Password == ""{
		return errors.New("A senha do usuário é obrigatória")
	}

	validRoles := map[string]bool{
		"admin": true,
		"coordenador": true,
		"avaliador": true,
		"aluno": true,
	}

	if !validRoles[user.Role]{
		return errors.New("O papel do usuário é inválido")
	}

	//Chama repository para persistir usuário
	return s.repository.Create(user)
}

//GetAllUsers retorna todos os usuários cadastrados
func (s *UserService) GetAllUsers() ([]models.User, error){
	return s.repository.FindAll()
}

//GetUserByID busca usuário por ID
func (s *UserService) GetUserByID(id uint)(*models.User, error){
	return s.repository.FindByID(id)
}