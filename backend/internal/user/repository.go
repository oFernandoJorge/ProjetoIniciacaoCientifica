package user

import (
	"ProjetoIniciacaoCientifica/internal/config"
)

//UserRepository é responsável pelas operações de acesso ao banco relacionadas aos usuários
//
//Não contém regra de negócio
//
//Apenas comunicação com banco de dados
type repository struct{}

//Cria uma nova instância do UserRepository
func NewRepository() Repository{
	return &repository{}
}

//Insere um novo usuário no banco de dados
func (r *repository) Create(user *User) error{
	return config.DB.Create(user).Error
}

//FindAll retorna todos os usuários cadastrados no banco de dados
func (r* repository) FindAll()([]User, error){
	var users []User

	err := config.DB.Find(&users).Error

	return users, err
}

//FindByID retorna um usuário específico com base no ID fornecido
func (r *repository) FindByID(id uint)(*User, error){
	var user User

	err := config.DB.First(&user, id).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}