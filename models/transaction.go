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
	WalletID  uuid.UUID       `gorm:"type:uuid;not null"`
	CreatedAt time.Time       `gorm:"autoCreateTime"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt  `gorm:"index"`
}

// Transaction Type
type TransactionType string

const (
	Deposit  TransactionType = "DEPOSIT"
	Withdraw TransactionType = "WITHDRAW"
	Transfer TransactionType = "TRANSFER"
)
