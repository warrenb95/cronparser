package cron

import (
	"strconv"
	"strings"
)

func getOutputString(minVal, maxVal, interval int) string {
	outputSlice := []string{}
	for i := minVal; i <= maxVal; i += interval {
		outputSlice = append(outputSlice, strconv.Itoa(i))
	}

	return strings.Join(outputSlice, " ")
}
