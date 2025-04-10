package cmd

import (
	"fmt"

	"github.com/am8850/cliai/pkg/services"
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

		system_prompt := findPrompt("kubectl")

		oaiSettings.ResponseFormat = "json_object"

		services.Process(system_prompt, prompt, !confirm, list, &oaiSettings)
	},
}
