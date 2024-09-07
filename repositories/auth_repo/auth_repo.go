package auth_repo

import (
	"ecomerce/config"
	"ecomerce/models"
	"ecomerce/repositories"
	"log"
	"time"
)

type AuthRepoImpl struct{}

func NewAuthRepo() repositories.AuthRepository {
	return &AuthRepoImpl{}
}

func (r *AuthRepoImpl) StoreToken(token *models.Token) error {
	query := `INSERT INTO tokens (user_id, token, expires_at) VALUES ($1, $2, $3) RETURNING id`
	err := config.DB.QueryRow(query, token.UserID, token.Token, token.ExpiresAt).Scan(&token.ID)
	if err != nil {
		log.Println("Error storing token:", err)
		return err
	}

	log.Println("Token stored with ID:", token.ID)
	return nil
}

func (r *AuthRepoImpl) GetToken(tokenString string) (*models.Token, error) {
	var token models.Token
	query := `SELECT id, user_id, token, expires_at FROM tokens WHERE token = $1`

	err := config.DB.QueryRow(query, tokenString).Scan(&token.ID, &token.UserID, &token.Token, &token.ExpiresAt)
	if err != nil {
		log.Println("Error fetching token:", err)
		return nil, err
	}

	return &token, nil
}

func (r *AuthRepoImpl) DeleteToken(tokenString string) error {
	query := `DELETE FROM tokens WHERE token = $1`
	_, err := config.DB.Exec(query, tokenString)
	if err != nil {
		log.Println("Error deleting token:", err)
		return err
	}

	log.Println("Token deleted:", tokenString)
	return nil
}

func (r *AuthRepoImpl) CleanExpiredTokens() error {
	query := `DELETE FROM tokens WHERE expires_at < $1`
	_, err := config.DB.Exec(query, time.Now())
	if err != nil {
		log.Println("Error cleaning expired tokens:", err)
		return err
	}

	log.Println("Expired tokens cleaned")
	return nil
}
