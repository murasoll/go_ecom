// services/auth_service.go

package services

import (
	"ecomerce/models"
	"ecomerce/repositories/auth_repo"
	"ecomerce/repositories/user_repo"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo user_repo.UserRepository
	authRepo auth_repo.AuthRepository
}

func NewAuthService(ur user_repo.UserRepository, ar auth_repo.AuthRepository) *AuthService {
	return &AuthService{
		userRepo: ur,
		authRepo: ar,
	}
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := s.generateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) Register(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.userRepo.CreateUser(*user)
}

func (s *AuthService) generateToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := []byte("your-secret-key") // Replace with a secure secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	dbToken := &models.Token{
		UserID:    user.ID,
		Token:     tokenString,
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}
	if err := s.authRepo.StoreToken(dbToken); err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *AuthService) ValidateToken(tokenString string) (*models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte("your-secret-key"), nil // Replace with your secret key
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := uint(claims["user_id"].(float64))
		role := models.Role(claims["role"].(string))

		user, err := s.userRepo.GetUserByID(userID)
		if err != nil {
			return nil, err
		}

		if user.Role != role {
			return nil, errors.New("invalid role")
		}

		// Check if token exists in the database
		if _, err := s.authRepo.GetToken(tokenString); err != nil {
			return nil, errors.New("token not found or expired")
		}

		return &user, nil
	}

	return nil, errors.New("invalid token")
}

func (s *AuthService) Logout(tokenString string) error {
	return s.authRepo.DeleteToken(tokenString)
}
