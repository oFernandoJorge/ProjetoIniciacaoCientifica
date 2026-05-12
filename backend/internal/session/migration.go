package session

import "gorm.io/gorm"

// Migrate executa migration
func Migrate(db *gorm.DB) error {

	return db.AutoMigrate(&Session{})
}