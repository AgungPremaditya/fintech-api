package repositories

import "gorm.io/gorm"

type Repositories struct {
	WalletRepository      *WalletRepository
	UserRepository        *UserRepository
	TransactionRepository *TransactionRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		WalletRepository:      NewWalletRepository(db),
		UserRepository:        NewUserRepository(db),
		TransactionRepository: NewTransactionRepository(db),
	}
}
