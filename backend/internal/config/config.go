package config

import (
	"os"
	"strconv"
)

// Config stores runtime configuration for the Fairy Castle API.
type Config struct {
	AppEnv   string
	HTTPPort string
	LogLevel string

	MySQLDSN string
	RedisURL string

	JWTSecret  string
	AIProvider string
}

// Load reads configuration from environment variables.
func Load() Config {
	return Config{
		AppEnv:     getEnv("APP_ENV", "development"),
		HTTPPort:   getEnv("HTTP_PORT", "8080"),
		LogLevel:   getEnv("LOG_LEVEL", "debug"),
		MySQLDSN:   getEnv("MYSQL_DSN", "fairy:fairy@tcp(mysql:3306)/fairy_castle?charset=utf8mb4&parseTime=True&loc=UTC"),
		RedisURL:   getEnv("REDIS_URL", "redis://redis:6379/0"),
		JWTSecret:  getEnv("JWT_SECRET", "change-me-in-production"),
		AIProvider: getEnv("AI_PROVIDER", "mock"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

// GetInt reads an integer environment variable with fallback.
func GetInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return parsed
}
