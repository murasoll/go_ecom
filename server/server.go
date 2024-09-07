package server

import (
	"ecomerce/api/v1/auth"
	"ecomerce/repositories/auth_repo"
	"ecomerce/repositories/user_repo"
	"ecomerce/routes"
	"ecomerce/services"

	"github.com/gorilla/mux"
)

func SetupServer() *mux.Router {
	r := mux.NewRouter()

	userRepo := user_repo.NewUserRepo()
	authRepo := auth_repo.NewAuthRepo()
	authService := services.NewAuthService(userRepo, authRepo)
	auth.InitAuthHandlers(authService)

	routes.SetupRoutes(r, authService) // Register API routes
	return r
}
