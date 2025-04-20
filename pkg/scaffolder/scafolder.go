package scaffolder

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/am8850/cliai/pkg/config"
	"github.com/am8850/cliai/pkg/console"
	"github.com/am8850/cliai/pkg/openai"
	"github.com/gookit/color"
)

func createFolderIfNotExists(filePath string) error {
	// Get the directory from the filepath
	dir := filepath.Dir(filePath)

	//fmt.Println("Creating directory:", dir)

	if dir == "." {
		return nil
	}

	if dir != "" {
		// Check if the directory exists
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			// If it doesn't exist, create it
			return os.MkdirAll(dir, os.ModePerm)
		}
	}

	return nil
}

func Scaffold(system_prompt, prompt string) {
	// Create the system and user messages
	system := config.Message{Role: "system", Content: system_prompt}
	user := config.Message{Role: "user", Content: prompt}
	messages := []config.Message{system, user}

	// Execute the chat completion
	jdata, err := openai.ChatCompletion(messages, "json_object")
	if err != nil {
		fmt.Println("Unable to generate a completion with error:")
		color.Red.Println(err)
		return
	}

	//fmt.Println("JSON:\n", jdata)

	// Unmarshal the JSON data into a slice of commands
	var codefiles config.CodeFiles
	err = json.Unmarshal([]byte(jdata), &codefiles)
	if err != nil {
		fmt.Println("Unable to parse the command with error:")
		color.Red.Println(err)
		fmt.Println("Failed Payload:\n", jdata)
		return
	}

	//fmt.Println("Generated code files:", codefiles)
	fmt.Print("Generated code:\n\n")
	for _, codefile := range codefiles.Files {
		color.Yellow.Println("File: " + codefile.Filepath)
		color.Cyan.Println(codefile.Code + "\n")
	}

	if console.AskForConfirmation("Do you want to write files?") {
		for _, codefile := range codefiles.Files {
			err := createFolderIfNotExists(codefile.Filepath)
			if err != nil {
				color.Red.Println("Error creating directory:", err)
				return
			}
			err = os.WriteFile(codefile.Filepath, []byte(codefile.Code), 0644)
			if err != nil {
				color.Red.Println("Error writing file:", err)
			}
		}
	}

}
