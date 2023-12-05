package dayn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	tests := []struct {
		name     string
		input    string // TODO: modify
		expected bool
		skip     bool
	}{
		{
			name:     "My test name",
			input:    "TODO",
			expected: false,
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			got := supportingFunc(test.input)
			assert.Equal(t, test.expected, got)
		})
	}
}
