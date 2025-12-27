package models

import "time"

type User struct {
	ID uint64 `gorm:"primaryKey;autoIncrement"`
	UUID string `gorm:"unique"`
	FullName string
	Email string `gorm:"unique"`
	EmailVerified bool
	ProfileImage string
	IsActive bool
	CreatedAt time.Time
}