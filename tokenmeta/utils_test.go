package tokenmeta

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_stringInFilter(t *testing.T) {
	tests := []struct {
		name        string
		term        string
		filter      []string
		expectMatch bool
	}{
		{
			name:        "without any filters",
			term:        "zswhq.token",
			expectMatch: true,
		},
		{
			name:        "with a non-matching filter",
			term:        "zswhq.token",
			filter:      []string{"baababbaba"},
			expectMatch: false,
		},
		{
			name:        "with a non-matching filter",
			term:        "zswhq.token",
			filter:      []string{"baababbaba", "zswhq.token"},
			expectMatch: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectMatch, stringInFilter(test.term, test.filter))
		})
	}
}
