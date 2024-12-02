package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/am8850/cliai/pkg/services"
	"github.com/spf13/cobra"
)

var (
	confirm     bool
	prompt      string
	list        bool
	oaiSettings services.OpenAISettings
)

// Create the root command
var RootCmd = &cobra.Command{
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
	err = json.Unmarshal(bytes, &oaiSettings)
	if err != nil {
		log.Fatal(err)
	}

	// Add the commands to the root command
	RootCmd.AddCommand(gitCmd)
	RootCmd.AddCommand(azCmd)
	RootCmd.AddCommand(k8sCmd)
	RootCmd.AddCommand(scCmd)
	RootCmd.AddCommand(versionCmd)

	// Add the flags to the root command
	RootCmd.PersistentFlags().StringVarP(&prompt, "prompt", "p", "", "natural language CLI prompt")
	RootCmd.PersistentFlags().BoolVarP(&confirm, "disable", "d", false, "disable command confirmation")
	RootCmd.PersistentFlags().BoolVarP(&list, "list", "l", false, "list all commands first")
}
