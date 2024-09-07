package routes

import (
	"ecomerce/api/v1/auth"
	"ecomerce/api/v1/products"
	"ecomerce/middleware"
	"ecomerce/models"
	"ecomerce/services"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router, authService *services.AuthService) {
	// Create auth middleware
	authMiddleware := middleware.NewAuthMiddleware(authService)

	// auth routes
	r.HandleFunc("/api/v1/login", auth.Login).Methods("POST")
	r.HandleFunc("/api/v1/register", auth.Register).Methods("POST")
	r.HandleFunc("/api/v1/logout", authMiddleware.Authenticate(auth.Logout)).Methods("POST")

	// Product routes
	r.HandleFunc("/api/v1/products", products.GetAllProducts).Methods("GET")
	r.HandleFunc("/api/v1/products", authMiddleware.Authenticate(authMiddleware.Authorize(models.RoleAdmin, models.RoleManager)(products.CreateProduct))).Methods("POST")
	r.HandleFunc("/api/v1/products/{id}", authMiddleware.Authenticate(products.GetProductByID)).Methods("GET")
	r.HandleFunc("/api/v1/products/{id}", authMiddleware.Authenticate(authMiddleware.Authorize(models.RoleAdmin, models.RoleManager)(products.UpdateProduct))).Methods("PUT")
	r.HandleFunc("/api/v1/products/{id}", authMiddleware.Authenticate(authMiddleware.Authorize(models.RoleAdmin)(products.DeleteProduct))).Methods("DELETE")
}
