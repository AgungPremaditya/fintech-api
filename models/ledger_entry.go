package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type LedgerEntry struct {
	ID        uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Type      string          `gorm:"type:varchar(255);not null"`
	Amount    decimal.Decimal `gorm:"type:decimal(36,18);not null"`
	Balance   decimal.Decimal `gorm:"type:decimal(36,18);not null"`
	CreatedAt time.Time       `gorm:"autoCreateTime"`

	// Wallet ID
	WalletID uuid.UUID `gorm:"type:uuid;not null"`
	Wallet   Wallet    `gorm:"foreignKey:WalletID"`
}

// Entry Types
type EntryType string

const (
	Debit  EntryType = "DEBIT"
	Credit EntryType = "CREDIT"
)
