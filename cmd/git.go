package cmd

import (
	"fmt"

	"github.com/am8850/cliai/pkg/config"
	"github.com/am8850/cliai/pkg/processor"
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

		template := config.FindTemplate("git")
		system_prompt := template.SystemPrompt

		processor.GenerateCommands(system_prompt, prompt, !confirm, list)
	},
}
