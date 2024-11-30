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

var (
	app     App
	confirm bool
	prompt  string
	list    bool
)

func init() {

	// Read the configuration JSON file
	bytes, err := os.ReadFile("./openai.json")
	if err != nil {
		// Get the executable path and directory
		exePath, err := os.Executable()
		if err != nil {
			log.Fatal(err)
		}
		exeDir := filepath.Dir(exePath)

		// Read configuration JSON file
		bytes, err = os.ReadFile(exeDir + "/openai.json")
		if err != nil {
			log.Fatal(err)
		}
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
		Short: "cliai - A simple CLI AI helper for git, az, and kubernetes",
		Long:  `cliai is a simple CLI AI helper that can generate git, Azure CLI (az), and Kubernetes (kubectl) commands based on user prompts.`,
		Run: func(cmd *cobra.Command, args []string) {
			text := ` ██████╗██╗     ██╗ █████╗ ██╗
██╔════╝██║     ██║██╔══██╗██║
██║     ██║     ██║███████║██║
██║     ██║     ██║██╔══██║██║
╚██████╗███████╗██║██║  ██║██║
 ╚═════╝╚══════╝╚═╝╚═╝  ╚═╝╚═╝
                              `
			fmt.Println()
			fmt.Println(text)
			fmt.Println("CLI Commander: A simple CLI AI helper for git, az, and kubectl")
		},
	}

	// Add the root command to the application
	gitCmd := &cobra.Command{
		Use:   "git",
		Short: "Generate and execute git CLI commands",
		Run: func(cmd *cobra.Command, args []string) {
			if prompt == "" {
				fmt.Println("Please provide a command. Example:")
				color.Cyan.Println("cliai git -p 'List all branches'")
				return
			}

			system_prompt := `You are an AI that can help generate git commands.
Rules:
- If configuring the user name or email address, put the user name or email address in double quotes and configure locally unless the user specifies global.
- If the user requests something not related to git, do not generate any commands.

No prologue or epilogue. Respond in the following JSON format:
[
	{ "command": "git", "args": ["add", "."] },
]`
			Process(system_prompt, prompt, !confirm, list)
		},
	}

	// Add the root command to the application
	azCmd := &cobra.Command{
		Use:   "az",
		Short: "Generate and execute Azure CLI commands",
		Run: func(cmd *cobra.Command, args []string) {
			if prompt == "" {
				fmt.Println("Please provide a command. Example:")
				color.Cyan.Println("cliai az -p 'Show account information'")
				return
			}
			system_prompt := `You are an AI that can help generate Azure CLI (az) commands.

Rules:
- If the user requests something not related to az commands or operations, do not generate any commands.

No prologue or epilogue. Respond in the following JSON format:
[
	{ "command": "az", "args": ["account", "show"] },
]`
			Process(system_prompt, prompt, !confirm, list)
		},
	}

	// Add the root command to the application
	k8sCmd := &cobra.Command{
		Use:   "k8s",
		Short: "Generate and execute Kubernetes CLI commands",
		Run: func(cmd *cobra.Command, args []string) {
			if prompt == "" {
				fmt.Println("Please provide a command. Example:")
				color.Cyan.Println("cliai k8s -p 'List all pods'")
				return
			}
			system_prompt := `You are an AI that can help generate Kubernetes (kubctl) commands.

Rules:
- If the user requests something not related to kubectl commands or operations, do not generate any commands.

No prologue or epilogue. Respond in the following JSON format:
[
	{ "command": "kubectl", "args": ["get", "-A"] },
]`
			Process(system_prompt, prompt, !confirm, list)
		},
	}

	// Create the version command
	scCmd := &cobra.Command{
		Use:   "sc",
		Short: "Scafold code",
		Run: func(cmd *cobra.Command, args []string) {
			system_prompt := `You are an AI that can help scaffold code in any programming language.

Rules:
- If the user requests something not related to scaffold code, do not generate any commands.
- Do your best to make the code very usable from the start.

No prologue or epilogue. Respond in the following JSON format:
[{
"filepath":"main.py",
"code":"print('Hello World')"
}]`
			Scafolder(system_prompt, prompt)
		},
	}

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
	rootCmd.AddCommand(scCmd)
	rootCmd.AddCommand(versionCmd)

	rootCmd.PersistentFlags().StringVarP(&prompt, "prompt", "p", "", "natural language CLI prompt")
	rootCmd.PersistentFlags().BoolVarP(&confirm, "disable", "d", false, "disable command confirmation")
	rootCmd.PersistentFlags().BoolVarP(&list, "list", "l", false, "list all commands first")

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
