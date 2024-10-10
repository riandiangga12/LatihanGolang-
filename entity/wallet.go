package entity

import "time"

type Wallet struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	UserID    int       `gorm:"not null" json:"user_id"`
	Balance   float32   `gorm:"not null" json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
