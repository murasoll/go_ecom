package server

import (
	v1 "ecomerce/api/v1"
	"ecomerce/middleware"
	"ecomerce/repositories/auth_repo"
	"ecomerce/repositories/cart_repo"
	"ecomerce/repositories/category_repo"
	"ecomerce/repositories/city_repo"
	"ecomerce/repositories/order_repo"
	"ecomerce/repositories/product_repo"
	"ecomerce/repositories/user_repo"
	"ecomerce/routes"
	"ecomerce/services"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupServer() http.Handler {
	r := mux.NewRouter()

	// Initialize repositories
	userRepo := user_repo.NewUserRepo()
	authRepo := auth_repo.NewAuthRepo()
	productRepo := product_repo.NewProductRepo()
	categoryRepo := category_repo.NewCategoryRepo()
	cartRepo := cart_repo.NewCartRepo()
	orderRepo := order_repo.NewOrderRepo()
	cityRepo := city_repo.NewCityRepo()

	// Initialize services
	authService := services.NewAuthService(userRepo, authRepo)
	productService := services.NewProductService(productRepo, categoryRepo)
	categoryService := services.NewCategoryService(categoryRepo)
	cartService := services.NewCartService(cartRepo, productRepo)
	orderService := services.NewOrderService(orderRepo, cartRepo, productRepo)
	shippingService := services.NewShippingService(cityRepo)

	// Initialize API handlers
	v1.InitAuthHandlers(authService)
	v1.InitProductHandlers(productService)
	v1.InitCategoryHandlers(categoryService)
	v1.InitCartHandlers(cartService)
	v1.InitOrderHandlers(orderService)
	v1.InitShippingHandlers(shippingService)

	// Initialize middleware
	authMiddleware := middleware.NewAuthMiddleware(authService)

	r.Use(middleware.CorsMiddleware())

	// Setup routes
	routes.SetupRoutes(r, authMiddleware)

	// Apply middlewares
	handler := middleware.CorsMiddleware()(r)

	return handler
}

func corsMiddleware(next http.Handler) http.Handler {
	return middleware.CorsMiddleware()(next)
}
