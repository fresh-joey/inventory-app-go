package handlers

import (
	"html/template"
	"net/http"
	"inventory-app/models"
	"inventory-app/database"
	"strconv"

	"github.com/gorilla/mux"
	// "gorm.io/gorm"
)

func ListProductsHandler(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	database.DB.Preload("Category").Find(&products)

	tmpl, _ := template.ParseFiles("assets/product_list.html")
	tmpl.Execute(w, products)
}

func AddProductPageHandler(w http.ResponseWriter, r *http.Request) {
	var categories []models.Category
	database.DB.Find(&categories)

	tmpl, _ := template.ParseFiles("assets/product_add.html")
	tmpl.Execute(w, categories)
}

func AddProductHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	description := r.FormValue("description")
	categoryID, _ := strconv.ParseUint(r.FormValue("category_id"), 10, 64)

	product := models.Product{
		Name:        name,
		Description: description,
		CategoryID:  uint(categoryID),
	}

	database.DB.Create(&product)
	http.Redirect(w, r, "/admin/products", http.StatusSeeOther)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)

	database.DB.Delete(&models.Product{}, id)
	http.Redirect(w, r, "/admin/products", http.StatusSeeOther)
}
