package cmd

import (
	"fmt"

	"github.com/am8850/cliai/pkg/services"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// Add the root command to the application
var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Generate and execute git CLI commands",
	Run: func(cmd *cobra.Command, args []string) {
		if prompt == "" {
			fmt.Println("Please provide a command. Example:")
			color.Cyan.Println("cliai git -p 'List all branches'")
			return
		}

		system_prompt := `You are an AI that can help generate git commands.
Rules:
- If configuring the user name or email address, put the user name or email address in double quotes and configure locally unless the user specifies global.
- If the user requests something not related to git, do not generate any commands.

No prologue or epilogue. Respond in the following JSON format:
[
{ "command": "git", "args": ["add", "."] },
]`
		services.Process(system_prompt, prompt, !confirm, list, &oaiSettings)
	},
}
