package main

import (
	"fmt"
	"github/warrenb95/cronparser/internal/cron"
	"log"
	"os"
	"strings"
)

func main() {
	inputArgs := os.Args[1:2]

	strArgs := strings.Fields(inputArgs[0])
	if len(strArgs) != 6 {
		log.Fatalf("input args len invalid: %v", len(strArgs))
	}

	processedArgs, err := cron.Parse(strArgs[:len(strArgs)-1])
	if err != nil {
		log.Fatalf("failed to parse cron input arg: %v", err)
	}

	processedArgs = append(processedArgs, strArgs[len(strArgs)-1])

	for i, title := range []string{
		"minute",
		"hour",
		"day of month",
		"month",
		"day of week",
		"command",
	} {
		fmt.Printf("%-14s%s\n", title, processedArgs[i])
	}
}
