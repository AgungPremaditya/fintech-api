package transaction_dtos

type CreateTransactionDTO struct {
	Type      string  `json:"type"`
	Amount    float64 `json:"amount"`
	Reference string  `json:"reference"`
	WalletID  string  `json:"wallet_id"`
}
