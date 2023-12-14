package dayn

import (
	"testing"

	"github.com/jeffrosenberg/advent-of-code-2023/go/pkg/aoc"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	p := NewPart1(aoc.ReadAocInput("../../../inputs/TODO_test.txt"))
	expected := []string{"TODO", "TODO"}
	got := parse(p)
	assert.Equal(t, expected, got)
}

func TestParseLine(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected any
		skip     bool
	}{
		{
			name:     "TODO",
			input:    "TODO",
			expected: "TODO",
		},
	}

	for _, test := range tests {
		t.Run(string(test.name), func(t *testing.T) {
			if test.skip {
				t.Skipf("Skipping %s", string(test.name))
			}

			t.Log(test.name)
			got := parseLine(test.input)
			assert.Equal(t, test.expected, got)
		})
	}
}
