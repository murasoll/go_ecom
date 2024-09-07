package middleware

import (
	"context"
	"ecomerce/helpers"
	"ecomerce/models"
	"ecomerce/services"
	"net/http"
)

type AuthMiddleware struct {
	authService *services.AuthService
}

func NewAuthMiddleware(authService *services.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

func (m *AuthMiddleware) Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			helpers.JsonResponse(w, map[string]string{"error": "No token provided"}, http.StatusUnauthorized)
			return
		}

		// Remove "Bearer " prefix if present
		tokenString = helpers.TrimBearerPrefix(tokenString)

		user, err := m.authService.ValidateToken(tokenString)
		if err != nil {
			helpers.JsonResponse(w, map[string]string{"error": "Invalid or expired token"}, http.StatusUnauthorized)
			return
		}

		// Add the user to the request context
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (m *AuthMiddleware) Authorize(roles ...models.Role) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			user, ok := r.Context().Value("user").(*models.User)
			if !ok {
				helpers.JsonResponse(w, map[string]string{"error": "Unauthorized"}, http.StatusUnauthorized)
				return
			}

			for _, role := range roles {
				if user.Role == role {
					next.ServeHTTP(w, r)
					return
				}
			}

			helpers.JsonResponse(w, map[string]string{"error": "Forbidden"}, http.StatusForbidden)
		}
	}
}
