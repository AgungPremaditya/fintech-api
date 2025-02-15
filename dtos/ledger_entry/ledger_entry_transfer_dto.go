package ledgerentry_dtos

import (
	"github.com/shopspring/decimal"
)

type LedgerEntryTransferDTO struct {
	ID       string          `json:"id"`
	Type     string          `json:"type"`
	Amount   decimal.Decimal `json:"amount"`
	Balance  decimal.Decimal `json:"balance"`
	WalletID string          `json:"wallet_id"`
}

type DetailTransferTransactionDTO struct {
	DebitEntry  LedgerEntryTransferDTO `json:"debit_entry"`
	CreditEntry LedgerEntryTransferDTO `json:"credit_entry"`
}
