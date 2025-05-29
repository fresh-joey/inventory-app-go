package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"inventory-app/database"
	"inventory-app/models"
	"strconv"

	"github.com/gorilla/mux"
)

// List all orders, preload order items and products
func ListOrdersHandler(w http.ResponseWriter, r *http.Request) {
	var orders []models.Order
	database.DB.Preload("Items.Product").Find(&orders)

	tmpl, err := template.ParseFiles("assets/order_list.html") // your order list template
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, orders)
}

// Serve the order placement form page
func AddOrderPageHandler(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	database.DB.Find(&products)

	tmpl, err := template.ParseFiles("assets/order.html") // your order form template
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, products)
}

// Add a new order with multiple items (expects JSON payload)
func AddOrderHandler(w http.ResponseWriter, r *http.Request) {
	type OrderItemInput struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}

	type OrderInput struct {
		CustomerName string          `json:"customer_name"`
		Items        []OrderItemInput `json:"items"`
	}

	var input OrderInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	order := models.Order{CustomerName: input.CustomerName}
	if err := database.DB.Create(&order).Error; err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	for _, item := range input.Items {
		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
		database.DB.Create(&orderItem)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Order created"})
}

// Delete an order by ID
func DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	if err := database.DB.Delete(&models.Order{}, id).Error; err != nil {
		http.Error(w, "Failed to delete order", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/orders", http.StatusSeeOther)
}