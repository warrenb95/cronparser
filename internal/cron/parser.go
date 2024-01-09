package cron

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const ()

// Parse the inputSlice cron values and return a string array of formatted values.
func Parse(inputSlice []string) ([]string, error) {
	ranges := [][]int{
		{0, 59}, // Minute
		{0, 23}, // Hour
		{1, 31}, // DayMonth
		{1, 12}, // Month
		{1, 7},  // DayWeek
	}

	outputArray := make([]string, len(inputSlice))
	for i, arg := range inputSlice {
		switch {
		case strings.Contains(arg, "/"):
			formattedValue, err := getWithInterval(arg, ranges[i][1])
			if err != nil {
				return nil, fmt.Errorf("parsing cron: %v", err)
			}
			outputArray[i] = formattedValue
		case strings.Contains(arg, "-"):
			min, max := getBounds(arg)
			formattedValue, err := getWithRange(min, max)
			if err != nil {
				return nil, fmt.Errorf("parsing cron: %v", err)
			}
			outputArray[i] = formattedValue
		case strings.Contains(arg, ","):
			argSlice := strings.Split(arg, ",")
			outputArray[i] = strings.Join(argSlice, " ")
		case strings.Contains(arg, "*"):
			formattedValue, err := getWithRange(ranges[i][0], ranges[i][1])
			if err != nil {
				return nil, fmt.Errorf("parsing cron: %v", err)
			}
			outputArray[i] = formattedValue
		default:
			outputArray[i] = arg
		}
	}

	return outputArray, nil
}

func getWithInterval(arg string, maxVal int) (string, error) {
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
			return "", fmt.Errorf("can't parse minVal %v", minStr)
		}
		minVal = v
	}

	interval, err := strconv.Atoi(intervalStr)
	if err != nil {
		return "", fmt.Errorf("parsing interval string while getWithInterval: %v", err)
	}

	return getOutputString(minVal, maxVal, interval), nil
}

func getWithRange(minVal, maxVal int) (string, error) {
	if minVal > maxVal {
		err := fmt.Errorf("invalid range %d-%d - %d must be smaller than %d", minVal, maxVal, minVal, maxVal)
		log.Print(err)
		return "", err
	}
	return getOutputString(minVal, maxVal, int(1)), nil
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
