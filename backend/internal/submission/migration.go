package submission

import "gorm.io/gorm"

// Migrate executa migration do módulo submission
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Submission{})
}