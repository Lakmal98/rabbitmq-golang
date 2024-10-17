package main

import (
	"log"

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

	// Declare the same queue to ensure it exists
	rabbitmq.DeclareQueue("rabbit")

	// Consume messages from the queue
	msgs := rabbitmq.ConsumeMessages("rabbit")

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	// Goroutine to handle incoming messages
	go func() {
		for d := range msgs {
			// Print the received message
			log.Printf("Received a message: %s", d.Body)

			// Read the message count from Redis
			count := redis.GetMessageCount()

			// Log the current message count
			log.Printf("Total messages sent: %d", count)
		}
	}()

	select {} // Keep the program alive
}
