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
	var sentTransactions []models.Transaction
	var receivedTransactions []models.Transaction

	// Fetch sent transactions
	if err := r.db.Where("from_wallet_id = ?", id).Find(&sentTransactions).Error; err != nil {
		return decimal.Zero, err
	}

	// Fetch received transactions
	if err := r.db.Where("to_wallet_id = ?", id).Find(&receivedTransactions).Error; err != nil {
		return decimal.Zero, err
	}

	// Count Balance
	var balance decimal.Decimal
	for _, tx := range sentTransactions {
		balance = balance.Sub(tx.Amount)
	}
	for _, tx := range receivedTransactions {
		balance = balance.Add(tx.Amount)
	}

	return balance, nil
}

func (r *WalletRepository) CreateWallet(wallet *models.Wallet) (*models.Wallet, error) {
	if err := r.db.Create(wallet).Error; err != nil {
		return nil, err
	}
	return wallet, nil
}
