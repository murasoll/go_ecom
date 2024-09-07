package main

import (
	"ecomerce/config"
	"ecomerce/server"
	"log"
	"net/http"
)

func main() {
	config.ConnectDatabase()
	s := server.SetupServer()
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", s))
}
