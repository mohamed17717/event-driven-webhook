package consumers

import (
	"encoding/json"
	"event-driven-webhook/config"
	"event-driven-webhook/utils"
	"fmt"
	"log"
)

func ConsumeNotifySubscribers() {
	consumeMsgs := config.CreateConsumer(config.QUEUE_NOTIFY_SUBSCRIBERS)
	forever := make(chan bool)

	go func() {
		for d := range consumeMsgs {
			taskData := string(d.Body)
			fmt.Printf(" [x] Received %s", taskData)

			var payload map[string]interface{}
			json.Unmarshal([]byte(taskData), &payload)

			// Extract required data
			webhookLink := payload["webhook_link"].(string)
			secretToken := payload["secret_token"].(string)
			actionData := payload["data"]

			// Prepare and send webhook request
			utils.Post(webhookLink, actionData, secretToken)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
