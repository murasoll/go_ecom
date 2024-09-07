package auth

import (
	"ecomerce/helpers"
	"ecomerce/models"
	"ecomerce/services"
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, `{"error": "Invalid request payload"}`, http.StatusBadRequest)
		return
	}

	// Call authentication logic from the services layer
	isAuthenticated := services.AuthenticateUser(user.Email, user.Password)
	if !isAuthenticated {
		http.Error(w, `{"error": "Invalid credentials"}`, http.StatusUnauthorized)
		return
	}

	// Respond with success in JSON format
	helpers.JsonResponse(w, map[string]string{
		"message": "Login successful",
	}, http.StatusOK)
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, `{"error": "Invalid request payload"}`, http.StatusBadRequest)
		return
	}

	// Call registration logic from the services layer
	err = services.RegisterUser(user)
	if err != nil {
		http.Error(w, `{"error": "Error registering user"}`, http.StatusInternalServerError)
		return
	}

	// Respond with success in JSON format
	helpers.JsonResponse(w, map[string]string{
		"message": "User registered successfully",
	}, http.StatusCreated)
}
