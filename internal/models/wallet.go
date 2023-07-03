package models

import (
	"github.com/google/uuid"
	"time"
)

type Wallet struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID    uuid.UUID `json:"user_id"`
	Type      string    `json:"type" gorm:"default:non-identified"`
	Balance   uint64    `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Transactions []Transaction `json:"-" gorm:"foreignKey:WalletID"`
}
