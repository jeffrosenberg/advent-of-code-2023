package day8

import (
	"testing"

	"github.com/jeffrosenberg/advent-of-code-2023/go/pkg/aoc"
	"github.com/stretchr/testify/assert"
)

// Doesn't actually test anything,
// this can be used to debug the solution run.
// func TestDebug(t *testing.T) {
// 	p := NewPart2(aoc.ReadAocInput("../../../inputs/8.txt"))
// 	p.Solve()
// }

func TestTraversePart1(t *testing.T) {
	tests := []struct {
		name         string
		p            *Part1
		nodes        map[string]node
		instructions string
		start        string
		expected     int
		skip         bool
	}{
		{
			name: "RL example",
			p: &Part1{
				nodes: map[string]node{
					"AAA": {left: "BBB", right: "CCC"},
					"BBB": {left: "DDD", right: "EEE"},
					"CCC": {left: "ZZZ", right: "GGG"},
					"DDD": {left: "DDD", right: "DDD"},
					"EEE": {left: "EEE", right: "EEE"},
					"GGG": {left: "GGG", right: "GGG"},
					"ZZZ": {left: "ZZZ", right: "ZZZ"},
				},
				instructions: "RL",
				startNodes:   []string{"AAA"},
			},
			expected: 2,
		},
		{
			name: "LLR example",
			p: &Part1{
				nodes: map[string]node{
					"AAA": {left: "BBB", right: "BBB"},
					"BBB": {left: "AAA", right: "ZZZ"},
					"ZZZ": {left: "ZZZ", right: "ZZZ"},
				},
				instructions: "LLR",
				startNodes:   []string{"AAA"},
			},
			expected: 6,
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			got := traverse(test.p)
			assert.Equal(t, test.expected, got)
		})
	}
}

func TestTraversePart2(t *testing.T) {
	tests := []struct {
		name         string
		p            *Part2
		nodes        map[string]node
		instructions string
		start        string
		expected     int
		skip         bool
	}{
		{
			name: "LR example",
			p: &Part2{
				nodes: map[string]node{
					"11A": {left: "11B", right: "XXX"},
					"11B": {left: "XXX", right: "11Z"},
					"11Z": {left: "11B", right: "XXX"},
					"22A": {left: "22B", right: "XXX"},
					"22B": {left: "22C", right: "22C"},
					"22C": {left: "22Z", right: "22Z"},
					"22Z": {left: "22B", right: "22B"},
					"XXX": {left: "XXX", right: "XXX"},
				},
				instructions: "LR",
				startNodes:   []string{"11A", "22A"},
			},
			expected: 6,
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			got := traverse(test.p)
			assert.Equal(t, test.expected, got)
		})
	}
}

func TestFindMatchingEndNodes(t *testing.T) {
	tests := []struct {
		name         string
		input        [][]int
		expectedInt  int
		expectedBool bool
		skip         bool
	}{
		{
			name: "No match",
			input: [][]int{
				{5, 10, 15},
				{22, 44},
			},
			expectedInt:  0,
			expectedBool: false,
		},
		{
			name: "Two inputs",
			input: [][]int{
				{5, 10, 15, 20},
				{4, 8, 12, 20},
			},
			expectedInt:  20,
			expectedBool: true,
		},
		{
			name: "Four inputs",
			input: [][]int{
				{5, 10, 15, 20, 25, 30, 35, 40},
				{10, 20, 30, 40, 50, 60},
				{20, 40, 60},
				{8, 16, 24, 32, 40, 48, 56},
			},
			expectedInt:  40,
			expectedBool: true,
		},
	}

	var matches chan (int) = make(chan int)
	var answer chan (int) = make(chan int, 1)
	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)

			go findMatchingEndNodes(matches, test.expectedInt, answer)
			for val := range test.input {
				matches <- val
			}
			got := <-answer
			assert.Equal(t, test.expectedInt, got)
		})
	}
}

func TestParse(t *testing.T) {
	p := NewPart1(aoc.ReadAocInput("../../../inputs/8_test.txt"))
	expectedInstructions := "RL"
	expectedNodes := map[string]node{
		"AAA": {left: "BBB", right: "CCC"},
		"BBB": {left: "DDD", right: "EEE"},
		"CCC": {left: "ZZZ", right: "GGG"},
		"DDD": {left: "DDD", right: "DDD"},
		"EEE": {left: "EEE", right: "EEE"},
		"GGG": {left: "GGG", right: "GGG"},
		"ZZZ": {left: "ZZZ", right: "ZZZ"},
	}
	expectedStart := []string{"AAA"}
	i, n, s := parse(p)
	assert.Equal(t, i, expectedInstructions)
	assert.Equal(t, n, expectedNodes)
	assert.Equal(t, s, expectedStart)
}

func TestParseLine(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		expectedKey  string
		expectedNode node
		skip         bool
	}{
		{
			name:         "AAA",
			input:        "AAA = (BBB, CCC)",
			expectedKey:  "AAA",
			expectedNode: node{left: "BBB", right: "CCC"},
		},
		{
			name:         "BBB",
			input:        "BBB = (DDD, EEE)",
			expectedKey:  "BBB",
			expectedNode: node{left: "DDD", right: "EEE"},
		},
		{
			name:         "Random letters",
			input:        "LOL = (WTF, OMG)",
			expectedKey:  "LOL",
			expectedNode: node{left: "WTF", right: "OMG"},
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			gotKey, gotNode := parseLine(test.input)
			assert.Equal(t, test.expectedKey, gotKey)
			assert.Equal(t, test.expectedNode, gotNode)
		})
	}
}
