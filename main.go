package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

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

var app App

func askForConfirmation(s string) bool {
	var response string
	color.Yellow.Printf("%s (y/n): ", s)
	fmt.Scanln(&response)
	return response == "y" || response == "Y"
}

func exec_shell(confirm bool, command string, args []string) {

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

func chat_completion(messages []Message, model string, temperature float64) (string, error) {
	if model == "" {
		model = app.ChatModel
	}
	// Create a new payload
	payload := ChatRequest{
		Messages:    messages,
		Model:       model,
		Temperature: temperature,
	}

	// Marshal the payload into JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", app.Endpoint, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// Set the request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", app.Key)

	// Create a new HTTP client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
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

func init() {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Executable path:", exePath)
	exeDir := filepath.Dir(exePath)

	// Read JSON file
	bytes, err := os.ReadFile(exeDir + "/openai.json")
	if err != nil {
		log.Fatal(err)
	}
	// Unmarshal JSON
	err = json.Unmarshal(bytes, &app)
	if err != nil {
		log.Fatal(err)
	}
	//log.Print("Config file loaded", app)
}

type Command struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

func process(systemMessage, prompt string, confirm bool) {
	system := Message{Role: "system", Content: systemMessage}
	user := Message{Role: "user", Content: prompt}
	messages := []Message{system, user}
	jdata, err := chat_completion(messages, app.ChatModel, 0.1)
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
		exec_shell(confirm, command.Command, command.Args)
	}
}

// var jdata string = `[
//         { "command": "git", "args": ["checkout", "main"] },
//         { "command": "git", "args": ["reset", "--hard"] }
// ]`

func main() {
	// var commands []Command
	// err := json.Unmarshal([]byte(jdata), &commands)
	// if err != nil {
	// 	log.Fatal("Unable to parse the command:", err)
	// 	return
	// }

	// Create the root command
	rootCmd := &cobra.Command{
		Use:   "cliai",
		Short: "A simple CLI AI helper for git, az, and kubernetes",
	}

	// Add the root command to the application
	gitCmd := &cobra.Command{
		Use:   "git",
		Short: "execute git CLI commands",
		Run: func(cmd *cobra.Command, args []string) {
			prompt, err := cmd.Flags().GetString("prompt")
			if err != nil {
				log.Fatal(err)
			}
			if prompt == "" {
				log.Println("You need to provide a prompt")
				return
			}

			confirm, err := cmd.Flags().GetString("confirm")
			if err != nil {
				log.Println("Unable to get the confirm flag:", err)
				return
			}
			system := `You are an AI that can help generate git commands.
Rules:
- If configuring the user name or email address, put the user name or email address in double quotes.

No prologue or epilogue. Respond in the following JSON format:
[
	{ "command": "git", "args": ["add", "."] },
]`
			process(system, prompt, confirm == "y")
		},
	}
	gitCmd.Flags().StringP("prompt", "p", "", "Prompt to execute")
	gitCmd.Flags().StringP("confirm", "c", "y", "Confirm execution")

	// Add the root command to the application
	azCmd := &cobra.Command{
		Use:   "az",
		Short: "az - execute Azure CLI commands",
		Run: func(cmd *cobra.Command, args []string) {
			prompt, err := cmd.Flags().GetString("prompt")
			if err != nil {
				log.Fatal(err)
			}
			if prompt != "" {
				fmt.Println("You entered:", prompt)
			}
			system := `You are an AI that can help generate Azure CLI (az) commands.

No prologue or epilogue. Respond in the following JSON format:
[
	{ "command": "az", "args": ["account", "show"] },
]`
			process(system, prompt, true)
		},
	}
	azCmd.Flags().StringP("prompt", "p", "", "Pass a prompt to the AI model")

	// Add the root command to the application
	k8sCmd := &cobra.Command{
		Use:   "k8s",
		Short: "k8s - execute Kubernetes CLI commands",
		Run: func(cmd *cobra.Command, args []string) {
			prompt, err := cmd.Flags().GetString("prompt")
			if err != nil {
				log.Fatal(err)
			}
			if prompt != "" {
				fmt.Println("You entered:", prompt)
			}
			system := `You are an AI that can help generate Kubernetes (kubctl) commands.

No prologue or epilogue. Respond in the following JSON format:
[
	{ "command": "kubectl", "args": ["get", "-A"] },
]`
			process(system, prompt, true)
		},
	}
	k8sCmd.Flags().StringP("prompt", "p", "", "Pass a prompt to the AI model")

	// Create the version command
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the application version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("MyApp version 1.0")
		},
	}

	// Add the commands to the root command
	rootCmd.AddCommand(gitCmd)
	rootCmd.AddCommand(azCmd)
	rootCmd.AddCommand(k8sCmd)
	rootCmd.AddCommand(versionCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
