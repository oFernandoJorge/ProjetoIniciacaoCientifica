package user

import "gorm.io/gorm"

// Migrate é responsabilidade do módulo User
// Criar/atualizar tabelas relacionadas ao domínio User
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
	)
}