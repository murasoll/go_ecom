package user_repo

import (
	"ecomerce/config"
	"ecomerce/models"
	"log"
)

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	query := `SELECT id, name, email, password FROM users WHERE email = $1`

	err := config.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		log.Println("Error fetching user by email:", err)
		return user, err
	}

	return user, nil
}

func CreateUser(user models.User) error {
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	err := config.DB.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		log.Println("Error creating user:", err)
		return err
	}

	log.Println("User created with ID:", user.ID)
	return nil
}
