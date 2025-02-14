package services

import (
	dtos "ledger-system/dtos/wallet"
	"ledger-system/models"
	"ledger-system/repositories"
	"log"

	"github.com/shopspring/decimal"
)

type WalletService struct {
	walletRepo *repositories.WalletRepository
	userRepo   *repositories.UserRepository
}

func NewWalletService(
	repo *repositories.WalletRepository,
	userRepo *repositories.UserRepository,
) *WalletService {
	return &WalletService{
		walletRepo: repo,
		userRepo:   userRepo,
	}
}

func (s *WalletService) GetWalletsService() ([]models.Wallet, error) {
	var wallets []models.Wallet

	wallets, err := s.walletRepo.GetWallets()
	if err != nil {
		log.Println("Error getting wallets:", err)
		return nil, err
	}

	return wallets, nil
}

func (s *WalletService) GetWalletService(id string) (models.Wallet, error) {
	var wallet models.Wallet

	wallet, err := s.walletRepo.GetWallet(id)
	if err != nil {
		log.Println("Error getting wallet:", err)
		return models.Wallet{}, err
	}

	return wallet, nil
}

func (s *WalletService) CreateWalletService(wallet dtos.CreateWalletDTO) (models.Wallet, error) {
	user, err := s.userRepo.FindUser(wallet.UserID)
	if err != nil {
		log.Println("Error getting user:", err)
		return models.Wallet{}, err
	}

	newWallet := models.Wallet{
		Name:    wallet.Name,
		Balance: decimal.NewFromFloat(wallet.Balance),
		UserID:  user.ID,
		User:    user,
	}

	createdWallet, err := s.walletRepo.CreateWallet(&newWallet)
	if err != nil {
		log.Println("Error creating wallet:", err)
		return models.Wallet{}, err
	}

	return *createdWallet, nil
}
