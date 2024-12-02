package main

type App struct {
	Endpoint  string `json:"endpoint"`
	Key       string `json:"key"`
	Version   string `json:"version"`
	ChatModel string `json:"chat_model"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Payload represents the JSON payload to be sent
type ChatRequest struct {
	Messages    []Message `json:"messages"`
	Model       string    `json:"model"`
	Temperature float64   `json:"temperature"`
}

// Response represents the JSON response from the API
type ChatResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

// Command represents the command to be executed
type Command struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

type CodeFile struct {
	Filepath string `json:"filepath"`
	Code     string `json:"code"`
}
