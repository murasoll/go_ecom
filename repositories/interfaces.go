package repositories

import (
	"ecomerce/models"
)

type UserRepository interface {
	GetUserByEmail(email string) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	GetUserByID(id uint) (models.User, error)
	CreateUser(user models.User) error
	UpdateUser(user models.User) error
	DeleteUser(id uint) error
}

type AuthRepository interface {
	StoreToken(token *models.Token) error
	GetToken(tokenString string) (*models.Token, error)
	DeleteToken(tokenString string) error
	CleanExpiredTokens() error
}
