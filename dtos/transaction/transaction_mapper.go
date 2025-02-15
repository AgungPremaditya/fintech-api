package transaction_dtos

import (
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
		FromWallet: m.walletMapper.ToWalletEmbedDTO(transaction.FromWallet),
		ToWallet:   m.walletMapper.ToWalletEmbedDTO(transaction.ToWallet),
	}
}
