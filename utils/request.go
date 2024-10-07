package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func Post(url string, data interface{}, authorization string) (*http.Response, error) {
	// Convert data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error marshaling data: %s", err)
		return nil, err
	}

	// Create request
	var req *http.Request
	req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error creating request: %s", err)
		return nil, err
	}

	// Set headers, including secret token
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+authorization)

	// Send the request
	var resp *http.Response

	client := &http.Client{}
	resp, err = client.Do(req)
	LogOnError(err, "Error sending request")

	defer resp.Body.Close()
	return resp, err
}
