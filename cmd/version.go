package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Create the version command
var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "Print the application version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CLIAI version 1.0")
	},
}
