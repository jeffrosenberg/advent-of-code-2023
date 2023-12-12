package day11

import (
	"testing"

	"github.com/jeffrosenberg/advent-of-code-2023/go/pkg/aoc"
	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	p := NewPart1(aoc.ReadAocInput("../../../inputs/11_test.txt"))
	expected := 374
	p.Solve()
	assert.Equal(t, expected, p.Value())
}

func TestSolvePart2(t *testing.T) {
	tests := []struct {
		name            string
		expansionFactor int
		expected        int
		skip            bool
	}{
		{
			name:            "Expand x 10",
			expansionFactor: 10,
			expected:        1030,
		},
		{
			name:            "Expand x 100",
			expansionFactor: 100,
			expected:        8410,
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			p := NewPart2(aoc.ReadAocInput("../../../inputs/11_test.txt"))

			// Can't just use p.Solve(), because the example answer
			// uses only factors of 10 and 100
			points := parseAndExpand(p, test.expansionFactor)
			p.value = solve(points)
			assert.Equal(t, test.expected, p.Value())
		})
	}
}

func TestGetPathLength(t *testing.T) {
	tests := []struct {
		name     string
		from     point
		to       point
		expected int
		skip     bool
	}{
		{
			name:     "Example 1",
			from:     point{x: 1, y: 5},
			to:       point{x: 5, y: 0},
			expected: 9,
		},
		{
			name:     "Example 2",
			from:     point{x: 4, y: 11},
			to:       point{x: 9, y: 1},
			expected: 15,
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			got := getPathLength(test.from, test.to)
			assert.Equal(t, test.expected, got)
		})
	}
}

func TestParseAndExpand(t *testing.T) {
	p := NewPart1(aoc.ReadAocInput("../../../inputs/11_test.txt"))
	// Following the order they're numbered in the example
	expected := []point{
		{x: 4, y: 11},
		{x: 9, y: 10},
		{x: 0, y: 9},
		{x: 8, y: 6},
		{x: 1, y: 5},
		{x: 12, y: 4},
		{x: 9, y: 1},
		{x: 0, y: 0},
		{x: 5, y: 0},
	}
	got := parseAndExpand(p, 2)
	assert.Equal(t, expected, got)
}

func TestParseLine(t *testing.T) {
	tests := []struct {
		name            string
		input           string
		line            int
		expectedPoints  []point
		expectedColumns []int
		skip            bool
	}{
		{
			name:  "Line 0",
			input: "#...#.....",
			line:  0,
			expectedPoints: []point{
				{x: 0, y: 0},
				{x: 4, y: 0},
			},
			expectedColumns: []int{0, 4},
		},
		{
			name:  "Line 1",
			input: ".......#..",
			line:  1,
			expectedPoints: []point{
				{x: 7, y: 1},
			},
			expectedColumns: []int{7},
		},
		{
			name:            "Line 2",
			input:           "..........",
			line:            2,
			expectedPoints:  []point{},
			expectedColumns: []int{},
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			gotPoints, gotColumns := parseLine(test.input, test.line)
			assert.Equal(t, len(test.expectedPoints), len(gotPoints))
			for i, exp := range test.expectedPoints {
				got := gotPoints[i]
				assert.Equal(t, exp, *got)
			}
			assert.Equal(t, test.expectedColumns, gotColumns)
		})
	}
}
