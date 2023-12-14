package day12

import (
	"testing"

	"github.com/jeffrosenberg/advent-of-code-2023/go/pkg/aoc"
	"github.com/stretchr/testify/assert"
)

func TestCalculateArrangements(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
		skip     bool
	}{
		{
			name:     "Example line 1",
			input:    "???.### 1,1,3",
			expected: 1,
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			got := calculateArrangements(parseLine(test.input))
			assert.Equal(t, test.expected, got)
		})
	}
}

func TestArrangementMeetsConditions(t *testing.T) {
	tests := []struct {
		name     string
		groups   []group
		wanted   []int
		expected bool
		skip     bool
	}{
		{
			name: "Simple match",
			groups: []group{
				{position: 0, length: 3, condition: CONDITION_DAMAGED},
				{position: 5, length: 1, condition: CONDITION_DAMAGED},
				{position: 7, length: 2, condition: CONDITION_DAMAGED},
			},
			wanted:   []int{3, 1, 2},
			expected: true,
		},
		{
			name: "Wrong number of groups",
			groups: []group{
				{position: 0, length: 3, condition: CONDITION_DAMAGED},
				{position: 5, length: 1, condition: CONDITION_DAMAGED},
				{position: 7, length: 2, condition: CONDITION_DAMAGED},
				{position: 10, length: 2, condition: CONDITION_DAMAGED},
			},
			wanted:   []int{3, 1, 2},
			expected: true,
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			got := arrangementMeetsConditions(test.groups, test.wanted)
			assert.Equal(t, test.expected, got)
		})
	}
}

func TestParse(t *testing.T) {
	p := NewPart1(aoc.ReadAocInput("../../../inputs/12_test.txt"))
	expectedInputs := []string{
		"???.###",
		".??..??...?##.",
		"?#?#?#?#?#?#?#?",
		"????.#...#...",
		"????.######..#####.",
		"?###????????",
	}
	expectedGroups := [][]group{
		{
			{position: 0, length: 3, condition: CONDITION_UNKNOWN},
			{position: 3, length: 1, condition: CONDITION_OPERATIONAL},
			{position: 4, length: 3, condition: CONDITION_DAMAGED},
		},
		{
			{position: 0, length: 1, condition: CONDITION_OPERATIONAL},
			{position: 1, length: 2, condition: CONDITION_UNKNOWN},
			{position: 3, length: 2, condition: CONDITION_OPERATIONAL},
			{position: 5, length: 2, condition: CONDITION_UNKNOWN},
			{position: 7, length: 3, condition: CONDITION_OPERATIONAL},
			{position: 10, length: 1, condition: CONDITION_UNKNOWN},
			{position: 11, length: 2, condition: CONDITION_DAMAGED},
			{position: 13, length: 1, condition: CONDITION_OPERATIONAL},
		},
		{
			{position: 0, length: 1, condition: CONDITION_UNKNOWN},
			{position: 1, length: 1, condition: CONDITION_DAMAGED},
			{position: 2, length: 1, condition: CONDITION_UNKNOWN},
			{position: 3, length: 1, condition: CONDITION_DAMAGED},
			{position: 4, length: 1, condition: CONDITION_UNKNOWN},
			{position: 5, length: 1, condition: CONDITION_DAMAGED},
			{position: 6, length: 1, condition: CONDITION_UNKNOWN},
			{position: 7, length: 1, condition: CONDITION_DAMAGED},
			{position: 8, length: 1, condition: CONDITION_UNKNOWN},
			{position: 9, length: 1, condition: CONDITION_DAMAGED},
			{position: 10, length: 1, condition: CONDITION_UNKNOWN},
			{position: 11, length: 1, condition: CONDITION_DAMAGED},
			{position: 12, length: 1, condition: CONDITION_UNKNOWN},
			{position: 13, length: 1, condition: CONDITION_DAMAGED},
			{position: 14, length: 1, condition: CONDITION_UNKNOWN},
		},
		{
			{position: 0, length: 4, condition: CONDITION_UNKNOWN},
			{position: 4, length: 1, condition: CONDITION_OPERATIONAL},
			{position: 5, length: 1, condition: CONDITION_DAMAGED},
			{position: 6, length: 3, condition: CONDITION_OPERATIONAL},
			{position: 9, length: 1, condition: CONDITION_DAMAGED},
			{position: 10, length: 3, condition: CONDITION_OPERATIONAL},
		},
		{
			{position: 0, length: 4, condition: CONDITION_UNKNOWN},
			{position: 4, length: 1, condition: CONDITION_OPERATIONAL},
			{position: 5, length: 6, condition: CONDITION_DAMAGED},
			{position: 11, length: 2, condition: CONDITION_OPERATIONAL},
			{position: 13, length: 5, condition: CONDITION_DAMAGED},
			{position: 18, length: 1, condition: CONDITION_OPERATIONAL},
		},
		{
			{position: 0, length: 1, condition: CONDITION_UNKNOWN},
			{position: 1, length: 3, condition: CONDITION_DAMAGED},
			{position: 4, length: 8, condition: CONDITION_UNKNOWN},
		},
	}
	expectedWanted := [][]int{
		{1, 1, 3},
		{1, 1, 3},
		{1, 3, 1, 6},
		{4, 1, 1},
		{1, 6, 5},
		{3, 2, 1},
	}

	gotInputs, gotGroups, gotWanted := parse(p)
	assert.Equal(t, expectedInputs, gotInputs)
	assert.Equal(t, expectedGroups, gotGroups)
	assert.Equal(t, expectedWanted, gotWanted)
}

func TestParseLine(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedGroup  []group
		expectedWanted []int
		skip           bool
	}{
		{
			name:  "Example line 1",
			input: "???.### 1,1,3",
			expectedGroup: []group{
				{
					position:  0,
					length:    3,
					condition: CONDITION_UNKNOWN,
				},
				{
					position:  3,
					length:    1,
					condition: CONDITION_OPERATIONAL,
				},
				{
					position:  4,
					length:    3,
					condition: CONDITION_DAMAGED,
				},
			},
			expectedWanted: []int{1, 1, 3},
		},
		{
			name:  "Example line 3",
			input: ".#.###.#.###### 1,3,1,6",
			expectedGroup: []group{
				{
					position:  0,
					length:    1,
					condition: CONDITION_OPERATIONAL,
				},
				{
					position:  1,
					length:    1,
					condition: CONDITION_DAMAGED,
				},
				{
					position:  2,
					length:    1,
					condition: CONDITION_OPERATIONAL,
				},
				{
					position:  3,
					length:    3,
					condition: CONDITION_DAMAGED,
				},
				{
					position:  6,
					length:    1,
					condition: CONDITION_OPERATIONAL,
				},
				{
					position:  7,
					length:    1,
					condition: CONDITION_DAMAGED,
				},
				{
					position:  8,
					length:    1,
					condition: CONDITION_OPERATIONAL,
				},
				{
					position:  9,
					length:    6,
					condition: CONDITION_DAMAGED,
				},
			},
			expectedWanted: []int{1, 3, 1, 6},
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			_, gotGroup, gotWanted := parseLine(test.input)
			assert.Equal(t, test.expectedGroup, gotGroup)
			assert.Equal(t, test.expectedWanted, gotWanted)
		})
	}
}
