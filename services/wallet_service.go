package services

import (
	"ledger-system/models"
	"ledger-system/repositories"
	"log"
)

func GetWalletsService() ([]models.Wallet, error) {
	var wallets []models.Wallet

	wallets, err := repositories.GetWallets()
	if err != nil {
		log.Println("Error getting wallets:", err)
		return nil, err
	}

	return wallets, nil
}

func GetWalletService(id string) (models.Wallet, error) {
	var wallet models.Wallet

	wallet, err := repositories.GetWallet(id)
	if err != nil {
		log.Println("Error getting wallet:", err)
		return models.Wallet{}, err
	}

	return wallet, nil
}
