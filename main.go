package main

import (
	"github/warrenb95/cronparser/pkg/cron"
	"github/warrenb95/cronparser/pkg/printer"
	"log"
	"os"
	"strings"
)

func main() {
	// get input
	inputArgs := os.Args[1:2]

	// validate input
	strArgs := strings.Fields(inputArgs[0])
	if len(strArgs) != 6 {
		log.Fatalf("input args len invalid: %v", len(strArgs))
	}

	//process the input
	// do not send the last strArg
	processedArgs := cron.Parse(strArgs[:len(strArgs)-1])
	processedArgs = append(processedArgs, strArgs[len(strArgs)-1])

	// print the value to console
	printer.PrintCron([]string{
		"minute",
		"hour",
		"day of month",
		"month",
		"day of week",
		"command",
	}, processedArgs)
}
