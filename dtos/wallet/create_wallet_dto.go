package wallet_dtos

type CreateWalletDTO struct {
	UserID  string `json:"user_id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
