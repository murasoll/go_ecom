package services

import (
	"ecomerce/models"
	"ecomerce/repositories/user_repo"
	"errors"
)

func AuthenticateUser(email, password string) bool {
	// Fetch user from the database
	user, err := user_repo.GetUserByEmail(email)
	if err != nil {
		return false
	}

	// Compare the provided password with the stored hashed password
	// (Password comparison logic not implemented)
	if user.Password == password {
		return true
	}

	return false
}

func RegisterUser(user models.User) error {
	// Perform necessary validation on the user data

	// Save the user to the database
	err := user_repo.CreateUser(user)
	if err != nil {
		return errors.New("failed to register user")
	}

	return nil
}
