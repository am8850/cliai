package cmd

import (
	"fmt"

	"github.com/am8850/cliai/pkg/config"
	"github.com/am8850/cliai/pkg/processor"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// Add the root command to the application
var cmdAzcopy = &cobra.Command{
	Use:   "azcopy",
	Short: "Generate and execute azcopy commands",
	Run: func(cmd *cobra.Command, args []string) {
		if prompt == "" {
			fmt.Println("Please provide a command. Example:")
			color.Cyan.Println("cliai azcopy -p 'List all blobs in a container'")
			return
		}

		template := config.FindTemplate("azcopy")
		system_prompt := template.SystemPrompt

		processor.GenerateCommands(system_prompt, prompt, !confirm, list)
	},
}
