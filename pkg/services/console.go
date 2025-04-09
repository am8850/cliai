package services

import (
	"fmt"

	"github.com/gookit/color"
)

func askForConfirmation(s string) bool {
	var response string
	color.Yellow.Printf("%s (y/n): ", s)
	fmt.Scanln(&response)
	return response == "y" || response == "Y"
}
