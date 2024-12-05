package services

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"

	"github.com/gookit/color"
)

func askForConfirmation(s string) bool {
	var response string
	color.Yellow.Printf("%s (y/n): ", s)
	fmt.Scanln(&response)
	return response == "y" || response == "Y"
}

func execShell(confirm bool, command string, args []string) {

	fmt.Print("Generated command: ")
	//sargs := strings.Join(args, "")
	color.Cyan.Println(command, args)
	if confirm {
		confirm = askForConfirmation("Do you want to execute the command?")
		if !confirm {
			fmt.Println("")
			return
		}
	}

	// Execute the command
	if command == "kubectl" || command == "az" || command == "git" {
		out, err := exec.Command(command, args...).Output()
		if err != nil {
			log.Fatal(err)
		}
		// Print the output
		fmt.Println()
		color.Green.Print(string(out))
	}
}

func Process(systemMessage, prompt string, confirm, list bool, settings *OpenAISettings) {
	// Create the system and user messages
	system := Message{Role: "system", Content: systemMessage}
	user := Message{Role: "user", Content: prompt}
	messages := []Message{system, user}

	// Execute the chat completion
	jdata, err := ChatCompletion(messages, settings.ChatModel, 0.1, settings)
	if err != nil {
		fmt.Println("Unable to generate a completion with error:")
		color.Red.Println(err)
		return
	}

	// Unmarshal the JSON data into a slice of commands
	var commands []Command
	err = json.Unmarshal([]byte(jdata), &commands)
	if err != nil {
		fmt.Println("Unable to parse the command with error:")
		color.Red.Println(err)
		return
	}

	// if requested, list the generated commands first
	if list {
		fmt.Println("Generated commands:")
		for _, command := range commands {
			color.Cyan.Println(command.Command, command.Args)
		}
		fmt.Println()
	}

	// Execute the commands
	for _, command := range commands {
		execShell(confirm, command.Command, command.Args)
	}
}
