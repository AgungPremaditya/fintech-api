package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Transaction Models
type Transaction struct {
	ID        uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Type      string          `gorm:"type:varchar(255);not null"`
	Amount    decimal.Decimal `gorm:"type:decimal(36,18);not null"`
	Reference string          `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time       `gorm:"autoCreateTime"`

	// Wallet relationships
	WalletID uuid.UUID `gorm:"type:uuid;not null"`
	Wallet   Wallet    `gorm:"foreignKey:WalletID"`
}

// Transaction Type
type TransactionType string

const (
	Deposit  TransactionType = "DEPOSIT"
	Withdraw TransactionType = "WITHDRAW"
	Transfer TransactionType = "TRANSFER"
)
