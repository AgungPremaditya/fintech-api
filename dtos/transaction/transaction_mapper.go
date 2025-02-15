package transaction_dtos

import (
	general_dtos "ledger-system/dtos/general"
	wallet_dtos "ledger-system/dtos/wallet"
	"ledger-system/models"

	"github.com/shopspring/decimal"
)

type Mapper struct {
	walletMapper *wallet_dtos.Mapper
}

func NewTransactionMapper() *Mapper {
	return &Mapper{
		walletMapper: wallet_dtos.NewWalletMapper(),
	}
}

func (m *Mapper) ToTransactionModel(transaction *CreateTransactionDTO, wallet *models.Wallet) (*models.Transaction, error) {
	// Set new transaction
	newTransaction := models.Transaction{
		Type:      string(transaction.Type),
		Amount:    decimal.NewFromFloat(transaction.Amount),
		Reference: transaction.Reference,
		WalletID:  wallet.ID,
	}

	return &newTransaction, nil
}

func (m *Mapper) ToTransactionDetailResponse(transaction *models.Transaction) *DetailTransactionDTO {
	return &DetailTransactionDTO{
		ID:        transaction.ID.String(),
		Type:      models.TransactionType(transaction.Type),
		Amount:    transaction.Amount,
		Reference: transaction.Reference,
		Wallet:    m.walletMapper.ToWalletEmbedDTO(&transaction.Wallet),
	}
}

func (m *Mapper) ToTransactionResponse(transaction *models.Transaction) TransactionIndexDTO {
	return TransactionIndexDTO{
		ID:        transaction.ID.String(),
		Type:      transaction.Type,
		Wallet:    m.walletMapper.ToWalletEmbedDTO(&transaction.Wallet),
		Amount:    transaction.Amount.String(),
		Reference: transaction.Reference,
	}
}

func (m *Mapper) ToTransactionListResponse(transactions []models.Transaction) []TransactionIndexDTO {
	result := make([]TransactionIndexDTO, len(transactions))
	for i, transaction := range transactions {
		result[i] = m.ToTransactionResponse(&transaction)
	}
	return result
}

func (m *Mapper) ToTransactionPaginatedResponse(transaction *[]TransactionIndexDTO, meta *general_dtos.PaginationMeta) *TransactionPaginatedDTO {
	return &TransactionPaginatedDTO{
		Transactions: *transaction,
		Meta:         *meta,
	}
}

func (m *Mapper) ToTransferTransaction(transferPayload *TransferTransactionDTO, fromWallet *models.Wallet, toWallet *models.Wallet) *models.Transaction {
	// Set new transaction
	newTransaction := models.Transaction{
		Type:      string(models.Transfer),
		Amount:    decimal.NewFromFloat(transferPayload.Amount),
		Reference: transferPayload.Reference,
		Wallet:    *fromWallet,
	}

	return &newTransaction
}
