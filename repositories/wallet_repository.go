package repositories

import (
	"fmt"
	"ledger-system/models"

	"github.com/shopspring/decimal"
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

func (r *WalletRepository) GetWalletBalance(id string) (decimal.Decimal, error) {
	var sentTotal, receivedTotal decimal.Decimal

	// Sum sent amount
	if err := r.db.Model(&models.Transaction{}).
		Where("from_wallet_id = ?", id).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&sentTotal).Error; err != nil {
		return decimal.Zero, err
	}

	// Sum received amount
	if err := r.db.Model(&models.Transaction{}).
		Where("to_wallet_id = ?", id).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&receivedTotal).Error; err != nil {
		return decimal.Zero, err
	}

	// Calculate balance
	balance := receivedTotal.Sub(sentTotal)

	return balance, nil
}

func (r *WalletRepository) CreateWallet(wallet *models.Wallet) (*models.Wallet, error) {
	if err := r.db.Create(wallet).Error; err != nil {
		return nil, err
	}
	return wallet, nil
}
