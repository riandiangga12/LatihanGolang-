package entity

import "time"

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name" binding:"required,min=3"`
	Email     string    `gorm:"size:100;unique;not null" json:"email" binding:"required,email"`
	Password  string    `gorm:"size:255" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
