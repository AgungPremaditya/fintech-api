package services

import (
	"ledger-system/models"
	"ledger-system/repositories"
	"log"
)

type WalletService struct {
	walletRepo *repositories.WalletRepository
}

func NewWalletService(repo *repositories.WalletRepository) *WalletService {
	return &WalletService{
		walletRepo: repo,
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
