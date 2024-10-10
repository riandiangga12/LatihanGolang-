package entity

import "time"

type Transaction struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	WalletID  int       `gorm:"not null" json:"wallet_id"`
	WalletTo  int       `gorm:"null" json:"wallet_to"`
	Amount    float32   `gorm:"not null" json:"amount"`
	Type      string    `gorm:"size:100;not null" json:"type"` // topup, withdraw, transfer
	CreatedAt time.Time `json:"created_at"`
}
