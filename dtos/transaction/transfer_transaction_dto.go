package transaction_dtos

type TransferTransactionDTO struct {
	Amount       float64 `json:"amount"`
	Reference    string  `json:"reference"`
	FromWalletID string  `json:"from_wallet_id"`
	ToWalletID   string  `json:"to_wallet_id"`
}
