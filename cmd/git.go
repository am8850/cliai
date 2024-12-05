package cmd

import (
	"fmt"

	"github.com/am8850/cliai/pkg/services"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// Add the root command to the application
var cmdGit = &cobra.Command{
	Use:   "git",
	Short: "Generate and execute git CLI commands",
	Run: func(cmd *cobra.Command, args []string) {
		if prompt == "" {
			fmt.Println("Please provide a command. Example:")
			color.Cyan.Println("cliai git -p 'List all branches'")
			return
		}

		system_prompt := findPrompt("git")

		oaiSettings.ResponseFormat = "text"

		services.Process(system_prompt, prompt, !confirm, list, &oaiSettings)
	},
}
