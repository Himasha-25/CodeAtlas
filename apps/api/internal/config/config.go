package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
	GroqAPIKey string
	Port       string
	Env        string
}

func (c *Config) DatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET is required")
	}
	groqAPIKey := os.Getenv("GROQ_API_KEY")
	if groqAPIKey == "" {
		return nil, fmt.Errorf("GROQ_API_KEY is required")
	}

	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "codeatlas"),
		DBPassword: getEnv("DB_PASSWORD", "codeatlas"),
		DBName:     getEnv("DB_NAME", "codeatlas"),
		JWTSecret:  jwtSecret,
		GroqAPIKey: groqAPIKey,
		Port:       getEnv("PORT", "8080"),
		Env:        getEnv("ENV", "development"),
	}, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
