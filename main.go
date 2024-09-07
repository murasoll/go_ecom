package main

import (
	"ecomerce/config"
	"ecomerce/server"
	"log"
	"net/http"
)

func main() {
	config.ConnectDatabase()  // Ensure the database is connected and tables are created
	s := server.SetupServer() // Initialize the server
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", s)) // Start the HTTP server
}
