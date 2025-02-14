package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// Wallet Models
type Wallet struct {
	ID           uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name         string          `gorm:"type:varchar(255);not null"`
	Balance      decimal.Decimal `gorm:"type:decimal(36,18);not null"`
	UserID       uuid.UUID       `gorm:"type:uuid;not null"`
	CreatedAt    time.Time       `gorm:"autoCreateTime"`
	UpdatedAt    time.Time       `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt  `gorm:"index"`
	Transactions []Transaction   `gorm:"foreignKey:WalletID"`
}
