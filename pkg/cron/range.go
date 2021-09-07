package cron

import (
	"log"
	"strconv"
	"strings"
)

// getWithRange will get the range of the input
func getWithRange(minVal, maxVal int) string {

	// validate min < max
	if minVal > maxVal {
		log.Fatalf("invalid range %v-%v - %v must be smaller than %v", minVal, maxVal, minVal, maxVal)
	}
	return getOutputString(minVal, maxVal, int(1))
}

func getBounds(arg string) (int, int) {
	splitstring := strings.Split(arg, "-")
	lowstr := splitstring[0]
	highstr := splitstring[1]

	low, err := strconv.Atoi(lowstr)
	if err != nil {
		log.Fatalf("can't parse lowstr %v", err)
	}

	high, err := strconv.Atoi(highstr)
	if err != nil {
		log.Fatalf("can't parse highstr %v", err)
	}

	return low, high
}
