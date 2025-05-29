package handlers

import (
	"html/template"
	"net/http"
	"inventory-app/database"
	"inventory-app/models"
	"strconv"

	"github.com/gorilla/mux"
)

func ListCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	var categories []models.Category
	database.DB.Find(&categories)

	tmpl, _ := template.ParseFiles("assets/categories.html")
	tmpl.Execute(w, categories)
}

func AddCategoryPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("assets/category_add.html")
	tmpl.Execute(w, nil)
}

func AddCategoryHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")

	category := models.Category{Name: name}
	database.DB.Create(&category)
	http.Redirect(w, r, "/admin/categories", http.StatusSeeOther)
}

func DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)

	database.DB.Delete(&models.Category{}, id)
	http.Redirect(w, r, "/admin/categories", http.StatusSeeOther)
}
