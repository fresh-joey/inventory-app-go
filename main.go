package main

import (
	"fmt"
	"log"
	"net/http"

	"inventory-app/database"
	"inventory-app/handlers"
	"inventory-app/middleware"
	"inventory-app/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/gorilla/mux"
)

func main() {
	database.Connect()
	database.Migrate()
	database.SeedData()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("defaultpassword"), bcrypt.DefaultCost)

	database.DB.FirstOrCreate(&models.User{}, models.User{
		Username: "Admin",
		Password: string(hashedPassword),
	})


	r := mux.NewRouter()

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	r.HandleFunc("/", handlers.DashboardHandler)
	r.HandleFunc("/login", handlers.LoginPageHandler).Methods("GET")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("GET")

	// Protected routes
	admin := r.PathPrefix("/admin").Subrouter()
	admin.Use(middleware.AuthMiddleware)
	admin.HandleFunc("/dashboard", handlers.DashboardHandler)
	admin.HandleFunc("/products", handlers.ListProductsHandler).Methods("GET")
	admin.HandleFunc("/products", handlers.AddProductHandler).Methods("POST")

	//Orders routes
	admin.HandleFunc("/orders/new", handlers.AddOrderPageHandler).Methods("GET")
	admin.HandleFunc("/orders", handlers.ListOrdersHandler).Methods("GET")
	admin.HandleFunc("/orders", handlers.AddOrderHandler).Methods("POST")
	admin.HandleFunc("/orders/{id}", handlers.DeleteOrderHandler).Methods("DELETE")

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
