package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Calculation(t *testing.T) {
	tests := []struct {
		name     string
		input    NumberToken
		line     int
		lines    []string
		expected bool
		skip     bool
	}{
		{
			name:  "No adjacent symbols",
			input: NumberToken{143, 1, 3},
			line:  1,
			lines: []string{
				".....",
				".143.",
				".....",
				".....",
			},
			expected: false,
		},
		{
			name:  "Symbol above (1)",
			input: NumberToken{143, 1, 3},
			line:  1,
			lines: []string{
				".*...",
				".143.",
				".....",
				".....",
			},
			expected: true,
		},
		{
			name:  "Symbol above (2)",
			input: NumberToken{143, 0, 2},
			line:  1,
			lines: []string{
				".&...",
				"143..",
				".....",
				".....",
			},
			expected: true,
		},
		{
			name:  "Symbol above (EOF)",
			input: NumberToken{143, 1, 3},
			line:  3,
			lines: []string{
				".....",
				".....",
				"..-..",
				".143.",
			},
			expected: true,
		},
		{
			name:  "Symbol below",
			input: NumberToken{143, 1, 3},
			line:  0,
			lines: []string{
				".143.",
				"..&..",
				".....",
				".....",
			},
			expected: true,
		},
		{
			name:  "Symbol next to",
			input: NumberToken{143, 1, 3},
			line:  1,
			lines: []string{
				".....",
				".143#",
				".....",
				".....",
			},
			expected: true,
		},
		{
			name:  "Symbol next to (EOL)",
			input: NumberToken{143, 2, 4},
			line:  1,
			lines: []string{
				".....",
				".+143",
				".....",
				".....",
			},
			expected: true,
		},
		{
			name:  "Symbol diagonal (1)",
			input: NumberToken{143, 1, 3},
			line:  1,
			lines: []string{
				"....@",
				".143.",
				".....",
				".....",
			},
			expected: true,
		},
		{
			name:  "Symbol diagonal (2)",
			input: NumberToken{143, 1, 3},
			line:  1,
			lines: []string{
				".....",
				".143.",
				"!....",
				".....",
			},
			expected: true,
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			sut := NewPart1(test.lines)
			got := calculateSymbols(sut, test.input, test.line)
			assert.Equal(t, test.expected, got)
		})
	}
}

func TestPart2Calculation(t *testing.T) {
	tests := []struct {
		name     string
		input    NumberToken
		line     int
		lines    []string
		expected map[string][]int
		skip     bool
	}{
		{
			name:  "No adjacent symbols",
			input: NumberToken{143, 1, 3},
			line:  1,
			lines: []string{
				".....",
				".143.",
				".....",
				".....",
			},
			expected: map[string][]int{},
		},
		{
			name:  "Star above",
			input: NumberToken{143, 1, 3},
			line:  1,
			lines: []string{
				".*...",
				".143.",
				".....",
				".....",
			},
			expected: map[string][]int{
				"0.1": []int{143},
			},
		},
		{
			name:  "Other symbol above",
			input: NumberToken{143, 0, 2},
			line:  1,
			lines: []string{
				".&...",
				"143..",
				".....",
				".....",
			},
			expected: map[string][]int{},
		},
		{
			name:  "Star above (EOF)",
			input: NumberToken{143, 1, 3},
			line:  3,
			lines: []string{
				".....",
				".....",
				"..*..",
				".143.",
			},
			expected: map[string][]int{
				"2.2": []int{143},
			},
		},
		{
			name:  "Star above and below",
			input: NumberToken{143, 1, 3},
			line:  3,
			lines: []string{
				".....",
				".....",
				"..*..",
				".143.",
				"....*",
			},
			expected: map[string][]int{
				"2.2": []int{143},
				"4.4": []int{143},
			},
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			sut := NewPart2(test.lines)
			calculateSymbols(sut, test.input, test.line)
			got := sut.PossibleGears()
			assert.Equal(t, test.expected, got)
		})
	}
}

func TestParseLine(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []NumberToken
		skip     bool
	}{
		{
			name:  "Single digit at start",
			input: "1...",
			expected: []NumberToken{
				{1, 0, 0},
			},
		},
		{
			name:  "Multiple digits at start",
			input: "143...",
			expected: []NumberToken{
				{143, 0, 2},
			},
		},
		{
			name:  "Two tokens",
			input: "143..26.",
			expected: []NumberToken{
				{143, 0, 2},
				{26, 5, 6},
			},
		},
		{
			name:  "Three tokens",
			input: "143..26..%.2168.",
			expected: []NumberToken{
				{143, 0, 2},
				{26, 5, 6},
				{2168, 11, 14},
			},
		},
		{
			name:     "No tokens",
			input:    ".&..#...",
			expected: []NumberToken{},
		},
		{
			name:  "Token at EOL",
			input: "143..26",
			expected: []NumberToken{
				{143, 0, 2},
				{26, 5, 6},
			},
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

func TestSeekNumber(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		pos           int
		expectedToken NumberToken
		expectedBool  bool
		skip          bool
	}{
		{
			name:          "Single digit at start",
			input:         "1...",
			expectedToken: NumberToken{1, 0, 0},
			expectedBool:  true,
		},
		{
			name:          "Multiple digits at start",
			input:         "143...",
			expectedToken: NumberToken{143, 0, 2},
			expectedBool:  true,
		},
		{
			name:          "Multiple digits in middle",
			input:         "..143...",
			expectedToken: NumberToken{143, 2, 4},
			expectedBool:  true,
		},
		{
			name:          "Multiple digits at end",
			input:         "..143",
			expectedToken: NumberToken{143, 2, 4},
			expectedBool:  true,
		},
		{
			name:          "Multiple digits with symbol",
			input:         "..143&..",
			expectedToken: NumberToken{143, 2, 4},
			expectedBool:  true,
		},
		{
			name:          "No digits",
			input:         ".....",
			expectedToken: NumberToken{},
			expectedBool:  false,
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			gotToken, gotBool := seekNumber([]rune(test.input), &test.pos)
			assert.Equal(t, test.expectedToken, gotToken)
			assert.Equal(t, test.expectedBool, gotBool)
		})
	}
}

func TestIsDigit(t *testing.T) {
	tests := []struct {
		input    rune
		expected bool
		skip     bool
	}{
		{
			input:    rune('9'),
			expected: true,
		},
		{
			input:    rune('0'),
			expected: true,
		},
		{
			input:    rune('1'),
			expected: true,
		},
		{
			input:    rune('&'),
			expected: false,
		},
		{
			input:    rune('*'),
			expected: false,
		},
		{
			input:    rune('.'),
			expected: false,
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.input))
		}

		t.Run(string(test.input), func(t *testing.T) {
			t.Log(test.input)
			got := isDigit(test.input)
			assert.Equal(t, test.expected, got)
		})
	}
}

func TestIsSymbol(t *testing.T) {
	tests := []struct {
		input    byte
		expected bool
		skip     bool
	}{
		{
			input:    byte('9'),
			expected: false,
		},
		{
			input:    byte('0'),
			expected: false,
		},
		{
			input:    byte('1'),
			expected: false,
		},
		{
			input:    byte('&'),
			expected: true,
		},
		{
			input:    byte('*'),
			expected: true,
		},
		{
			input:    byte('.'),
			expected: false,
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.input))
		}

		t.Run(string(test.input), func(t *testing.T) {
			t.Log(test.input)
			got := isSymbol(test.input)
			assert.Equal(t, test.expected, got)
		})
	}
}
