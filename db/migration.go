package db

import (
	"ledger-system/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	// Enable UUID extension
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	// Import models and migrate them
	return db.AutoMigrate(
		&models.User{},
		&models.Wallet{},
		&models.Transaction{},
	)
}

func Revert(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&models.User{},
		&models.Wallet{},
		&models.Transaction{},
	)
}
