package consumers

import (
	"bytes"
	"encoding/json"
	"event-driven-webhook/config"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func ConsumeNotifySubscribers() {
	msgs, err := config.CHANNEL.Consume(
		config.QUEUE_NOTIFY_SUBSCRIBERS, // queue
		"",                              // consumer
		true,                            // auto-ack
		false,                           // exclusive
		false,                           // no-local
		false,                           // no-wait
		nil,                             // arguments
	)

	if err != nil {
		fmt.Printf("Error consuming from queue: %s\n", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			taskData := string(d.Body)
			fmt.Printf(" [x] Received %s", taskData)

			var payload map[string]interface{}
			json.Unmarshal([]byte(taskData), &payload)

			// Extract required data
			webhookLink := payload["webhook_link"].(string)
			secretToken := payload["secret_token"].(string)
			actionData := payload["data"]

			// Prepare and send webhook request
			sendWebhook(webhookLink, secretToken, actionData)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func sendWebhook(webhookLink, secretToken string, data interface{}) {
	// Convert data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error marshaling data: %s", err)
		return
	}

	// Create request
	req, err := http.NewRequest("POST", webhookLink, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error creating request: %s", err)
		return
	}

	// Set headers, including secret token
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+secretToken)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %s", err)
		return
	}
	defer resp.Body.Close()

	// Read and log the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response: %s", err)
		return
	}

	log.Printf("Webhook response: %s", string(body))
}
