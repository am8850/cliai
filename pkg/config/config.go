package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
)

const (
	MODEL_CONFIG_FILE     = "cliaiopenai.json"
	TEMPLATES_CONFIG_FILE = "cliaitemplates.json"
)

type ModelConfig struct {
	Endpoint       string `json:"endpoint"`
	Key            string `json:"key"`
	ChatModel      string `json:"chat_model"`
	ResponseFormat string `json:"response_format"`
}

type Template struct {
	Command      string `json:"command"`
	SystemPrompt string `json:"system"`
}

type Config struct {
	ModelConfig ModelConfig
	Templates   []Template
}

var (
	instance *Config
	once     sync.Once
)

func readfile(fileName string) ([]byte, error) {

	file := filepath.Join(".", fileName)

	bytes, err := os.ReadFile(file)
	if err == nil {
		return bytes, nil
	}

	execPath, err := os.Executable()
	file = filepath.Join(execPath, fileName)

	bytes, err = os.ReadFile(file)
	if err == nil {
		return bytes, nil
	}
	return nil, err

}

func FindTemplate(command string) *Template {
	for _, p := range instance.Templates {
		if p.Command == command {
			return &p
		}
	}
	return nil
}

// GetConfig returns the singleton instance of Config by loading from configuration files
func GetConfig() (*Config, error) {
	var err error
	once.Do(func() {
		instance = &Config{}

		// Load OpenAI settings
		modelJSON, err := readfile(MODEL_CONFIG_FILE)
		if err != nil {
			panic("Unable to read file: " + MODEL_CONFIG_FILE)
		}

		// Unmarshal OpenAI modelConfig
		var modelConfig ModelConfig
		if err = json.Unmarshal(modelJSON, &modelConfig); err != nil {
			panic("Unable to parse file: " + MODEL_CONFIG_FILE)
		}

		if modelConfig.Endpoint == "" || modelConfig.Key == "" || modelConfig.ChatModel == "" {
			panic("Missing required fields endpoint, key, or chat_model in " + MODEL_CONFIG_FILE)
		}

		instance.ModelConfig = modelConfig

		// Load system prompts
		templatesDtata, err := readfile(TEMPLATES_CONFIG_FILE)
		if err != nil {
			panic("Unable to read file: " + TEMPLATES_CONFIG_FILE)
		}

		// Unmarshal system templates
		var templates []Template
		if err = json.Unmarshal(templatesDtata, &templates); err != nil {
			panic("Unable to parse file: " + TEMPLATES_CONFIG_FILE)
		}
		if len(templates) == 0 {
			panic("No templates found in " + TEMPLATES_CONFIG_FILE)
		}
		instance.Templates = templates

	})

	return instance, err
}
