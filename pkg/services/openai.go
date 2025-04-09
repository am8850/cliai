package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	client = &http.Client{}
)

func ChatCompletion(messages []Message, model string, temperature float64, settings *OpenAISettings) (string, error) {
	if model == "" {
		model = settings.ChatModel
	}

	response_format := "json_object"
	if settings.ResponseFormat != "" {
		response_format = settings.ResponseFormat
	}

	// Create a new payload
	payload := ChatRequest{
		Messages:       messages,
		Model:          model,
		Temperature:    temperature,
		ResponseFormat: &ChatResponsFormatType{Type: response_format},
	}

	// Marshal the payload into JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	//fmt.Println("Calling openai API", string(jsonPayload), app.Endpoint)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", settings.Endpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Set the request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", settings.Key)

	// Create a new HTTP client
	//client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	//fmt.Println("Response Status:", resp.Status)
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Error: %s", resp.Status)
	}

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

func DisposeClient() {
	if client != nil {
		client.CloseIdleConnections()
	}
}
