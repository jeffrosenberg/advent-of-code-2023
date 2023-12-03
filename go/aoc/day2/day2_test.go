package day2

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGameId(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expected      int
		expectedError bool
		skip          bool
	}{
		{
			name:     "Game 1",
			input:    "Game 1: 1 green, 1 blue, 1 red;",
			expected: 1,
		},
		{
			name:     "Game 2",
			input:    "Game 2: 9 red, 7 green, 3 blue;",
			expected: 2,
		},
		{
			name:          "Scan error",
			input:         "",
			expected:      0,
			expectedError: true,
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			scanner := bufio.NewScanner(strings.NewReader(test.input))
			scanner.Split(bufio.ScanWords)
			got, err := getGameId(scanner)
			if test.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
		})
	}
}

func TestGetCubeDraw(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedColor string
		expectedNum   int
		expectedError bool
		skip          bool
	}{
		{
			name:          "With comma",
			input:         "1 green,",
			expectedColor: "green",
			expectedNum:   1,
		},
		{
			name:          "With semicolon",
			input:         "3 blue;",
			expectedColor: "blue",
			expectedNum:   3,
		},
		{
			name:          "No punctuation",
			input:         "20 red",
			expectedColor: "red",
			expectedNum:   20,
		},
		{
			name:          "EOF",
			input:         "",
			expectedColor: "",
			expectedNum:   0,
		},
		{
			name:          "Scan error",
			input:         "2",
			expectedError: true,
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			scanner := bufio.NewScanner(strings.NewReader(test.input))
			scanner.Split(bufio.ScanWords)
			gotColor, gotNum, err := getCubeDraw(scanner)
			if test.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expectedColor, gotColor)
				assert.Equal(t, test.expectedNum, gotNum)
			}
		})
	}
}

func TestIsPossible(t *testing.T) {
	tests := []struct {
		name     string
		color    string
		quantity int
		expected bool
		skip     bool
	}{
		{
			name:     "red 11",
			color:    "red",
			quantity: 11,
			expected: true,
		},
		{
			name:     "red 12",
			color:    "red",
			quantity: 12,
			expected: true,
		},
		{
			name:     "red 13",
			color:    "red",
			quantity: 13,
			expected: false,
		},
		{
			name:     "green 13",
			color:    "green",
			quantity: 13,
			expected: true,
		},
		{
			name:     "green 14",
			color:    "green",
			quantity: 14,
			expected: false,
		},
		{
			name:     "blue 14",
			color:    "blue",
			quantity: 14,
			expected: true,
		},
		{
			name:     "blue 15",
			color:    "blue",
			quantity: 15,
			expected: false,
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			got := isPossible(test.color, test.quantity)
			assert.Equal(t, test.expected, got)
		})
	}
}
