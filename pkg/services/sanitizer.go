package services

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gookit/color"
)

type SanitizerResponse struct {
	ReadabilityScore  int    `json:"readability_score"`
	ReadabilityReason string `json:"readability_reason"`
	CyclomaticScore   int    `json:"cyclomatic_score"`
	CyclomaticReason  string `json:"cyclomatic_reason"`
	ImprovedCode      string `json:"improved_code"`
}

func Sanitizer(system_prompt, file, output string, settings *OpenAISettings) {

	// Read the text in a file
	prompt, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading the input file:", err)
		return
	}

	if string(prompt) == "" {
		fmt.Println("The file is empty.")
		return
	}

	system := Message{Role: "system", Content: system_prompt}
	user := Message{Role: "user", Content: string(prompt)}
	messages := []Message{system, user}

	// Execute the chat completion
	jdata, err := ChatCompletion(messages, settings.ChatModel, 0.1, settings)
	if err != nil {
		fmt.Println("Unable to generate a completion with error:")
		color.Red.Println(err)
		return
	}
	//fmt.Println("JSON:\n", jdata)

	var sanitizedResponse SanitizerResponse
	err = json.Unmarshal([]byte(jdata), &sanitizedResponse)
	if err != nil {
		fmt.Println("Unable to parse the command with error:")
		color.Red.Println(err)
		return
	}

	fmt.Println("Propose code:\n")
	color.Cyan.Println(sanitizedResponse.ImprovedCode)

	if askForConfirmation("Do you want to write the file?") {
		// Write the sanitized code to a file
		if output != "" {
			err = os.WriteFile(output, []byte(sanitizedResponse.ImprovedCode), 0644)
		} else {
			// Add _sanitized to the file name before the extension
			fileParts := strings.Split(file, ".")
			if len(fileParts) >= 2 && fileParts[0] != "" && fileParts[1] != "" {
				file = fileParts[0] + "_sanitized." + fileParts[1]
				fmt.Println(file)
				err = os.WriteFile(file, []byte(sanitizedResponse.ImprovedCode), 0644)
			}
		}
	}

}
