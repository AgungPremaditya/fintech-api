package repositories

import (
	"ledger-system/models"

	"gorm.io/gorm"
)

type LedgerEntryRepository struct {
	db *gorm.DB
}

func NewLedgerEntryRepository(db *gorm.DB) *LedgerEntryRepository {
	return &LedgerEntryRepository{db: db}
}

func (r *LedgerEntryRepository) CreateLedgerEntry(ledgerEntry *models.LedgerEntry) (*models.LedgerEntry, error) {
	if err := r.db.Create(ledgerEntry).Error; err != nil {
		return nil, err
	}

	return ledgerEntry, nil
}

func (r *LedgerEntryRepository) ProcessTransferLedgerEntry(debitEntry *models.LedgerEntry, creditEntry *models.LedgerEntry) (*models.LedgerEntry, *models.LedgerEntry, error) {
	// Transaction
	return debitEntry, creditEntry, r.db.Transaction(func(tx *gorm.DB) error {
		// Create ledger entry debit
		if err := tx.Create(debitEntry).Error; err != nil {
			return err
		}

		// Create ledger entry credit
		if err := tx.Create(creditEntry).Error; err != nil {
			return err
		}

		// Update debit wallet balance
		var debitWallet models.Wallet
		if err := tx.First(&debitWallet, "id = ?", debitEntry.WalletID).Error; err != nil {
			return err
		}

		debitBalance := debitWallet.Balance.Sub(debitEntry.Amount)
		if err := tx.Model(&debitWallet).Update("balance", debitBalance).Error; err != nil {
			return err
		}

		// Update credit wallet balance
		var creditWallet models.Wallet
		if err := tx.First(&creditWallet, "id = ?", creditEntry.WalletID).Error; err != nil {
			return err
		}

		creditBalance := creditWallet.Balance.Add(creditEntry.Amount)
		if err := tx.Model(&creditWallet).Update("balance", creditBalance).Error; err != nil {
			return err
		}

		// Create debit transaction
		debitTransaction := models.Transaction{
			Type:      string(models.Transfer),
			EntryType: string(models.Debit),
			Amount:    debitEntry.Amount,
			Reference: "Transfer to " + creditWallet.Address,
			WalletID:  debitEntry.WalletID,
		}
		if err := tx.Create(&debitTransaction).Error; err != nil {
			return err
		}

		// Create credit transaction
		creditTransaction := models.Transaction{
			Type:      string(models.Transfer),
			EntryType: string(models.Credit),
			Amount:    creditEntry.Amount,
			Reference: "Transfer from " + debitWallet.Address,
			WalletID:  creditEntry.WalletID,
		}
		if err := tx.Create(&creditTransaction).Error; err != nil {
			return err
		}

		return nil
	})
}
