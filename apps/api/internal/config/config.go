package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL    string
	JWTSecret      string
	Port           string
	Env            string
	OllamaBaseURL  string
	OllamaModel    string
}

func Load() *Config {
	_ = godotenv.Load()
	return &Config{
		DatabaseURL:   getEnv("DATABASE_URL", "postgres://codeatlas:codeatlas@localhost:5432/codeatlas?sslmode=disable"),
		JWTSecret:     getEnv("JWT_SECRET", "dev-secret"),
		Port:          getEnv("PORT", "8080"),
		Env:           getEnv("ENV", "development"),
		OllamaBaseURL: getEnv("OLLAMA_BASE_URL", "http://localhost:11434"),
		OllamaModel:   getEnv("OLLAMA_MODEL", "llama3"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
