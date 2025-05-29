package models

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"not null"`
	Description string
	CategoryID  uint
	Category    Category
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
