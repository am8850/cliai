package main

import (
	"fmt"

	"github.com/am8850/cliai/cmd"
	"github.com/am8850/cliai/pkg/services"
)

func main() {
	// Execute the root command
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}

	// Dispose of the OpenAI client
	services.DisposeClient()
}
