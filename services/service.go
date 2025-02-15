package services

import "ledger-system/repositories"

type Service struct {
	WalletService      *WalletService
	TransactionService *TransactionService
}

func NewServices(repos *repositories.Repositories) *Service {
	return &Service{
		WalletService:      NewWalletService(repos.WalletRepository, repos.UserRepository),
		TransactionService: NewTransactionService(repos.TransactionRepository, repos.WalletRepository, repos.LedgerEntryRepository),
	}
}
