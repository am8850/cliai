package cmd

import (
	"fmt"

	"github.com/am8850/cliai/pkg/config"
	"github.com/am8850/cliai/pkg/refactorer"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var (
	file   string
	output string
)

// Create the version command
var cmdRefactor = &cobra.Command{
	Use:     "refactor",
	Aliases: []string{"re"},
	Short:   "Evaluate code for clarity and complexity",
	Run: func(cmd *cobra.Command, args []string) {

		if file == "" {
			fmt.Println("Please provide a command. Example:")
			color.Cyan.Println("cliai re -f app.py -o app_sanitized.py")
			return
		}

		template := config.FindTemplate("refactor")
		system_prompt := template.SystemPrompt

		refactorer.Refactor(system_prompt, file, output)
	},
}

func init() {
	cmdRefactor.PersistentFlags().StringVarP(&file, "file", "f", "", "The file path to sanitize [required]")
	cmdRefactor.PersistentFlags().StringVarP(&output, "output", "o", "", "The output file path and name")
}
