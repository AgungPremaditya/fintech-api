package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Wallet Models
type Wallet struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string         `gorm:"type:varchar(255);not null"`
	Address   string         `gorm:"type:varchar(255);not null;uniqueIndex"`
	UserID    uuid.UUID      `gorm:"type:uuid;not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	User      User           `gorm:"foreignKey:UserID"`

	// Transactions relationships
	SentTransactions     []Transaction `json:"sent_transactions" gorm:"foreignKey:FromWalletID"`
	ReceivedTransactions []Transaction `json:"received_transactions" gorm:"foreignKey:ToWalletID"`
}
