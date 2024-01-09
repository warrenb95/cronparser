package cron

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWithRange(t *testing.T) {
	t.Parallel()

	// Test Cases
	tests := map[string]struct {
		min int
		max int

		expectedOutput string
		errContains    string
	}{
		"pass": {
			min:            1,
			max:            5,
			expectedOutput: "1 2 3 4 5",
		},
	}

	// Testing
	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result, err := getWithRange(test.min, test.max)
			if test.errContains != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), test.errContains)
				return
			}
			require.NoError(t, err)

			assert.Equal(t, test.expectedOutput, result)
		})
	}
}

func TestWithInverval(t *testing.T) {
	t.Parallel()

	// Test Cases
	tests := map[string]struct {
		input  string
		maxVal int

		expectedOutput string
		errContains    string
	}{
		"every 10": {
			input:          "*/10",
			maxVal:         59,
			expectedOutput: "0 10 20 30 40 50",
		},

		"every 1 past 10": {
			input:          "1/10",
			maxVal:         30,
			expectedOutput: "1 11 21",
		},

		"every 5 between 5 and 40, ignoring max val": {
			input:          "5-40/10",
			maxVal:         60,
			expectedOutput: "5 15 25 35",
		},
	}

	// Testing
	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result, err := getWithInterval(test.input, test.maxVal)
			if test.errContains != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), test.errContains)
				return
			}
			require.NoError(t, err)

			assert.Equal(t, test.expectedOutput, result)
		})
	}
}
