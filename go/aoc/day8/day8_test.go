package day8

import (
	"testing"

	"github.com/jeffrosenberg/advent-of-code-2023/go/pkg/aoc"
	"github.com/stretchr/testify/assert"
)

func TestTraverse(t *testing.T) {
	tests := []struct {
		name         string
		nodes        map[string]node
		instructions string
		start        string
		expected     int
		skip         bool
	}{
		{
			name: "RL example",
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
			expected:     2,
		},
		{
			name: "LLR example",
			nodes: map[string]node{
				"AAA": {left: "BBB", right: "BBB"},
				"BBB": {left: "AAA", right: "ZZZ"},
				"ZZZ": {left: "ZZZ", right: "ZZZ"},
			},
			instructions: "LLR",
			expected:     6,
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			got := traverse(test.nodes, test.instructions)
			assert.Equal(t, test.expected, got)
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
	i, n := parse(p)
	assert.Equal(t, i, expectedInstructions)
	assert.Equal(t, n, expectedNodes)
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
