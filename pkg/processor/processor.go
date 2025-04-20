package processor

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"

	"github.com/am8850/cliai/pkg/config"
	"github.com/am8850/cliai/pkg/console"
	"github.com/am8850/cliai/pkg/openai"
	"github.com/gookit/color"
)

func execShell(confirm bool, command string, args []string) {

	fmt.Print("Command: ")
	//sargs := strings.Join(args, "")
	color.Cyan.Println(command, args)
	if confirm {
		confirm = console.AskForConfirmation("Do you want to execute this command?")
		if !confirm {
			fmt.Println("")
			return
		}
	}

	// Execute the command
	if command == "kubectl" || command == "az" || command == "git" || command == "docker" {
		out, err := exec.Command(command, args...).Output()
		if err != nil {
			log.Fatal(err)
		}
		// Print the output
		fmt.Println()
		color.Green.Print(string(out))
	}
}

func Any(system, prompt, format string) {
	// Create the system and user messages
	systemMessage := config.Message{Role: "system", Content: system}
	userMessage := config.Message{Role: "user", Content: prompt}
	messages := []config.Message{systemMessage, userMessage}

	// Execute the chat completion
	res, err := openai.ChatCompletion(messages, format)
	if err != nil {
		fmt.Println("Unable to generate a completion with error:")
		color.Red.Println(err)
		return
	}

	fmt.Println("\n", res)
}

func GenerateCommands(systemMessage, prompt string, confirm, list bool) {
	// Create the system and user messages
	system := config.Message{Role: "system", Content: systemMessage}
	user := config.Message{Role: "user", Content: prompt}
	messages := []config.Message{system, user}

	// Execute the chat completion
	jdata, err := openai.ChatCompletion(messages, "json_object")
	if err != nil {
		fmt.Println("Unable to generate a completion with error:")
		color.Red.Println(err)
		return
	}

	//fmt.Println("Payload:\n", jdata)

	// Unmarshal the JSON data into a slice of commands
	var commands config.Commands
	err = json.Unmarshal([]byte(jdata), &commands)
	if err != nil {
		fmt.Println("Unable to parse the command with error:")
		color.Red.Println(err)
		fmt.Println("Failed Payload:\n", jdata)
		return
	}

	// if requested, list the generated commands first
	if list {
		fmt.Println("Generated commands:")
		for _, command := range commands.Commands {
			color.Cyan.Println(command.Command, command.Args, "->", command.Explanation)
		}
		fmt.Println()
	}

	// Execute the commands
	for _, command := range commands.Commands {
		execShell(confirm, command.Command, command.Args)
	}
}
