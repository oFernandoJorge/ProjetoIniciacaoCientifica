package room

import "gorm.io/gorm"

// Migrate executa migration do módulo
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Room{})
}