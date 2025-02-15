package wallet_dtos

import (
	user_dtos "ledger-system/dtos/user"
	"ledger-system/models"
	"ledger-system/repositories"
	"log"

	"github.com/shopspring/decimal"
)

type Mapper struct {
	userRepo   *repositories.UserRepository
	userMapper *user_dtos.Mapper
}

func NewWalletMapper(
	userRepo *repositories.UserRepository,
) *Mapper {
	return &Mapper{
		userRepo:   userRepo,
		userMapper: user_dtos.NewUserMapper(),
	}
}

func (m *Mapper) ToWalletResponse(wallet *models.Wallet) WalletIndexDTO {
	return WalletIndexDTO{
		ID:        wallet.ID.String(),
		Name:      wallet.Name,
		Address:   wallet.Address,
		CreatedAt: wallet.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: wallet.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func (m *Mapper) ToWalletResponseList(wallets []models.Wallet) []WalletIndexDTO {
	result := make([]WalletIndexDTO, len(wallets))
	for i, wallet := range wallets {
		result[i] = m.ToWalletResponse(&wallet)
	}
	return result
}

func (m *Mapper) ToWalletModel(dto CreateWalletDTO) (models.Wallet, error) {
	user, err := m.userRepo.FindUser(dto.UserID)
	if err != nil {
		log.Println("Error getting user:", err)
		return models.Wallet{}, err
	}

	return models.Wallet{
		Name:    dto.Name,
		UserID:  user.ID,
		Address: dto.Address,
	}, nil
}

func (m *Mapper) ToWalletDetailResponse(wallet models.Wallet, balance *decimal.Decimal) WalletDetailDTO {
	// Map user data
	userEmbed := m.userMapper.ToUserEmbedDTO(wallet.User)

	return WalletDetailDTO{
		ID:        wallet.ID.String(),
		Name:      wallet.Name,
		Address:   wallet.Address,
		User:      userEmbed,
		Balance:   *balance,
		CreatedAt: wallet.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: wallet.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
