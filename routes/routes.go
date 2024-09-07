package routes

import (
	"ecomerce/api/v1/auth"
	"ecomerce/api/v1/products"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	// Authentication routes
	r.HandleFunc("/api/v1/login", auth.Login).Methods("POST")
	r.HandleFunc("/api/v1/register", auth.Register).Methods("POST")

	// Product routes
	r.HandleFunc("/api/v1/products", products.GetAllProducts).Methods("GET")
	r.HandleFunc("/api/v1/products", products.CreateProduct).Methods("POST")
	r.HandleFunc("/api/v1/products/{id}", products.GetProductByID).Methods("GET")
	r.HandleFunc("/api/v1/products/{id}", products.UpdateProduct).Methods("PUT")
	r.HandleFunc("/api/v1/products/{id}", products.DeleteProduct).Methods("DELETE")
}
