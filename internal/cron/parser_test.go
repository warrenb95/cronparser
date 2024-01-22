package cron_test

import (
	"github/warrenb95/cronparser/internal/cron"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	t.Parallel()

	// Test Cases
	tests := map[string]struct {
		input          []string
		expectedOutput []string
		errContains    string
	}{
		"pass": {
			input: []string{
				"*/15",
				"0",
				"1,15",
				"*",
				"1-5",
			},
			expectedOutput: []string{
				"0 15 30 45",
				"0",
				"1 15",
				"1 2 3 4 5 6 7 8 9 10 11 12",
				"1 2 3 4 5",
			},
		},
		"pass circular range": {
			input: []string{
				"*/15",
				"0",
				"30-4",
				"*",
				"1-5",
			},
			expectedOutput: []string{
				"0 15 30 45",
				"0",
				"30 31 1 2 3 4",
				"1 2 3 4 5 6 7 8 9 10 11 12",
				"1 2 3 4 5",
			},
		},
		"error parsing min value": {
			input: []string{
				"oops/15",
				"0",
				"1,15",
				"*",
				"1-5",
			},
			errContains: "can't parse minVal",
		},
	}

	// Testing
	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			res, err := cron.Parse(test.input)

			if test.errContains != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), test.errContains)
				return
			}
			require.NoError(t, err)

			assert.Equal(t, test.expectedOutput, res)
		})
	}
}
