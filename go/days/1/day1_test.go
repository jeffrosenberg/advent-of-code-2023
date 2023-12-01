package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractCalibrationValue(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		skip     bool
	}{
		{
			input:    "1abc2",
			expected: 12,
		},
		{
			input:    "pqr3stu8vwx",
			expected: 38,
		},
		{
			input:    "a1b2c3d4e5f",
			expected: 15,
		},
		{
			input:    "treb7uchet",
			expected: 77,
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", test.input)
		}

		t.Run(test.input, func(t *testing.T) {
			t.Log(test.input)
			got, err := extractCalibrationValue(test.input)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, got)
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
			input:    'a',
			expected: false,
		},
		{
			input:    '9',
			expected: true,
		},
		{
			input:    't',
			expected: false,
		},
		{
			input:    '0',
			expected: true,
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
