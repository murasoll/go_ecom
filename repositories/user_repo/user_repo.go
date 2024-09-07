package user_repo

import (
	"ecomerce/config"
	"ecomerce/models"
	"ecomerce/repositories"
	"log"
)

type UserRepoImpl struct{}

func NewUserRepo() repositories.UserRepository {
	return &UserRepoImpl{}
}

func (r *UserRepoImpl) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	query := `SELECT id, username, email, password, role FROM users WHERE email = $1`

	err := config.DB.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)
	if err != nil {
		log.Println("Error fetching user by email:", err)
		return user, err
	}

	return user, nil
}

func (r *UserRepoImpl) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	query := `SELECT id, username, email, password, role FROM users WHERE username = $1`

	err := config.DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)
	if err != nil {
		log.Println("Error fetching user by username:", err)
		return user, err
	}

	return user, nil
}

func (r *UserRepoImpl) GetUserByID(id uint) (models.User, error) {
	var user models.User
	query := `SELECT id, username, email, password, role FROM users WHERE id = $1`

	err := config.DB.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)
	if err != nil {
		log.Println("Error fetching user by ID:", err)
		return user, err
	}

	return user, nil
}

func (r *UserRepoImpl) CreateUser(user models.User) error {
	query := `INSERT INTO users (username, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id`
	err := config.DB.QueryRow(query, user.Username, user.Email, user.Password, user.Role).Scan(&user.ID)
	if err != nil {
		log.Println("Error creating user:", err)
		return err
	}

	log.Println("User created with ID:", user.ID)
	return nil
}

func (r *UserRepoImpl) UpdateUser(user models.User) error {
	query := `UPDATE users SET username = $1, email = $2, password = $3, role = $4 WHERE id = $5`
	_, err := config.DB.Exec(query, user.Username, user.Email, user.Password, user.Role, user.ID)
	if err != nil {
		log.Println("Error updating user:", err)
		return err
	}

	log.Println("User updated with ID:", user.ID)
	return nil
}

func (r *UserRepoImpl) DeleteUser(id uint) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := config.DB.Exec(query, id)
	if err != nil {
		log.Println("Error deleting user:", err)
		return err
	}

	log.Println("User deleted with ID:", id)
	return nil
}
