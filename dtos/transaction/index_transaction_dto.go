package transaction_dtos

import (
	general_dtos "ledger-system/dtos/general"
	wallet_dtos "ledger-system/dtos/wallet"
)

type TransactionIndexDTO struct {
	ID        string                      `json:"id"`
	Type      string                      `json:"type"`
	Wallet    *wallet_dtos.EmbedWalletDTO `json:"from_wallet"`
	Amount    string                      `json:"amount"`
	Reference string                      `json:"reference"`
	CreatedAt string                      `json:"created_at"`
}

type TransactionPaginatedDTO struct {
	Transactions []TransactionIndexDTO       `json:"transactions"`
	Meta         general_dtos.PaginationMeta `json:"meta"`
}
