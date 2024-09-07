package v1

import (
	"ecomerce/helpers"
	"ecomerce/models"
	"ecomerce/services"
	"encoding/json"
	"net/http"
)

var authService *services.AuthService

func InitAuthHandlers(as *services.AuthService) {
	authService = as
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
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
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Invalid request payload"}, http.StatusBadRequest)
		return
	}

	user.Role = models.RoleCustomer
	if err := authService.Register(&user); err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Error registering user"}, http.StatusInternalServerError)
		return
	}

	helpers.JsonResponse(w, map[string]string{
		"message": "User registered successfully",
	}, http.StatusCreated)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		helpers.JsonResponse(w, map[string]string{"error": "No token provided"}, http.StatusBadRequest)
		return
	}

	tokenString = helpers.TrimBearerPrefix(tokenString)
	if err := authService.Logout(tokenString); err != nil {
		helpers.JsonResponse(w, map[string]string{"error": "Error logging out"}, http.StatusInternalServerError)
		return
	}

	helpers.JsonResponse(w, map[string]string{
		"message": "Logged out successfully",
	}, http.StatusOK)
}
