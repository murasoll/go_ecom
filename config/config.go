package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var DB *sql.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to ping the database:", err)
	}

	log.Println("Database connected successfully")

	autoMigrate()
	createAdminUser()
}

func autoMigrate() {
	createUsersTable := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(100) UNIQUE NOT NULL,
        email VARCHAR(100) UNIQUE NOT NULL,
        password VARCHAR(100) NOT NULL,
        role VARCHAR(20) NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );`

	createProductsTable := `
    CREATE TABLE IF NOT EXISTS products (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100),
        price NUMERIC(10, 2) NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );`

	createTokensTable := `
    CREATE TABLE IF NOT EXISTS tokens (
        id SERIAL PRIMARY KEY,
        user_id INTEGER NOT NULL REFERENCES users(id),
        token VARCHAR(255) UNIQUE NOT NULL,
        expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		log.Fatal("Error creating users table:", err)
	}

	_, err = DB.Exec(createProductsTable)
	if err != nil {
		log.Fatal("Error creating products table:", err)
	}

	_, err = DB.Exec(createTokensTable)
	if err != nil {
		log.Fatal("Error creating tokens table:", err)
	}

	log.Println("Tables ensured/created successfully")
}

func createAdminUser() {
	adminUsername := os.Getenv("ADMIN_USERNAME")
	adminEmail := os.Getenv("ADMIN_EMAIL")
	adminPassword := os.Getenv("ADMIN_PASSWORD")

	if adminUsername == "" || adminEmail == "" || adminPassword == "" {
		log.Fatal("Admin credentials not set in environment variables")
	}

	// Check if admin user already exists
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM users WHERE username = $1", adminUsername).Scan(&count)
	if err != nil {
		log.Fatal("Error checking for existing admin user:", err)
	}

	if count > 0 {
		log.Println("Admin user already exists")
		return
	}

	// Hash the admin password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Error hashing admin password:", err)
	}

	// Insert the admin user
	_, err = DB.Exec(`
        INSERT INTO users (username, email, password, role, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $5)
    `, adminUsername, adminEmail, string(hashedPassword), "admin", time.Now())

	if err != nil {
		log.Fatal("Error creating admin user:", err)
	}

	log.Println("Admin user created successfully")
}
