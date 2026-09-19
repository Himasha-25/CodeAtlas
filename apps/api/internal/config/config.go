package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	JWTSecret   string
	Port        string
	Env         string
	GroqAPIKey  string
}

func Load() *Config {
	_ = godotenv.Load()
	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://codeatlas:codeatlas@localhost:5432/codeatlas?sslmode=disable"),
		JWTSecret:   getEnv("JWT_SECRET", "dev-secret"),
		Port:        getEnv("PORT", "8080"),
		Env:         getEnv("ENV", "development"),
		GroqAPIKey:  os.Getenv("GROQ_API_KEY"), // intentionally no default; callers must check
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
