package auth

import (
	"ecomerce/helpers"
	"ecomerce/models"
	"ecomerce/services"
	"encoding/json"
	"net/http"
)

var authService *services.AuthService

// InitAuthHandlers initializes the auth handlers with the necessary dependencies
func InitAuthHandlers(as *services.AuthService) {
	authService = as
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid request payload"}, http.StatusBadRequest)
		return
	}

	token, err := authService.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid credentials"}, http.StatusUnauthorized)
		return
	}

	helpers.JsonResponse(w, map[string]string{
		"message": "Login successful",
		"token":   token,
	}, http.StatusOK)
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid request payload"}, http.StatusBadRequest)
		return
	}

	// Set a default role for new users (e.g., customer)
	user.Role = models.RoleCustomer

	err = authService.Register(&user)
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Error registering user"}, http.StatusInternalServerError)
		return
	}

	helpers.JsonResponse(w, map[string]string{
		"message": "User registered successfully",
	}, http.StatusCreated)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		helpers.JsonResponse(w, map[string]string{"error": "No token provided"}, http.StatusBadRequest)
		return
	}

	// Remove "Bearer " prefix if present
	tokenString = helpers.TrimBearerPrefix(tokenString)

	err := authService.Logout(tokenString)
	if err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Error logging out"}, http.StatusInternalServerError)
		return
	}

	helpers.JsonResponse(w, map[string]string{
		"message": "Logged out successfully",
	}, http.StatusOK)
}
