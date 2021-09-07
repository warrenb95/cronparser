package cron

import (
	"strings"
)

const ()

func Parse(inputSlice []string) []string {
	ranges := [][]int{
		{0, 59}, // Minute
		{0, 23}, // Hour
		{1, 31}, // DayMonth
		{1, 12}, // Month
		{1, 7},  // DayWeek
	}

	for i, arg := range inputSlice {
		switch {
		case strings.Contains(arg, "/"):
			inputSlice[i] = getWithInterval(arg, ranges[i][1])
		case strings.Contains(arg, "-"):
			min, max := getBounds(arg)
			inputSlice[i] = getWithRange(min, max)
		case strings.Contains(arg, ","):
			argSlice := strings.Split(arg, ",")
			inputSlice[i] = strings.Join(argSlice, " ")
		case strings.Contains(arg, "*"):
			inputSlice[i] = getWithRange(ranges[i][0], ranges[i][1])
		default:
			inputSlice[i] = arg
		}
	}

	return inputSlice
}
