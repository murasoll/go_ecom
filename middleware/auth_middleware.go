// middleware/auth_middleware.go

package middleware

import (
	"ecomerce/helpers"
	"ecomerce/models"
	"ecomerce/services"
	"net/http"
)

type AuthMiddleware struct {
	authService *services.AuthService
}

func NewAuthMiddleware(as *services.AuthService) *AuthMiddleware {
	return &AuthMiddleware{authService: as}
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
		ctx := services.ContextWithUser(r.Context(), user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (m *AuthMiddleware) Authorize(roles ...models.Role) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			user, err := services.UserFromContext(r.Context())
			if err != nil {
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
