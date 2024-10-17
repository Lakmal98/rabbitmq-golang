package rabbitmq

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var conn *amqp.Connection
var ch *amqp.Channel

func InitRabbitMQ(url string) {
	var err error
	conn, err = amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")
}

func DeclareQueue(queueName string) {
	_, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")
}

func PublishMessage(ctx context.Context, queueName string, body []byte) {
	err := ch.PublishWithContext(ctx,
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
	failOnError(err, "Failed to publish a message")
}

func ConsumeMessages(queueName string) <-chan amqp.Delivery {
	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")
	return msgs
}

func Close() {
	if err := ch.Close(); err != nil {
		log.Fatalf("Failed to close channel: %s", err)
	}
	if err := conn.Close(); err != nil {
		log.Fatalf("Failed to close connection: %s", err)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
