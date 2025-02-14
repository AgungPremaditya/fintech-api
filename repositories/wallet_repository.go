package repositories

import (
	"fmt"
	"ledger-system/db"
	"ledger-system/models"

	"gorm.io/gorm"
)

var connection *gorm.DB

func Init() {
	connection = db.Init()
}
func GetWallets() ([]models.Wallet, error) {
	var wallets []models.Wallet
	result := connection.Find(&wallets)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch wallets: %v", result.Error)
	}

	return wallets, nil
}

func GetWallet(id string) (models.Wallet, error) {
	var wallet models.Wallet
	result := connection.First(&wallet, "id = ?", id)

	if result.Error != nil {
		return models.Wallet{}, fmt.Errorf("failed to fetch wallet: %v", result.Error)
	}

	return wallet, nil
}
