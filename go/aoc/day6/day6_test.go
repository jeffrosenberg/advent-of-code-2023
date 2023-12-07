package day6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuadratic(t *testing.T) {
	tests := []struct {
		name string
		// a is always 1!
		b           int
		c           int
		expectedMin int
		expectedMax int
		skip        bool
	}{
		{
			name:        "First example test",
			b:           -7,
			c:           9,
			expectedMin: 2,
			expectedMax: 5,
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			gotMin, gotMax := quadratic(1, test.b, test.c)
			assert.Equal(t, test.expectedMin, gotMin)
			assert.Equal(t, test.expectedMax, gotMax)
		})
	}
}
