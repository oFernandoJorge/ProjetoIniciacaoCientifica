package models

import(
	"gorm.io/gorm"
)

//User representa usuários do sistema
//
//Pode assumir diferente papeis:
//
// - admin
//
// - coordenador
//
// - avaliador
//
// - aluno
type User struct {

	//Campos padrão adicionados pelo GORM
	//
	//ID
	//
	//CreatedAt
	//
	//UpdatedAt
	//
	//DeletedAt
	gorm.Model

	//Nome completo do usuário
	Name     string `json:"name"`
	//Email do usuário, deve ser único
	Email    string `json:"email" gorm:"unique"`
	//Senha do usuário, deve ser armazenada de forma segura (hash)
	Password string `json:"password"`
	//Papel/Permissão do usuário no sistema
	Role	 string `json:"role"`
}
