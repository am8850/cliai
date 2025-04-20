package refactorer

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/am8850/cliai/pkg/config"
	"github.com/am8850/cliai/pkg/console"
	"github.com/am8850/cliai/pkg/openai"
	"github.com/gookit/color"
)

func Refactor(system_prompt, file, output string) {

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

	system := config.Message{Role: "system", Content: system_prompt}
	user := config.Message{Role: "user", Content: string(prompt)}
	messages := []config.Message{system, user}

	// Execute the chat completion
	jdata, err := openai.ChatCompletion(messages, "json_object")
	if err != nil {
		fmt.Println("Unable to generate a completion with error:")
		color.Red.Println(err)
		return
	}
	//fmt.Println("JSON:\n", jdata)

	var sanitizedResponse config.SanitizerResponse
	err = json.Unmarshal([]byte(jdata), &sanitizedResponse)
	if err != nil {
		fmt.Println("Unable to parse the command with error:")
		color.Red.Println(err)
		fmt.Println("Failed Payload:\n", jdata)
		return
	}

	fmt.Printf("\nCode information:\n\n")

	fmt.Printf("Readability score: ")
	if sanitizedResponse.ReadabilityScore < 5 {
		color.Red.Println(sanitizedResponse.ReadabilityScore)
	} else {
		color.Cyan.Println(sanitizedResponse.ReadabilityScore)
	}
	fmt.Printf("Readability score reason:\n")
	color.Cyan.Println(sanitizedResponse.ReadabilityReason)

	fmt.Printf("\nCyclomatic complexity score: ")
	if sanitizedResponse.CyclomaticScore > 5 {
		color.Red.Println(sanitizedResponse.CyclomaticScore)
	} else {
		color.Cyan.Println(sanitizedResponse.CyclomaticScore)
	}
	fmt.Printf("Cyclomatic complexity score reason:\n")
	color.Cyan.Println(sanitizedResponse.CyclomaticReason)

	if !console.AskForConfirmation("\nContinue to view the proposed code?") {
		return
	}

	// fmt.Printf("\n\nOriginal code:\n\n")
	// color.Cyan.Println(string(prompt))

	fmt.Printf("\nProposed code changes:\n\n")
	color.Green.Println(sanitizedResponse.ImprovedCode)

	if console.AskForConfirmation("Write the code to a file?") {
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
