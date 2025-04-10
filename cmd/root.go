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
	prompts     []services.SystemPrompt
)

// Create the root command
var RootCmd = &cobra.Command{
	Use:   "cliai",
	Short: "cliai a simple CLI AI helper for git, az, kubernetes, scaffolding and refactoring code.",
	Long:  `cliai a simple CLI AI helper for generating git, az and kubectl commands from prompts. Also for scaffolding and refactoring code.`,
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
		fmt.Println("CLI AI Commander: A simple CLI AI helper for generating git, az and kubectl commands from prompts. Also for scaffolding and refactoring code.")
	},
}

func execFolder() string {
	// Get the executable path and directory
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting the executable path:", err)
		log.Fatal(err)
	}

	return filepath.Dir(exePath)
}

// commands are git, az, k8s, scaffod, refactor
func findPrompt(command string) string {
	for _, p := range prompts {
		if p.Command == command {
			return p.SystemPrompt
		}
	}

	return ""
}

func init() {

	// Read the configuration JSON file
	bytes, err := os.ReadFile("./cliaiopenai.json")
	if err != nil {
		execFolder := execFolder()

		// Read configuration JSON file
		bytes, err = os.ReadFile(execFolder + "/cliaiopenai.json")
		if err != nil {
			log.Fatal(err)
		}
	}

	// Unmarshal the configuration JSON
	err = json.Unmarshal(bytes, &oaiSettings)
	if err != nil {
		log.Fatal(err)
	}

	// Add the flags to the root command
	bytes, err = os.ReadFile("./cliaitemplates.json")
	if err != nil {
		execFolder := execFolder()

		// Read configuration JSON file
		bytes, err = os.ReadFile(execFolder + "/cliaitemplates.json")
		if err != nil {
			log.Fatal(err)
		}
	}

	// Unmarshal the configuration JSON
	err = json.Unmarshal(bytes, &prompts)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(prompts[0].Command, prompts[0].SystemPrompt)

	// Add the commands to the root command
	RootCmd.AddCommand(cmdGit)
	RootCmd.AddCommand(cmdDocker)
	RootCmd.AddCommand(cmdAzCli)
	RootCmd.AddCommand(cmdK8s)
	RootCmd.AddCommand(cmdScaffod)
	RootCmd.AddCommand(cmdRefactor)
	RootCmd.AddCommand(cmdVersion)
	RootCmd.AddCommand(cmdAny)

	// Add the flags to the root command
	RootCmd.PersistentFlags().StringVarP(&prompt, "prompt", "p", "", "natural language CLI prompt")
	RootCmd.PersistentFlags().BoolVarP(&confirm, "disable", "d", false, "disable command confirmation")
	RootCmd.PersistentFlags().BoolVarP(&list, "list", "l", true, "list all commands first")
}
