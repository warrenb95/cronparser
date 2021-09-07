package cron

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithInterval(t *testing.T) {
	type ttest struct {
		input  string
		output string
		maxVal int
	}

	tests := []ttest{
		{
			input:  "*/10",
			output: "0 10 20 30 40 50",
			maxVal: 59,
		},
		{
			input:  "1/10",
			output: "1 11 21",
			maxVal: 30,
		},
		{
			input:  "5-40/10",
			output: "5 15 25 35",
			maxVal: 60,
		},
	}

	for _, tt := range tests {
		result := getWithInterval(tt.input, tt.maxVal)
		assert.Equal(t, tt.output, result)
	}
}

func TestWithRange(t *testing.T) {
	type ttest struct {
		min    int
		max    int
		output string
	}

	tests := []ttest{
		{
			min:    1,
			max:    5,
			output: "1 2 3 4 5",
		},
	}

	for _, tt := range tests {
		result := getWithRange(tt.min, tt.max)
		assert.Equal(t, tt.output, result)
	}
}

func TestParse(t *testing.T) {
	type ttest struct {
		input  []string
		output []string
	}

	tests := []ttest{
		{
			input: []string{
				"*/15",
				"0",
				"1,15",
				"*",
				"1-5",
			},
			output: []string{
				"0 15 30 45",
				"0",
				"1 15",
				"1 2 3 4 5 6 7 8 9 10 11 12",
				"1 2 3 4 5",
			},
		},
	}

	for _, tt := range tests {
		res := Parse(tt.input)
		assert.Equal(t, tt.output, res)
	}
}
