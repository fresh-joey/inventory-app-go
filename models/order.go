package models

import "time"

type Order struct {
	ID        uint      `gorm:"primaryKey"`
	ProductID uint
	Product   Product
    CustomerName string      `json:"customer_name"`
	Quantity  int       `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Items []OrderItem
}
