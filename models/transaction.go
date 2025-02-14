package models

import (
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
	FromWalletID uuid.UUID `gorm:"type:uuid;not null"`
	FromWallet   Wallet    `gorm:"foreignKey:FromWalletID"`

	ToWalletID uuid.UUID `gorm:"type:uuid;not null"`
	ToWallet   Wallet    `gorm:"foreignKey:ToWalletID"`
}

// Transaction Type
type TransactionType string

const (
	Deposit  TransactionType = "DEPOSIT"
	Withdraw TransactionType = "WITHDRAW"
	Transfer TransactionType = "TRANSFER"
)
