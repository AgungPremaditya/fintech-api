package services

import "ledger-system/repositories"

type Service struct {
	WalletService *WalletService
}

func NewServices(repos *repositories.Repositories) *Service {
	return &Service{
		WalletService: NewWalletService(repos.WalletRepository),
	}
}
