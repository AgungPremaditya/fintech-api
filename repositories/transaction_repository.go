package repositories

import (
	general_dtos "ledger-system/dtos/general"
	"ledger-system/models"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) CreateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	if err := r.db.Create(transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func (r *TransactionRepository) GetTransactionsByWalletId(walletId string, pagination general_dtos.PaginationRequest) ([]models.Transaction, error) {
	var transactions []models.Transaction

	query := r.db.Preload("FromWallet").Preload("ToWallet").Where("from_wallet_id = ? OR to_wallet_id = ?", walletId, walletId)

	if pagination.Search != "" {
		query = query.Where("reference LIKE ?", "%"+pagination.Search+"%")
	}

	offset := (pagination.Page - 1) * pagination.PerPage
	result := query.Offset(offset).Limit(pagination.PerPage).Find(&transactions)

	if result.Error != nil {
		return nil, result.Error
	}

	return transactions, nil
}

func (r *TransactionRepository) GetTransaction(walletId string) (*int64, error) {
	var totalCount int64
	r.db.Model(&models.Transaction{}).
		Where("from_wallet_id = ? OR to_wallet_id = ?", walletId, walletId).
		Count(&totalCount)

	return &totalCount, nil
}
