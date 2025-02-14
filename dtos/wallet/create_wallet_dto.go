package dtos

type CreateWalletDTO struct {
	UserID  string  `json:"user_id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}
