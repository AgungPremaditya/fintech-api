package wallet_dtos

import (
	user_dtos "ledger-system/dtos/user"
	"ledger-system/models"

	"github.com/google/uuid"
)

type Mapper struct {
	userMapper *user_dtos.Mapper
}

func NewWalletMapper() *Mapper {
	return &Mapper{
		userMapper: user_dtos.NewUserMapper(),
	}
}

func (m *Mapper) ToWalletResponse(wallet *models.Wallet) WalletIndexDTO {
	return WalletIndexDTO{
		ID:        wallet.ID.String(),
		Name:      wallet.Name,
		Address:   wallet.Address,
		CreatedAt: wallet.CreatedAt.String(),
		UpdatedAt: wallet.UpdatedAt.String(),
	}
}

func (m *Mapper) ToWalletResponseList(wallets []models.Wallet) []WalletIndexDTO {
	result := make([]WalletIndexDTO, len(wallets))
	for i, wallet := range wallets {
		result[i] = m.ToWalletResponse(&wallet)
	}
	return result
}

func (m *Mapper) ToWalletModel(dto *CreateWalletDTO, user *models.User) (*models.Wallet, error) {
	return &models.Wallet{
		Name:    dto.Name,
		UserID:  user.ID,
		Address: dto.Address,
	}, nil
}

func (m *Mapper) ToWalletDetailResponse(wallet models.Wallet) WalletDetailDTO {
	// Map user data
	userEmbed := m.userMapper.ToUserEmbedDTO(wallet.User)

	return WalletDetailDTO{
		ID:        wallet.ID.String(),
		Name:      wallet.Name,
		Address:   wallet.Address,
		User:      userEmbed,
		Balance:   wallet.Balance,
		CreatedAt: wallet.CreatedAt.String(),
		UpdatedAt: wallet.UpdatedAt.String(),
	}
}

func (m *Mapper) ToWalletEmbedDTO(wallet *models.Wallet) *EmbedWalletDTO {
	if wallet.ID == uuid.Nil || wallet == nil {
		return nil
	}

	return &EmbedWalletDTO{
		ID:      wallet.ID.String(),
		Name:    wallet.Name,
		Address: wallet.Address,
	}
}
