package cmd

import (
	"fmt"

	"github.com/am8850/cliai/pkg/services"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// Create the version command
var cmdScaffod = &cobra.Command{
	Use:     "scaffold",
	Short:   "Scafold code for any language, SQL, CSV, json, bash, devcontainer, etc.",
	Aliases: []string{"sc"},
	Run: func(cmd *cobra.Command, args []string) {

		if prompt == "" {
			fmt.Println("Please provide a prompt. Example:")
			color.Cyan.Println("cliai sc -p 'Generate a Python script that prints Hello World'")
			return
		}

		system_prompt := findPrompt("scaffold")
		oaiSettings.ResponseFormat = "text"
		services.Scafolder(system_prompt, prompt, &oaiSettings)
	},
}
