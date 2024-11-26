package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var app App

func askForConfirmation(s string) bool {
	var response string
	color.Yellow.Printf("%s (y/n): ", s)
	fmt.Scanln(&response)
	return response == "y" || response == "Y"
}

func init() {
	// Get the executable path and directory
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exeDir := filepath.Dir(exePath)

	// Read JSON file
	bytes, err := os.ReadFile(exeDir + "/openai.json")
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the configuration JSON
	err = json.Unmarshal(bytes, &app)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	// Create the root command
	rootCmd := &cobra.Command{
		Use:   "cliai",
		Short: "A simple CLI AI helper for git, az, and kubernetes",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to CLIAI!")
		},
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
- If configuring the user name or email address, put the user name or email address in double quotes and configure locally unless the user specifies global.

No prologue or epilogue. Respond in the following JSON format:
[
	{ "command": "git", "args": ["add", "."] },
]`
			Process(system, prompt, confirm == "y")
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
			Process(system, prompt, true)
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
			Process(system, prompt, true)
		},
	}
	k8sCmd.Flags().StringP("prompt", "p", "", "Pass a prompt to the AI model")

	// Create the version command
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the application version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("CLIAI version 1.0")
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
