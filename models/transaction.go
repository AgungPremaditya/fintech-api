package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// Transaction Models
type Transaction struct {
	ID        uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Type      string          `gorm:"type:varchar(255);not null"`
	Amount    decimal.Decimal `gorm:"type:decimal(36,18);not null"`
	Reference string          `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time       `gorm:"autoCreateTime"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt  `gorm:"index"`

	// Relationships
	FromWalletID uuid.NullUUID `gorm:"type:uuid"`
	FromWallet   Wallet        `gorm:"foreignKey:FromWalletID"`

	ToWalletID uuid.NullUUID `gorm:"type:uuid"`
	ToWallet   Wallet        `gorm:"foreignKey:ToWalletID"`
}

// Transaction Type
type TransactionType string

const (
	Deposit  TransactionType = "DEPOSIT"
	Withdraw TransactionType = "WITHDRAW"
	Transfer TransactionType = "TRANSFER"
)

func ParseTransactionType(s string) (TransactionType, error) {
	switch s {
	case "DEPOSIT":
		return Deposit, nil
	case "WITHDRAW":
		return Withdraw, nil
	case "TRANSFER":
		return Transfer, nil
	default:
		return "", fmt.Errorf("invalid transaction type: %s", s)
	}
}
