package main

import (
	"log"

	"github.com/am8850/cliai/cmd"
	"github.com/am8850/cliai/pkg/config"
	"github.com/am8850/cliai/pkg/openai"
)

func main() {
	defer openai.DisposeClient()

	config.GetConfig()

	// Execute the root command
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}

}
