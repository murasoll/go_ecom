package config

import "os"

var (
	// DevelopmentOrigins is the list of allowed origins for the development environment
	DevelopmentOrigins = []string{
		"http://localhost:5173",
	}

	// ProductionOrigins is the list of allowed origins for the production environment
	ProductionOrigins = []string{
		"https://yourdomain.com",
		"https://app.yourdomain.com",
	}
)

// GetAllowedOrigins returns the appropriate list of allowed origins based on the environment
func GetAllowedOrigins() []string {
	if IsDevelopment() {
		return DevelopmentOrigins
	}
	return ProductionOrigins
}

// IsDevelopment checks if the current environment is development
// You can implement this based on your own logic, e.g., checking an ENV variable
func IsDevelopment() bool {
	// Implement your logic here
	// For example:
	return os.Getenv("GO_ENV") == "dev"
	// return true // Default to development for this example
}
