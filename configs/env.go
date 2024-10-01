package config

import "os"

// Config struct holds all configuration values needed for the application
type Config struct {
	PublicHost string // The public-facing hostname of the server (e.g., "https://localhost")
	Port       string // The port the server listens on (e.g., "8080")
	DBUser     string // The database username
	DBPassword string // The database password
	DBAddress  string // The database host address (e.g., "localhost")
	DBName     string // The database name (e.g., "pivault")
}

// Envs is a global variable holding the environment configuration values
// It is initialized by calling InitConfig(), which fetches the environment variables
var Envs = InitConfig()

// InitConfig initializes and returns the configuration struct by fetching environment variables
func InitConfig() Config {
	// Returns a Config struct, populating fields with environment variables or fallback defaults
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "https://localhost"), // Retrieves PUBLIC_HOST or defaults to "https://localhost"
		Port:       getEnv("PORT", "8080"),                     // Retrieves PORT or defaults to "8080"
		DBUser:     getEnv("DB_USER", "root"),                  // Retrieves DB_USER or defaults to "root"
		DBPassword: getEnv("DB_PASSWORD", ""),                  // Retrieves DB_PASSWORD or defaults to an empty string
		DBAddress:  getEnv("DB_HOST", "localhost"),             // Retrieves DB_HOST or defaults to "localhost"
		DBName:     getEnv("DB_NAME", "pivault"),               // Retrieves DB_NAME or defaults to "pivault"
	}
}

// getEnv checks if an environment variable is set, otherwise returns a fallback value
// `key` is the name of the environment variable, `fallback` is the default value
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value // If the environment variable is set, return its value
	}
	return fallback // If not set, return the provided fallback value
}
