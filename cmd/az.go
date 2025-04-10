package cmd

import (
	"fmt"

	"github.com/am8850/cliai/pkg/services"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// Add the root command to the application
var cmdAzCli = &cobra.Command{
	Use:   "az",
	Short: "Generate and execute Azure CLI commands",
	Run: func(cmd *cobra.Command, args []string) {
		if prompt == "" {
			fmt.Println("Please provide a command. Example:")
			color.Cyan.Println("cliai az -p 'Show account information'")
			return
		}

		system_prompt := findPrompt("az")

		oaiSettings.ResponseFormat = "json_object"

		services.Process(system_prompt, prompt, !confirm, list, &oaiSettings)
	},
}
