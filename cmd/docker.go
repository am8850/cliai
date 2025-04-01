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
			color.Cyan.Println("cliai docker -p 'Show account information'")
			return
		}
		system_prompt := `You are an AI that can help generate docker and docker compose CLI commands.

Rules:
- If the user requests something not related to docker commands or operations, do not generate any commands.

No prologue or epilogue. Respond in the following JSON format:
[
	{ "command": "docker", "args": ["image", "ls"] },
]`
		services.Process(system_prompt, prompt, !confirm, list, &oaiSettings)
	},
}
