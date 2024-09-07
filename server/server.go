package server

import (
	"ecomerce/routes"

	"github.com/gorilla/mux"
)

func SetupServer() *mux.Router {
	r := mux.NewRouter()
	routes.RegisterRoutes(r) // Register API routes
	return r
}
