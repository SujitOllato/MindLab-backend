package models

import "time"

type AuthProvider struct {
	ID uint64 `gorm:"primaryKey;autoIncrement"`
	UserID uint64
	Provider string
	ProviderUserID string
	PasswordHash string
	CreatedAt time.Time
}