package wallet_dtos

import (
	dtoUser "ledger-system/dtos/user"

	"github.com/shopspring/decimal"
)

type WalletDetailDTO struct {
	ID        string               `json:"id"`
	Name      string               `json:"name"`
	Address   string               `json:"address"`
	User      dtoUser.UserEmbedDTO `json:"user"`
	Balance   decimal.Decimal      `json:"balance"`
	CreatedAt string               `json:"created_at"`
	UpdatedAt string               `json:"updated_at"`
}
