package cmd

import (
	"fmt"

	"github.com/am8850/cliai/pkg/services"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// Add the root command to the application
var cmdK8s = &cobra.Command{
	Use:   "k8s",
	Short: "Generate and execute Kubernetes CLI commands",
	Run: func(cmd *cobra.Command, args []string) {
		if prompt == "" {
			fmt.Println("Please provide a command. Example:")
			color.Cyan.Println("cliai k8s -p 'List all pods'")
			return
		}
		system_prompt := `You are an AI that can help generate Kubernetes (kubctl) commands.

Rules:
- If the user requests something not related to kubectl commands or operations, do not generate any commands.

No prologue or epilogue. Respond in the following JSON format:
[
	{ "command": "kubectl", "args": ["get", "-A"] },
]`
		services.Process(system_prompt, prompt, !confirm, list, &oaiSettings)
	},
}
