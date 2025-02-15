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
	return transaction, r.db.Transaction(func(tx *gorm.DB) error {
		// Update wallet balance
		var wallet models.Wallet
		if err := tx.First(&wallet, "id = ?", transaction.WalletID).Error; err != nil {
			return err
		}

		if transaction.Type == string(models.Withdraw) {
			wallet.Balance = wallet.Balance.Sub(transaction.Amount)
		} else {
			wallet.Balance = wallet.Balance.Add(transaction.Amount)
		}

		if err := tx.Save(&wallet).Error; err != nil {
			return err
		}

		// Create ledger entry
		if err := tx.Create(&models.LedgerEntry{
			WalletID: transaction.WalletID,
			Amount:   transaction.Amount,
			Type:     string(transaction.EntryType),
			Balance:  wallet.Balance,
		}).Error; err != nil {
			return err
		}

		// Create transaction
		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *TransactionRepository) GetTransactionsByWalletId(walletId string, pagination general_dtos.PaginationRequest) ([]models.Transaction, error) {
	var transactions []models.Transaction

	query := r.db.Preload("Wallet").Where("wallet_id = ?", walletId)

	if pagination.Search != "" {
		query = query.Where("reference LIKE ?", "%"+pagination.Search+"%")
	}

	query = query.Order("created_at DESC")

	offset := (pagination.Page - 1) * pagination.PerPage
	result := query.Offset(offset).Limit(pagination.PerPage).Find(&transactions)

	if result.Error != nil {
		return nil, result.Error
	}

	return transactions, nil
}

func (r *TransactionRepository) CountTransaction(walletId string) (*int64, error) {
	var totalCount int64
	r.db.Model(&models.Transaction{}).
		Where("from_wallet_id = ? OR to_wallet_id = ?", walletId, walletId).
		Count(&totalCount)

	return &totalCount, nil
}
