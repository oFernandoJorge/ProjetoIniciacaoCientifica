package repositories

import (
	"ProjetoIniciacaoCientifica/internal/config"
	"ProjetoIniciacaoCientifica/internal/models"
)

//UserRepository é responsável pelas operações de acesso ao banco relacionadas aos usuários
//
//Não contém regra de negócio
//
//Apenas comunicação com banco de dados
type UserRepository struct{}

//Cria uma nova instância do UserRepository
func NewUserRepository() *UserRepository{
	return &UserRepository{}
}

//Insere um novo usuário no banco de dados
func (r *UserRepository) Create(user *models.User) error{
	return config.DB.Create(user).Error
}

//FindAll retorna todos os usuários cadastrados no banco de dados
func (r* UserRepository) FindAll()([]models.User, error){
	var users []models.User

	err := config.DB.Find(&users).Error

	return users, err
}

//FindByID retorna um usuário específico com base no ID fornecido
func (r *UserRepository) FindByID(id uint)(*models.User, error){
	var user models.User

	err := config.DB.First(&user, id).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}