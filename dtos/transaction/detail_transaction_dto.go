package transaction_dtos

import (
	wallet_dtos "ledger-system/dtos/wallet"
	"ledger-system/models"

	"github.com/shopspring/decimal"
)

type DetailTransactionDTO struct {
	ID         string                      `json:"id"`
	Type       models.TransactionType      `json:"type"`
	Amount     decimal.Decimal             `json:"amount"`
	Reference  string                      `json:"reference"`
	FromWallet *wallet_dtos.EmbedWalletDTO `json:"from_wallet"`
	ToWallet   *wallet_dtos.EmbedWalletDTO `json:"to_wallet"`
}
