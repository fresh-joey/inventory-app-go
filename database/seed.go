package database

import (
	"log"
	"inventory-app/models"
)

func SeedData() {
	defaultCategories := []models.Category{
		{Name: "Electronics"},
		{Name: "Books"},
		{Name: "Clothing"},
		{Name: "Furniture"},
	}

	for _, category := range defaultCategories {
		var existing models.Category
		result := DB.First(&existing, "name = ?", category.Name)
		if result.RowsAffected == 0 {
			if err := DB.Create(&category).Error; err != nil {
				log.Printf("Failed to seed category %s: %v", category.Name, err)
			}
		}
	}
}
