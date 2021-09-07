package printer

import (
	"fmt"
)

// PrintCron will print the cron to the terminal
// TODO: could extend to take in a writer interface and the write to that
func PrintCron(titles, args []string) {
	for i, arg := range args {
		fmt.Printf("%-14v%v\n", titles[i], arg)
	}
}
