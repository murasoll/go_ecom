package middleware

import (
	"ecomerce/config"
	"log"
	"net/http"
)

func CorsMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			allowedOrigins := config.GetAllowedOrigins()

			origin := r.Header.Get("Origin")

			log.Printf("Received request from origin: %s", origin)

			if origin == "" {
				log.Printf("No Origin header present. Request from: %s", r.RemoteAddr)
				// For requests without Origin, we'll still serve the request
				// but won't add any CORS headers
				next.ServeHTTP(w, r)
				return
			}

			allowed := false
			for _, allowedOrigin := range allowedOrigins {
				if origin == allowedOrigin {
					allowed = true
					break
				}
			}

			if !allowed {
				log.Printf("Origin not allowed: %s", origin)
				http.Error(w, "Origin not allowed", http.StatusForbidden)
				return
			}

			// If we reach here, the origin is allowed
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Max-Age", "300")

			// Handle preflight requests
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
