package config

import (
	"os"
)

type Config struct {
	DatabaseURL string
}

// Load reads configuration from environment variables.
func Load() *Config {
	// If a full DATABASE_URL is provided, prefer it. Otherwise build from DB_* vars.
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		pass := os.Getenv("DB_PASSWORD")
		name := os.Getenv("DB_NAME")
		ssl := os.Getenv("DB_SSLMODE")
		if host != "" && port != "" && user != "" && name != "" {
			if ssl == "" {
				ssl = "disable"
			}
			// Build a standard PostgreSQL URL
			dbURL = "postgresql://" + user + ":" + pass + "@" + host + ":" + port + "/" + name + "?sslmode=" + ssl
		}
	}
	return &Config{DatabaseURL: dbURL}
}
