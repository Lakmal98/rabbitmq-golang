package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RabbitMQURL   string
	RedisAddr     string
	RedisPassword string
}

func LoadConfig() Config {
	// Load the local configuration first
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Printf("No local .env.local file found, falling back to .env")
	}

	// Load the generic configuration
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return Config{
		RabbitMQURL:   os.Getenv("RABBITMQ_URL"),
		RedisAddr:     os.Getenv("REDIS_ADDR"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
	}
}
