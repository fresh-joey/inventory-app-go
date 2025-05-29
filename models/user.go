package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"` // stored as bcrypt hash
	CreatedAt time.Time
	UpdatedAt time.Time
}
