package day1

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
		{
			input:    "two1nine",
			expected: 29,
		},
		{
			input:    "eightwothree",
			expected: 83,
		},
		{
			input:    "abcone2threexyz",
			expected: 13,
		},
		{
			input:    "xtwone3four",
			expected: 24,
		},
		{
			input:    "4nineeightseven2",
			expected: 42,
		},
		{
			input:    "zoneight234",
			expected: 14,
		},
		{
			input:    "7pqrstsixteen",
			expected: 76,
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
		input        byte
		expectedInt  int
		expectedBool bool
		skip         bool
	}{
		{
			input:        byte('a'),
			expectedInt:  0,
			expectedBool: false,
		},
		{
			input:        byte('9'),
			expectedInt:  9,
			expectedBool: true,
		},
		{
			input:        byte('t'),
			expectedInt:  0,
			expectedBool: false,
		},
		{
			input:        byte('0'),
			expectedInt:  0,
			expectedBool: false,
		},
		{
			input:        byte('1'),
			expectedInt:  1,
			expectedBool: true,
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.input))
		}

		t.Run(string(test.input), func(t *testing.T) {
			t.Log(test.input)
			gotInt, gotBool := isDigit(test.input)
			assert.Equal(t, test.expectedInt, gotInt)
			assert.Equal(t, test.expectedBool, gotBool)
		})
	}
}

func TestIsSpelledDigit(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		index        int
		expectedInt  int
		expectedBool bool
		skip         bool
	}{
		{
			name:         "eightseventysix match eight",
			input:        "eightseventysix",
			index:        0,
			expectedInt:  8,
			expectedBool: true,
		},
		{
			name:         "eightseventysix no match",
			input:        "eightseventysix",
			index:        1,
			expectedInt:  0,
			expectedBool: false,
		},
		{
			name:         "eightseventysix match seven",
			input:        "eightseventysix",
			index:        5,
			expectedInt:  7,
			expectedBool: true,
		},
		{
			name:         "eightseventysix match six",
			input:        "eightseventysix",
			index:        12,
			expectedInt:  6,
			expectedBool: true,
		},
		{
			name:         "eightseventysix short circuit",
			input:        "eightseventysix",
			index:        13,
			expectedInt:  0,
			expectedBool: false,
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.input))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			gotInt, gotBool := isSpelledDigit(test.input, test.index)
			assert.Equal(t, test.expectedInt, gotInt)
			assert.Equal(t, test.expectedBool, gotBool)
		})
	}
}
