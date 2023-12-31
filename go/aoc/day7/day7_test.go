package day7

import (
	"testing"

	"github.com/jeffrosenberg/advent-of-code-2023/go/pkg/aoc"
	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	expected := 6440

	p := NewPart1(aoc.ReadAocInput("../../../inputs/7_test.txt"))
	p.Solve()

	assert.Equal(t, expected, p.Value())
}

func TestSolvePart2(t *testing.T) {
	expected := 5905

	p := NewPart2(aoc.ReadAocInput("../../../inputs/7_test.txt"))
	p.Solve()

	assert.Equal(t, expected, p.Value())
}

func TestSort(t *testing.T) {
	tests := []struct {
		name     string
		input    *Part1
		expected []string
		skip     bool
	}{
		{
			name: "Pair and two three of a kind",
			input: &Part1{
				hands: []hand{
					{
						cards:    "QQQJA",
						handType: HAND_THREEKIND,
						strength: "4CCCBE",
					},
					{
						cards:    "32T3K",
						handType: HAND_PAIR,
						strength: "232A3D",
					},
					{
						cards:    "T55J5",
						handType: HAND_THREEKIND,
						strength: "4A55B5",
					},
				},
			},
			expected: []string{
				"32T3K",
				"T55J5",
				"QQQJA",
			},
		},
		{
			name: "Two five of a kind and two full house",
			input: &Part1{
				hands: []hand{
					{
						cards:    "55555",
						handType: HAND_FIVEKIND,
						strength: "755555",
					},
					{
						cards:    "QQQJJ",
						handType: HAND_FULLHOUSE,
						strength: "5CCCBB",
					},
					{
						cards:    "JJJQQ",
						handType: HAND_FULLHOUSE,
						strength: "5CCBBB",
					},
					{
						cards:    "44444",
						handType: HAND_FIVEKIND,
						strength: "744444",
					},
				},
			},
			expected: []string{
				"JJJQQ",
				"QQQJJ",
				"44444",
				"55555",
			},
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			sortByStrength(test.input.hands)
			for i, exp := range test.expected {
				assert.Equal(t, exp, test.input.hands[i].cards)
			}
		})
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected hand
		skip     bool
	}{
		{
			name:  "First example",
			input: []string{"32T3K 765"},
			expected: hand{
				cards:    "32T3K",
				bid:      765,
				handType: HAND_PAIR,
				strength: "232A3D",
			},
		},
		{
			name:  "Second example",
			input: []string{"T55J5 684"},
			expected: hand{
				cards:    "T55J5",
				bid:      684,
				handType: HAND_THREEKIND,
				strength: "4A55B5",
			},
		},
		{
			name:  "Fifth example",
			input: []string{"QQQJA 483"},
			expected: hand{
				cards:    "QQQJA",
				bid:      483,
				handType: HAND_THREEKIND,
				strength: "4CCCBE",
			},
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			p := NewPart1(test.input)
			parse(p)
			assert.Equal(t, test.expected, p.hands[0])
		})
	}
}

func TestCalculateStrengthPart1(t *testing.T) {
	tests := []struct {
		name      string
		inputHand hand
		expected  string
		skip      bool
	}{
		{
			name: "First example - pair",
			inputHand: hand{
				cards: "32T3K",
			},
			expected: "232A3D",
		},
		{
			name: "Second example - three of a kind",
			inputHand: hand{
				cards: "T55J5",
			},
			expected: "4A55B5",
		},
		{
			name: "Third example - two pair",
			inputHand: hand{
				cards: "KK677",
			},
			expected: "3DD677",
		},
		{
			name: "Full House",
			inputHand: hand{
				cards: "QQQJJ",
			},
			expected: "5CCCBB",
		},
	}

	p := NewPart1([]string{}) // Don't actually care about any values here
	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			test.inputHand.calculateStrength(p.faceValues(), p.calculateHandType)
			assert.Equal(t, test.expected, test.inputHand.strength)
		})
	}
}

func TestCalculateStrengthPart2(t *testing.T) {
	tests := []struct {
		name      string
		inputHand hand
		expected  string
		skip      bool
	}{
		{
			name: "First example - pair",
			inputHand: hand{
				cards: "32T3K",
			},
			expected: "232A3D",
		},
		{
			name: "Second example - four of a kind",
			inputHand: hand{
				cards: "T55J5",
			},
			expected: "6A5515",
		},
		{
			name: "Third example - two pair",
			inputHand: hand{
				cards: "KK677",
			},
			expected: "3DD677",
		},
		{
			name: "Five of a kind",
			inputHand: hand{
				cards: "QQQJJ",
			},
			expected: "7CCC11",
		},
	}

	p := NewPart2([]string{}) // Don't actually care about any values here
	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			test.inputHand.calculateStrength(p.faceValues(), p.calculateHandType)
			assert.Equal(t, test.expected, test.inputHand.strength)
		})
	}
}
