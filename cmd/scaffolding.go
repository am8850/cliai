package cmd

import (
	"fmt"

	"github.com/am8850/cliai/pkg/services"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// Create the version command
var cmdScaffod = &cobra.Command{
	Use:   "sc",
	Short: "Scafold code",
	Run: func(cmd *cobra.Command, args []string) {

		if prompt == "" {
			fmt.Println("Please provide a prompt. Example:")
			color.Cyan.Println("cliai sc -p 'Generate a Python script that prints Hello World'")
			return
		}

		system_prompt := `You are an AI that can help scaffold code in any programming language.

Rules:
- If the user requests something not related to scaffold code, do not generate any commands.
- Do your best to make the code very usable from the start.

No prologue or epilogue. Respond in the following JSON format:
[{
"filepath":"main.py",
"code":"print('Hello World')"
}]`
		services.Scafolder(system_prompt, prompt, &oaiSettings)
	},
}
