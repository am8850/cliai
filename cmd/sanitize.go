package cmd

import (
	"fmt"

	"github.com/am8850/cliai/pkg/services"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var (
	file   string
	output string
)

// Create the version command
var cmdSanitize = &cobra.Command{
	Use:     "sanitize",
	Aliases: []string{"sa"},
	Short:   "Sanitize code",
	Run: func(cmd *cobra.Command, args []string) {

		if file == "" {
			fmt.Println("Please provide a command. Example:")
			color.Cyan.Println("cliai sa -f app.py -o app_sanitized.py")
			return
		}

		system_prompt := `system:

You are an AI that can evaluate the programming code for readability and cyclomatic complexity. 

Rules:
- Code can be in any programming language.
- Do your best to provide a score for readability and cyclomatic complexity.
- Provide a score from 1 to 10 for each category. 
- Provide reasons for the scores. 
- Generate version of the code that includes the proposed changes to improve readability and cyclomatic compexity. Do your best to provde the best possible version of the code. Add missing comments to the functions.
- The code should be in ISO-8859-1 encoding.
- No prologue or epilogue.
- Output in the following JSON format only: 
{
"readability_score":0,
"readability_reason":"",
"cyclomatic_score":0,
"cyclomatic_reason":"",
"improved_code":"import os\nmsg=\"Hello World\"\nprint(msg)",
}
`
		services.Sanitizer(system_prompt, file, output, &oaiSettings)
	},
}

func init() {
	cmdSanitize.PersistentFlags().StringVarP(&file, "file", "f", "", "The file path to sanitize [required]")
	cmdSanitize.PersistentFlags().StringVarP(&output, "output", "o", "", "The output file path and name")
}
