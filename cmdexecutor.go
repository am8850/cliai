package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"

	"github.com/gookit/color"
)

func ExecShell(confirm bool, command string, args []string) {

	fmt.Println("Execute command:")
	//sargs := strings.Join(args, "")
	color.Cyan.Println("  ", command, args)
	if confirm {
		confirm = askForConfirmation("Do you want to execute the command?")
		if !confirm {
			fmt.Println("")
			return
		}
	}
	// Execute the command
	out, err := exec.Command(command, args...).Output()
	if err != nil {
		log.Fatal(err)
	}
	// Print the output
	fmt.Println("\nOutput:")
	color.Green.Print(string(out))
}

func Process(systemMessage, prompt string, confirm bool) {
	system := Message{Role: "system", Content: systemMessage}
	user := Message{Role: "user", Content: prompt}
	messages := []Message{system, user}
	jdata, err := ChatCompletion(messages, app.ChatModel, 0.1)
	if err != nil {
		log.Println(err)
		return
	}
	//fmt.Println(jdata)

	var commands []Command
	err = json.Unmarshal([]byte(jdata), &commands)
	if err != nil {
		log.Println("Unable to parse the command:", err)
		return
	}
	//fmt.Println(confirm)
	for _, command := range commands {
		ExecShell(confirm, command.Command, command.Args)
	}
}
