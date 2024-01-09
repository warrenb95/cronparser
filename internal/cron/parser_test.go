package cron_test

import (
	"github/warrenb95/cronparser/internal/cron"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input  []string
		output []string
	}{
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
		res, err := cron.Parse(tt.input)
		require.NoError(t, err)

		assert.Equal(t, tt.output, res)
	}
}
