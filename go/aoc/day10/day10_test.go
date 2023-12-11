package day10

import (
	"testing"

	"github.com/jeffrosenberg/advent-of-code-2023/go/pkg/aoc"
	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name     string
		file     string
		expected int
		skip     bool
	}{
		{
			name:     "Example 1",
			file:     "../../../inputs/10_test1.txt",
			expected: 4,
		},
		{
			name:     "Example 2",
			file:     "../../../inputs/10_test2.txt",
			expected: 8,
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			p := NewPart1(aoc.ReadAocInput(test.file))
			p.Solve()
			assert.Equal(t, test.expected, p.Value())
		})
	}
}

func TestCalculateLoop(t *testing.T) {
	tests := []struct {
		name          string
		file          string
		expectedPipes []pipe
		skip          bool
	}{
		{
			name: "Example 1",
			file: "../../../inputs/10_test1.txt",
			expectedPipes: []pipe{
				{x: 1, y: 3, orientation: 'S'},
				{x: 2, y: 3, orientation: '-'},
				{x: 3, y: 3, orientation: '7'},
				{x: 3, y: 2, orientation: '|'},
				{x: 3, y: 1, orientation: 'J'},
				{x: 2, y: 1, orientation: '-'},
				{x: 1, y: 1, orientation: 'L'},
				{x: 1, y: 2, orientation: '|'},
				{x: 1, y: 3, orientation: 'S'},
			},
		},
		{
			name: "Example 2",
			file: "../../../inputs/10_test2.txt",
			expectedPipes: []pipe{
				{x: 0, y: 2, orientation: 'S'},
				{x: 1, y: 2, orientation: 'J'},
				{x: 1, y: 3, orientation: 'F'},
				{x: 2, y: 3, orientation: 'J'},
				{x: 2, y: 4, orientation: 'F'},
				{x: 3, y: 4, orientation: '7'},
				{x: 3, y: 3, orientation: '|'},
				{x: 3, y: 2, orientation: 'L'},
				{x: 4, y: 2, orientation: '7'},
				{x: 4, y: 1, orientation: 'J'},
				{x: 3, y: 1, orientation: '-'},
				{x: 2, y: 1, orientation: '-'},
				{x: 1, y: 1, orientation: 'F'},
				{x: 1, y: 0, orientation: 'J'},
				{x: 0, y: 0, orientation: 'L'},
				{x: 0, y: 1, orientation: '|'},
				{x: 0, y: 2, orientation: 'S'},
			},
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			p := NewPart1(aoc.ReadAocInput(test.file))
			p.Parse()

			got, _ := calculateLoop(p, []pipe{p.getStart()}, "")
			assert.Equal(t, test.expectedPipes, got)
		})
	}
}

func TestGetNext(t *testing.T) {
	tests := []struct {
		name     string
		input    pipe
		from     string
		expected []string
		skip     bool
	}{
		{
			name:     "North (|)",
			input:    pipe{x: 3, y: 3, orientation: '|'},
			from:     "3.2",
			expected: []string{"3.4"},
		},
		{
			name:     "East (L)",
			input:    pipe{x: 3, y: 3, orientation: 'L'},
			from:     "3.4",
			expected: []string{"4.3"},
		},
		{
			name:     "North (L)",
			input:    pipe{x: 3, y: 3, orientation: 'L'},
			from:     "4.3",
			expected: []string{"3.4"},
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			got := test.input.getNext(test.from)
			assert.Equal(t, test.expected, got)
		})
	}
}

func TestParse(t *testing.T) {
	p := NewPart1(aoc.ReadAocInput("../../../inputs/10_test1.txt"))
	expectedMap := map[string]pipe{
		"0.0": {x: 0, y: 0, orientation: 'L'},
		"1.0": {x: 1, y: 0, orientation: '|'},
		"2.0": {x: 2, y: 0, orientation: '-'},
		"3.0": {x: 3, y: 0, orientation: 'J'},
		"4.0": {x: 4, y: 0, orientation: 'F'},
		"0.1": {x: 0, y: 1, orientation: '-'},
		"1.1": {x: 1, y: 1, orientation: 'L'},
		"2.1": {x: 2, y: 1, orientation: '-'},
		"3.1": {x: 3, y: 1, orientation: 'J'},
		"4.1": {x: 4, y: 1, orientation: '|'},
		"0.2": {x: 0, y: 2, orientation: 'L'},
		"1.2": {x: 1, y: 2, orientation: '|'},
		"2.2": {x: 2, y: 2, orientation: '7'},
		"3.2": {x: 3, y: 2, orientation: '|'},
		"4.2": {x: 4, y: 2, orientation: '|'},
		"0.3": {x: 0, y: 3, orientation: '7'},
		"1.3": {x: 1, y: 3, orientation: 'S'},
		"2.3": {x: 2, y: 3, orientation: '-'},
		"3.3": {x: 3, y: 3, orientation: '7'},
		"4.3": {x: 4, y: 3, orientation: '|'},
		"0.4": {x: 0, y: 4, orientation: '-'},
		"1.4": {x: 1, y: 4, orientation: 'L'},
		"2.4": {x: 2, y: 4, orientation: '|'},
		"3.4": {x: 3, y: 4, orientation: 'F'},
		"4.4": {x: 4, y: 4, orientation: '7'},
	}
	expectedStart := &pipe{x: 1, y: 3, orientation: 'S'}
	gotMap, gotStart := parse(p)
	assert.Equal(t, expectedMap, gotMap)
	assert.Equal(t, expectedStart, gotStart)

	p.Parse()
	assert.Equal(t, expectedMap, p.getPipes())
	assert.Equal(t, *expectedStart, p.getStart())
}

func TestParseLine(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		line     int
		expected []pipe
		skip     bool
	}{
		{
			name:  "Example 1 line 1", // Line 1 = y-coordinate of 4 for the example
			input: "-L|F7",
			line:  4,
			expected: []pipe{
				{
					x:           0,
					y:           4,
					orientation: '-',
				},
				{
					x:           1,
					y:           4,
					orientation: 'L',
				},
				{
					x:           2,
					y:           4,
					orientation: '|',
				},
				{
					x:           3,
					y:           4,
					orientation: 'F',
				},
				{
					x:           4,
					y:           4,
					orientation: '7',
				},
			},
		},
		{
			name:  "Example 2 line 2", // Line 2 = y-coordinate of 4 for the example
			input: ".FJ|7",
			line:  4,
			expected: []pipe{
				{
					x:           1,
					y:           4,
					orientation: 'F',
				},
				{
					x:           2,
					y:           4,
					orientation: 'J',
				},
				{
					x:           3,
					y:           4,
					orientation: '|',
				},
				{
					x:           4,
					y:           4,
					orientation: '7',
				},
			},
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			got, _ := parseLine(test.input, test.line)
			assert.Equal(t, test.expected, got)
		})
	}
}
