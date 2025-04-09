package main

import (
	"log"

	"github.com/am8850/cliai/cmd"
	"github.com/am8850/cliai/pkg/services"
)

func main() {
	defer services.DisposeClient()

	// Execute the root command
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}

}
