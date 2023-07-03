package models

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	WalletID  uuid.UUID `json:"wallet_id"`
	Amount    uint64    `json:"amount"`
	CreatedAt time.Time `json:"created_at" gorm:"index:transaction_idx"`
	UpdatedAt time.Time `json:"updated_at"`
}
