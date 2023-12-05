package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
		skip     bool
	}{
		{
			name:     "Card 1",
			input:    []string{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"},
			expected: 8,
		},
		{
			name:     "Card 2",
			input:    []string{"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19"},
			expected: 2,
		},
		{
			name:     "Card 5",
			input:    []string{"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36"},
			expected: 0,
		},
		{
			name: "Cards 3 and 4",
			input: []string{
				"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
				"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
			},
			expected: 3,
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			s := NewPart1(test.input)
			s.Solve()
			assert.Equal(t, test.expected, s.Value())
		})
	}
}

func TestCalculateMatches(t *testing.T) {
	tests := []struct {
		name     string
		input    Line
		expected int
		skip     bool
	}{
		{
			name: "Card 1",
			input: Line{
				cardName: "Card 1",
				winners: map[int]bool{
					41: false,
					48: false,
					83: false,
					86: false,
					17: false,
				},
				given: []int{83, 86, 6, 31, 17, 9, 48, 53},
			},
			expected: 4,
		},
		{
			name: "Card 2",
			input: Line{
				cardName: "Card 2",
				winners: map[int]bool{
					13: false,
					32: false,
					20: false,
					16: false,
					61: false,
				},
				given: []int{61, 30, 68, 82, 17, 32, 24, 19},
			},
			expected: 2,
		},
		{
			name: "Card 5",
			input: Line{
				cardName: "Card 5",
				winners: map[int]bool{
					87: false,
					83: false,
					26: false,
					28: false,
					32: false,
				},
				given: []int{88, 30, 70, 12, 93, 22, 82, 36},
			},
			expected: 0,
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			test.input.CalculateMatches()
			assert.Equal(t, test.expected, test.input.matches)
		})
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Line
		skip     bool
	}{
		{
			name:  "Card 1",
			input: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			expected: Line{
				cardName: "Card 1",
				winners: map[int]bool{
					41: false,
					48: false,
					83: false,
					86: false,
					17: false,
				},
				given: []int{83, 86, 6, 31, 17, 9, 48, 53},
			},
		},
		{
			name:  "Card 2",
			input: "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
			expected: Line{
				cardName: "Card 2",
				winners: map[int]bool{
					13: false,
					32: false,
					20: false,
					16: false,
					61: false,
				},
				given: []int{61, 30, 68, 82, 17, 32, 24, 19},
			},
		},
		{
			name:  "Card 3",
			input: "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			expected: Line{
				cardName: "Card 3",
				winners: map[int]bool{
					1:  false,
					21: false,
					53: false,
					59: false,
					44: false,
				},
				given: []int{69, 82, 63, 72, 16, 21, 14, 1},
			},
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			got := parse(test.input)
			assert.Equal(t, test.expected, got)
		})
	}
}
