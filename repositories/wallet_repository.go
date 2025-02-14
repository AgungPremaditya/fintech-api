package repositories

import (
	"fmt"
	"ledger-system/models"

	"gorm.io/gorm"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (r *WalletRepository) GetWallets() ([]models.Wallet, error) {
	var wallets []models.Wallet
	result := r.db.Preload("User").Find(&wallets)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch wallets: %v", result.Error)
	}

	return wallets, nil
}

func (r *WalletRepository) GetWallet(id string) (models.Wallet, error) {
	var wallet models.Wallet
	result := r.db.Preload("User").First(&wallet, "id = ?", id)

	if result.Error != nil {
		return models.Wallet{}, fmt.Errorf("failed to fetch wallet: %v", result.Error)
	}

	return wallet, nil
}
