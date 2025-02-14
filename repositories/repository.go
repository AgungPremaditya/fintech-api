package repositories

import "gorm.io/gorm"

type Repositories struct {
	WalletRepository *WalletRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		WalletRepository: NewWalletRepository(db),
	}
}
