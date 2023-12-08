package day5

import (
	"testing"

	"github.com/jeffrosenberg/advent-of-code-2023/go/pkg/aoc"
	"github.com/stretchr/testify/assert"
)

func TestTraverse(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
		skip     bool
	}{
		{
			name:     "Seed 79 -> Location 82",
			input:    79,
			expected: 82,
		},
		{
			name:     "Seed 14 -> Location 43",
			input:    14,
			expected: 43,
		},
		{
			name:     "Seed 55 -> Location 86",
			input:    55,
			expected: 86,
		},
	}

	p := NewPart1(aoc.ReadAocInput("../../../inputs/5_test.txt"))
	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			p.parseSeedLine()
			p.d.parseMaps()
			got := p.d.traverse(test.input)
			assert.Equal(t, test.expected, got)
		})
	}
}

func TestGetDestination(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
		skip     bool
	}{
		{
			name:     "0 - 50: no change",
			input:    3,
			expected: 3,
		},
		{
			name:     "50 - 98: +2",
			input:    60,
			expected: 62,
		},
		{
			name:     "98 - 99: -48",
			input:    98,
			expected: 50,
		},
		{
			name:     "100+: no change",
			input:    200,
			expected: 200,
		},
	}

	expectedSeedToSoilMap := map[int]int{
		0:   0,
		98:  -48,
		50:  2,
		100: 0,
	}
	keys := []int{0, 98, 50, 100}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			got := getDestination(test.input, expectedSeedToSoilMap, keys)
			assert.Equal(t, test.expected, got)
		})
	}
}

func TestParsePart1(t *testing.T) {
	expectedSeeds := []int{79, 14, 55, 13}
	expectedSeedToSoilMap := map[int]int{
		0:   0,
		50:  2,
		98:  -48,
		100: 0,
	}
	expectedTemperatureToHumidityMap := map[int]int{
		0:  1,
		69: -69,
		70: 0,
	}
	p := NewPart1(aoc.ReadAocInput("../../../inputs/5_test.txt"))
	p.parseSeedLine()
	p.d.parseMaps()
	assert.Equal(t, expectedSeeds, p.seeds)
	assert.Equal(t, expectedSeedToSoilMap, p.d.maps[0])
	assert.Equal(t, expectedTemperatureToHumidityMap, p.d.maps[5])
}

func TestParsePart2(t *testing.T) {
	expectedSeeds := []seedRange{
		{start: 79, end: 93},
		{start: 55, end: 68},
	}
	p := NewPart2(aoc.ReadAocInput("../../../inputs/5_test.txt"))
	p.parseSeedLine()
	assert.Equal(t, expectedSeeds, p.seeds)
}

func TestBuildMap(t *testing.T) {
	tests := []struct {
		name          string
		inputString   string
		inputMap      map[int]int
		inputSlice    []int
		expectedMap   map[int]int
		expectedSlice []int
		skip          bool
	}{
		{
			name:        "Example test 1",
			inputString: "50 98 2",
			inputMap:    map[int]int{},
			inputSlice:  []int{},
			expectedMap: map[int]int{
				98:  -48,
				100: 0,
			},
			expectedSlice: []int{98, 100},
		},
		{
			name:        "Example test 2",
			inputString: "42 0 7",
			inputMap:    map[int]int{},
			inputSlice:  []int{},
			expectedMap: map[int]int{
				0: 42,
				7: 0,
			},
			expectedSlice: []int{0, 7},
		},
		{
			name:        "Combine over multiple lines",
			inputString: "42 0 7",
			inputMap: map[int]int{
				98:  -48,
				100: 0,
			},
			inputSlice: []int{98, 100},
			expectedMap: map[int]int{
				0:   42,
				7:   0,
				98:  -48,
				100: 0,
			},
			expectedSlice: []int{98, 100, 0, 7}, // Will get sorted in later methods
		},
	}

	for _, test := range tests {
		if test.skip {
			t.Skipf("Skipping %s", string(test.name))
		}

		t.Run(string(test.name), func(t *testing.T) {
			t.Log(test.name)
			gotKeys := buildMap(test.inputString, test.inputMap, test.inputSlice)
			// test.inputMap should be edited in-place
			assert.Equal(t, test.expectedMap, test.inputMap)
			assert.Equal(t, test.expectedSlice, gotKeys)
		})
	}
}
