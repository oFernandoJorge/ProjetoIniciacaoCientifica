package user

import "gorm.io/gorm"

// Migrate executa migration do módulo user
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
