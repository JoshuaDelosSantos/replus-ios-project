package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    AppPort    string
}

func LoadConfig() Config {
    // Load .env file if it exists
    if err := godotenv.Load(); err != nil {
        log.Println(".env file not found, using system environment variables")
    }

    return Config{
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     getEnv("DB_PORT", "5432"),
        DBUser:     getEnv("DB_USER", "postgres"),
        DBPassword: getEnv("DB_PASSWORD", "password"),
        DBName:     getEnv("DB_NAME", "mydb"),
        AppPort:    getEnv("APP_PORT", "8080"),
    }
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}