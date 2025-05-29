package handlers

import (
	"html/template"
	"net/http"

	"inventory-app/database"
	"inventory-app/models"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	var totalProducts int64
	var totalCategories int64
	var recentOrders []models.Order

	database.DB.Model(&models.Product{}).Count(&totalProducts)
	database.DB.Model(&models.Category{}).Count(&totalCategories)
	database.DB.Preload("Product").Order("created_at desc").Limit(5).Find(&recentOrders)

	tmpl, _ := template.ParseFiles("assets/dashboard.html")
	tmpl.Execute(w, map[string]interface{}{
		"TotalProducts":   totalProducts,
		"TotalCategories": totalCategories,
		"RecentOrders":    recentOrders,
	})
}
