package db

import (
	"database/sql"
	"ecomerce/config"
	"log"
)

func InitDB() *sql.DB {
	config.ConnectDatabase()
	log.Println("Database connected successfully")
	return config.DB
}
