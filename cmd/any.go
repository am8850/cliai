package cmd

import (
	"fmt"

	"github.com/am8850/cliai/pkg/processor"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var (
	system_prompt   string
	response_format string = "text"
)

// Add the root command to the application
var cmdAny = &cobra.Command{
	Use:   "any",
	Short: "Any general purpose prompt",
	Run: func(cmd *cobra.Command, args []string) {
		if prompt == "" {
			fmt.Println("Please provide a command. Example:")
			color.Cyan.Println("cliai any -p 'What is the speed of light?'")
			return
		}
		processor.Any(system_prompt, prompt, response_format)
	},
}

func init() {
	cmdAny.PersistentFlags().StringVarP(&system_prompt, "system", "s", "You are a general purpose AI assistant.", "Override the system prompt.")
	cmdAny.PersistentFlags().StringVarP(&response_format, "format", "f", "text", "Response format: text or json_object. Default is text.")
}
