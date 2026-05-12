package user

import "gorm.io/gorm"

// User representa usuário do sistema
type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Role     string
}
