package config

import (
	"errors"
	"event-driven-webhook/utils"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
)

var AMQP *amqp.Connection
var CHANNEL *amqp.Channel
var QUEUE_CHANGES = "changes"
var QUEUE_NOTIFY_SUBSCRIBERS = "notify_subscribers"
var QUEUE_WEBHOOK_FAILURE = "webhook_failure"

func createQueue(name string) (amqp.Queue, error) {
	return CHANNEL.QueueDeclare(
		name,  // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
}

func publish(queue, msg string) error {
	return CHANNEL.Publish(
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
}

func MQPublishChange(msg string) error {
	return publish(QUEUE_CHANGES, msg)
}

func MQPublishNotifySubscribers(msg string) error {
	return publish(QUEUE_NOTIFY_SUBSCRIBERS, msg)
}

func MQPublishWebhookFailure(msg string) error {
	return publish(QUEUE_WEBHOOK_FAILURE, msg)
}

func CreateConsumer(queueName string) <-chan amqp.Delivery {
	msgs, err := CHANNEL.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // arguments
	)
	utils.LogOnError(err, "Failed to register a consumer")
	return msgs
}

func ConnectToRabbitMQ() {
	var err error

	rabbitMqUrl := os.Getenv("RABBITMQ_URL")
	utils.FailOnError(errors.New("missed env variables"), "RABBITMQ_URL not set")

	AMQP, err = amqp.Dial(rabbitMqUrl)
	utils.FailOnError(err, "Failed to connect to RabbitMQ")

	CHANNEL, err = AMQP.Channel()
	utils.FailOnError(err, "Failed to open a channel")

	_, err = createQueue(QUEUE_CHANGES)
	utils.FailOnError(err, "Failed to create queue")

	_, err = createQueue(QUEUE_NOTIFY_SUBSCRIBERS)
	utils.FailOnError(err, "Failed to create queue")

	_, err = createQueue(QUEUE_WEBHOOK_FAILURE)
	utils.FailOnError(err, "Failed to create queue")
}
