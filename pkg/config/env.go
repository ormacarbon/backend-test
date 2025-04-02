package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config struct to abstract the environment variables
type Config struct {
	POSTGRES_HOST     string
	POSTGRES_PORT     int
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB	      string
	FIBER_PORT     	  string
	SMTP_EMAIL         string
	SMTP_HOST          string
	SMTP_PORT          int
	SMTP_USER          string
	SMTP_PASSWORD      string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := &Config{
		POSTGRES_HOST: getEnv("POSTGRES_HOST", "localhost"),
		POSTGRES_PORT: getEnvAsInt("POSTGRES_PORT", 5432),
		POSTGRES_USER: getEnv("POSTGRES_USER", "root"),
		POSTGRES_PASSWORD: getEnv("POSTGRES_PASSWORD", "root"),
		POSTGRES_DB: getEnv("POSTGRES_DB", "gss-db"),
		FIBER_PORT: getEnv("PORT", "3000"),
		SMTP_EMAIL: getEnv("SMTP_EMAIL", ""),
		SMTP_PORT: getEnvAsInt("SMTP_PORT", 587),
		SMTP_HOST: getEnv("SMTP_HOST", ""),
		SMTP_PASSWORD: getEnv("SMTP_PASSWORD", ""),

	}

	return config, nil
}

// Auxiliar functions to get the environment variables
func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}

func getEnvAsInt(name string, defaultValue int) int {
    valueStr := getEnv(name, "")
    if value, err := strconv.Atoi(valueStr); err == nil {
        return value
    }
    return defaultValue
}
