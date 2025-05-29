package models

import "gorm.io/gorm"

type OrderItem struct {
    gorm.Model
    OrderID   uint
    ProductID uint
    Quantity  int
    // Product   Product
}