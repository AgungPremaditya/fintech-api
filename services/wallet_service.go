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
