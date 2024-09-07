package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

func ConnectDatabase() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the database connection details from the environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	// Create the PostgreSQL connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode)

	// Open the database connection
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Check if the connection is successful
	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to ping the database:", err)
	}

	log.Println("Database connected successfully")

	autoMigrate() // Ensure tables are created automatically
}

func autoMigrate() {
	createUsersTable := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100),
        email VARCHAR(100) UNIQUE NOT NULL,
        password VARCHAR(100) NOT NULL
    );`

	createProductsTable := `
    CREATE TABLE IF NOT EXISTS products (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100),
        price NUMERIC(10, 2) NOT NULL
    );`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		log.Fatal("Error creating users table:", err)
	}

	_, err = DB.Exec(createProductsTable)
	if err != nil {
		log.Fatal("Error creating products table:", err)
	}

	log.Println("Tables ensured/created successfully")
}
