package config

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
)

var AMQP *amqp.Connection
var CHANNEL *amqp.Channel
var QUEUE_CHANGES = "changes"
var QUEUE_NOTIFY_SUBSCRIBERS = "notify_subscribers"
var QUEUE_WEBHOOK_FAILURE = "webhook_failure"

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

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
		"",    // exchange
		queue, // queue name
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		}, // message to publish
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

func ConnectToRabbitMQ() {
	rabbitMqUrl := os.Getenv("RABBITMQ_URL")
	if rabbitMqUrl == "" {
		fmt.Println("RABBITMQ_URL not set")
		return
	}

	var err error

	AMQP, err = amqp.Dial(rabbitMqUrl)
	failOnError(err, "Failed to connect to RabbitMQ")

	CHANNEL, err = AMQP.Channel()
	failOnError(err, "Failed to open a channel")

	_, err = createQueue(QUEUE_CHANGES)
	failOnError(err, "Failed to create queue")

	_, err = createQueue(QUEUE_NOTIFY_SUBSCRIBERS)
	failOnError(err, "Failed to create queue")

	_, err = createQueue(QUEUE_WEBHOOK_FAILURE)
	failOnError(err, "Failed to create queue")
}
