package main

import (
	"context"
	"log"
	"time"

	"github.com/Lakmak98/rabbitmq-golang/internal/config"
	"github.com/Lakmak98/rabbitmq-golang/internal/rabbitmq"
	"github.com/Lakmak98/rabbitmq-golang/internal/redis"
)

func main() {
	cfg := config.LoadConfig()

	// Initialize Redis client
	redis.InitRedis(cfg.RedisAddr, cfg.RedisPassword)

	// Initialize RabbitMQ
	rabbitmq.InitRabbitMQ(cfg.RabbitMQURL)
	defer rabbitmq.Close()

	// Declare a queue
	rabbitmq.DeclareQueue("rabbit")

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Message to be sent
	body := []byte("Hello World!")

	// Publish the message with context
	rabbitmq.PublishMessage(ctx, "rabbit", body)

	// Increment the Redis counter for messages
	redis.IncrementMessageCount()

	// Log the sent message
	log.Printf(" [x] Sent %s\n", body)
}
