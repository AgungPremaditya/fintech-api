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
	}

	transactionType := models.TransactionType(transaction.Type)

	// Check transaction type and set wallet
	if transactionType == models.Deposit {
		newTransaction.ToWalletID.UUID = wallet.ID
		newTransaction.ToWallet = *wallet
	} else if transactionType == models.Withdraw {
		newTransaction.FromWalletID.UUID = wallet.ID
		newTransaction.FromWallet = *wallet
	}

	return &newTransaction, nil
}

func (m *Mapper) ToTransactionDetailResponse(transaction *models.Transaction) *DetailTransactionDTO {
	return &DetailTransactionDTO{
		ID:         transaction.ID.String(),
		Type:       models.TransactionType(transaction.Type),
		Amount:     transaction.Amount,
		Reference:  transaction.Reference,
		FromWallet: *m.walletMapper.ToWalletEmbedDTO(&transaction.FromWallet),
		ToWallet:   *m.walletMapper.ToWalletEmbedDTO(&transaction.ToWallet),
	}
}

func (m *Mapper) ToTransactionResponse(transaction *models.Transaction) TransactionIndexDTO {
	return TransactionIndexDTO{
		ID:         transaction.ID.String(),
		Type:       transaction.Type,
		FromWallet: m.walletMapper.ToWalletEmbedDTO(&transaction.FromWallet),
		ToWallet:   m.walletMapper.ToWalletEmbedDTO(&transaction.ToWallet),
		Amount:     transaction.Amount.String(),
		Reference:  transaction.Reference,
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
