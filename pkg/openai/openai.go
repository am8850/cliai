package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/am8850/cliai/pkg/config"
)

var (
	client = &http.Client{}
)

func ChatCompletion(messages []config.Message, responseFormat string) (string, error) {
	conf, _ := config.GetConfig()
	if responseFormat == "" {
		responseFormat = "text"
	}

	// Create a new payload
	payload := config.ChatRequest{
		Messages:       messages,
		Model:          conf.ModelConfig.Model,
		Temperature:    conf.ModelConfig.Temperature,
		ResponseFormat: &config.ChatResponseFormatType{Type: responseFormat},
	}

	// Marshal the payload into JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	//fmt.Println("Calling openai API", string(jsonPayload), app.Endpoint)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", conf.ModelConfig.Endpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Set the request headers
	req.Header.Set("Content-Type", "application/json")
	if conf.ModelConfig.Type == "openai" {
		req.Header.Set("Authorization", "Bearer "+conf.ModelConfig.Key)
	} else {
		req.Header.Set("api-key", conf.ModelConfig.Key)
	}

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
	var response config.ChatResponse
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
