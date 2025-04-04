package cmd

import (
	"fmt"

	"github.com/am8850/cliai/pkg/services"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// Add the root command to the application
var dockerCmd = &cobra.Command{
	Use:     "docker",
	Short:   "Generate and execute Docker commands",
	Aliases: []string{"do"},
	Run: func(cmd *cobra.Command, args []string) {
		if prompt == "" {
			fmt.Println("Please provide a command. Example:")
			color.Cyan.Println("cliai docker -p 'List all running containers'")
			return
		}

		system_prompt := findPrompt("docker")

		oaiSettings.ResponseFormat = "text"

		services.Process(system_prompt, prompt, !confirm, list, &oaiSettings)
	},
}
