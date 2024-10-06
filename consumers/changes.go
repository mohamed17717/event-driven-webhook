package consumers

import (
	"encoding/json"
	"event-driven-webhook/config"
	"event-driven-webhook/models"
	"fmt"
	"log"
)

func ConsumeChanges() {
	msgs, err := config.CHANNEL.Consume(
		config.QUEUE_CHANGES, // queue
		"",                   // consumer
		true,                 // auto-ack
		false,                // exclusive
		false,                // no-local
		false,                // no-wait
		nil,                  // arguments
	)

	if err != nil {
		fmt.Printf("Error consuming from queue: %s\n", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			taskData := string(d.Body)
			fmt.Printf(" [x] Received %s", taskData)

			var data map[string]interface{}
			json.Unmarshal([]byte(taskData), &data)
			actionId := data["action_id"]

			//// Get the list of subscribers for the action
			subscribers := models.GetSubscribersForAction(actionId)
			for _, subscriber := range subscribers {
				noficiation := map[string]interface{}{
					"action_id":    actionId,
					"webhook_link": subscriber.WebhookLink,
					"secret_token": subscriber.SecretToken,
					"data":         data["data"],
				}
				jsonData, _ := json.Marshal(noficiation)
				config.MQPublishNotifySubscribers(string(jsonData))
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
