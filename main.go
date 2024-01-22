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

	// strArgs := strings.Fields(inputArgs[0])
	// if len(strArgs) != 8 {
	// 	log.Fatalf("input args len invalid: %v", len(strArgs))
	// }

	inputSlice := strings.Split(inputArgs[0], " ")

	fmt.Print(inputSlice[:5])

	processedArgs, err := cron.Parse(inputSlice[:5])
	if err != nil {
		log.Fatalf("failed to parse cron input arg: %v", err)
	}

	processedArgs = append(processedArgs, strings.Join(inputSlice[5:], " "))

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
