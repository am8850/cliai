package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	confirm bool
	prompt  string
	list    bool
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

func init() {
	// Add the commands to the root command
	RootCmd.AddCommand(cmdGit)
	RootCmd.AddCommand(cmdDocker)
	RootCmd.AddCommand(cmdAzCli)
	RootCmd.AddCommand(cmdAzcopy)
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
