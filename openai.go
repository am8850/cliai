package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ChatCompletion(messages []Message, model string, temperature float64) (string, error) {
	if model == "" {
		model = app.ChatModel
	}
	// Create a new payload
	payload := ChatRequest{
		Messages:    messages,
		Model:       model,
		Temperature: temperature,
	}

	// Marshal the payload into JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", app.Endpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Set the request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", app.Key)

	// Create a new HTTP client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Unmarshal the response into a struct
	var response ChatResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Print the response message
	return response.Choices[0].Message.Content, nil
}
