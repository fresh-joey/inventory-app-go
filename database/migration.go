// db/migrate.go
package database

import (
	"log"
	"inventory-app/models"
)

func Migrate() {
	err := DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}, &models.OrderItem{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	log.Println("Migration successful")
}
