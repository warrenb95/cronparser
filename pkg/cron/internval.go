package cron

import (
	"log"
	"strconv"
	"strings"
)

// getWithInterval will get values over an interval
// e.g arg = */10, maxVal = 60 -- return = 0 10 20 30 40 60
func getWithInterval(arg string, maxVal int) string {
	splitStr := strings.Split(arg, "/")
	minStr := splitStr[0]
	intervalStr := splitStr[1]

	// calc the min value, checking if it is a range
	var minVal int
	if strings.Contains(minStr, "-") {
		minVal, maxVal = getBounds(minStr)
	} else if minStr != "*" {
		v, err := strconv.Atoi(minStr)
		if err != nil {
			log.Fatalf("can't parse minVal %v", minStr)
		}
		minVal = v
	}

	interval, err := strconv.Atoi(intervalStr)
	if err != nil {
		log.Fatalf("cannot parse input %v error %v", intervalStr, err)
	}

	return getOutputString(minVal, maxVal, interval)
}
