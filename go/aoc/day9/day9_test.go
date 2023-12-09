package day9

import (
	"testing"

	"github.com/jeffrosenberg/advent-of-code-2023/go/pkg/aoc"
	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	expected := 114
	p := NewPart1(aoc.ReadAocInput("../../../inputs/9_test.txt"))
	p.Solve()
	assert.Equal(t, expected, p.Value())
}

func TestExtrapolate(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
		skip     bool
	}{
		{
			name:     "Example line 1",
			input:    []int{0, 3, 6, 9, 12, 15},
			expected: 18,
		},
		{
			name:     "Example line 2",
			input:    []int{1, 3, 6, 10, 15, 21},
			expected: 28,
		},
		{
			name:     "Example line 3",
			input:    []int{10, 13, 16, 21, 30, 45},
			expected: 68,
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			got := extrapolate(test.input)
			assert.Equal(t, test.expected, got)
		})
	}
}

func TestParse(t *testing.T) {
	p := NewPart1(aoc.ReadAocInput("../../../inputs/9_test.txt"))
	expected := [][]int{
		{0, 3, 6, 9, 12, 15},
		{1, 3, 6, 10, 15, 21},
		{10, 13, 16, 21, 30, 45},
	}
	got := parse(p)
	assert.Equal(t, expected, got)
}

func TestParseLine(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []int
		skip     bool
	}{
		{
			name:     "Example line 1",
			input:    "0 3 6 9 12 15",
			expected: []int{0, 3, 6, 9, 12, 15},
		},
		{
			name:     "Example line 2",
			input:    "1 3 6 10 15 21",
			expected: []int{1, 3, 6, 10, 15, 21},
		},
		{
			name:     "Example line 3",
			input:    "10 13 16 21 30 45",
			expected: []int{10, 13, 16, 21, 30, 45},
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			got := parseLine(test.input)
			assert.Equal(t, test.expected, got)
		})
	}
}
