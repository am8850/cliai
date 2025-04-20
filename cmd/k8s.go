package cmd

import (
	"fmt"

	"github.com/am8850/cliai/pkg/config"
	"github.com/am8850/cliai/pkg/processor"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// Add the root command to the application
var cmdK8s = &cobra.Command{
	Use:     "k8s",
	Aliases: []string{"k", "kubectl"},
	Short:   "Generate and execute Kubernetes CLI commands",
	Run: func(cmd *cobra.Command, args []string) {
		if prompt == "" {
			fmt.Println("Please provide a command. Example:")
			color.Cyan.Println("cliai k8s -p 'List all pods'")
			return
		}

		template := config.FindTemplate("kubectl")
		system_prompt := template.SystemPrompt

		processor.GenerateCommands(system_prompt, prompt, !confirm, list)
	},
}
