package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	UserServiceURL string
	PostServiceURL string
	Port           string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using system env variables")
	}

	cfg := &Config{
		UserServiceURL: getEnv("USER_SERVICE_URL", "http://localhost:8081"),
		PostServiceURL: getEnv("POST_SERVICE_URL", "http://localhost:8082"),
		Port:           getEnv("PORT", "8080"),
	}

	return cfg
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
