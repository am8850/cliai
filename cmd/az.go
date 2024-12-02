package cmd

import (
	"fmt"

	"github.com/am8850/cliai/pkg/services"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// Add the root command to the application
var azCmd = &cobra.Command{
	Use:   "az",
	Short: "Generate and execute Azure CLI commands",
	Run: func(cmd *cobra.Command, args []string) {
		if prompt == "" {
			fmt.Println("Please provide a command. Example:")
			color.Cyan.Println("cliai az -p 'Show account information'")
			return
		}
		system_prompt := `You are an AI that can help generate Azure CLI (az) commands.

Rules:
- If the user requests something not related to az commands or operations, do not generate any commands.

No prologue or epilogue. Respond in the following JSON format:
[
	{ "command": "az", "args": ["account", "show"] },
]`
		services.Process(system_prompt, prompt, !confirm, list, &oaiSettings)
	},
}
