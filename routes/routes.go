// routes/routes.go

package routes

import (
	v1 "ecomerce/api/v1"
	"ecomerce/middleware"
	"ecomerce/models"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Method       string
	Path         string
	Handler      http.HandlerFunc
	AuthRequired bool
	Roles        []models.Role
}

func SetupRoutes(r *mux.Router, authMiddleware *middleware.AuthMiddleware) {
	routes := []Route{
		// Public routes
		{Method: "POST", Path: "/login", Handler: v1.Login, AuthRequired: false},
		{Method: "POST", Path: "/register", Handler: v1.Register, AuthRequired: false},

		{Method: "GET", Path: "/products", Handler: v1.GetAllProducts, AuthRequired: false},
		{Method: "GET", Path: "/products/{id}", Handler: v1.GetProduct, AuthRequired: false},
		{Method: "POST", Path: "/products", Handler: v1.CreateProduct, AuthRequired: true, Roles: []models.Role{models.RoleAdmin, models.RoleManager}},
		{Method: "PUT", Path: "/products/{id}", Handler: v1.UpdateProduct, AuthRequired: true, Roles: []models.Role{models.RoleAdmin, models.RoleManager}},
		{Method: "DELETE", Path: "/products/{id}", Handler: v1.DeleteProduct, AuthRequired: true, Roles: []models.Role{models.RoleAdmin}},

		{Method: "GET", Path: "/categories", Handler: v1.GetAllCategories, AuthRequired: false},
		{Method: "POST", Path: "/categories", Handler: v1.CreateCategory, AuthRequired: true, Roles: []models.Role{models.RoleAdmin}},
		{Method: "PUT", Path: "/categories/{id}", Handler: v1.UpdateCategory, AuthRequired: true, Roles: []models.Role{models.RoleAdmin}},
		{Method: "DELETE", Path: "/categories/{id}", Handler: v1.DeleteCategory, AuthRequired: true, Roles: []models.Role{models.RoleAdmin}},

		{Method: "GET", Path: "/cart", Handler: v1.GetCart, AuthRequired: true},
		{Method: "POST", Path: "/cart", Handler: v1.AddToCart, AuthRequired: true},
		{Method: "PUT", Path: "/cart/{product_id}", Handler: v1.UpdateCartItem, AuthRequired: true},
		{Method: "DELETE", Path: "/cart/{product_id}", Handler: v1.RemoveFromCart, AuthRequired: true},

		{Method: "POST", Path: "/orders", Handler: v1.CreateOrder, AuthRequired: true},
		{Method: "GET", Path: "/orders", Handler: v1.GetOrders, AuthRequired: true},
		{Method: "GET", Path: "/orders/{id}", Handler: v1.GetOrder, AuthRequired: true},
		{Method: "PUT", Path: "/orders/{id}/status", Handler: v1.UpdateOrderStatus, AuthRequired: true, Roles: []models.Role{models.RoleAdmin, models.RoleManager}},

		{Method: "GET", Path: "/shipping/cost", Handler: v1.GetShippingCost, AuthRequired: true},
		{Method: "GET", Path: "/cities", Handler: v1.GetAllCities, AuthRequired: true},
		{Method: "POST", Path: "/cities", Handler: v1.CreateCity, AuthRequired: true, Roles: []models.Role{models.RoleAdmin}},
		{Method: "PUT", Path: "/cities/{id}", Handler: v1.UpdateCity, AuthRequired: true, Roles: []models.Role{models.RoleAdmin}},
		{Method: "DELETE", Path: "/cities/{id}", Handler: v1.DeleteCity, AuthRequired: true, Roles: []models.Role{models.RoleAdmin}},

		{Method: "POST", Path: "/logout", Handler: v1.Logout, AuthRequired: true},
	}

	api := r.PathPrefix("/api/v1").Subrouter()

	for _, route := range routes {
		handler := route.Handler
		if route.AuthRequired {
			handler = authMiddleware.Authenticate(handler)
			if len(route.Roles) > 0 {
				handler = authMiddleware.Authenticate(authMiddleware.Authorize(route.Roles...)(handler))
			}
		}
		api.HandleFunc(route.Path, handler).Methods(route.Method)
	}
}
