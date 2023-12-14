package day13

import (
	"testing"

	"github.com/jeffrosenberg/advent-of-code-2023/go/pkg/aoc"
	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	p := NewPart1(aoc.ReadAocInput("../../../inputs/13_test.txt"))
	expected := 405
	p.Solve()
	assert.Equal(t, expected, p.Value())
}

func TestHasReflection(t *testing.T) {
	tests := []struct {
		name          string
		input         []string
		expectedLine  int
		expectedFound bool
		skip          bool
	}{
		{
			name: "Simple true",
			input: []string{
				".##.",
				"..#.",
				"..#.",
				".##.",
			},
			expectedLine:  2,
			expectedFound: true,
		},
		{
			name: "Simple false",
			input: []string{
				".##.",
				".##.",
				"....",
				".##.",
			},
			expectedLine:  0,
			expectedFound: false,
		},
		{
			name: "Match one line, then no match",
			input: []string{
				"....",
				"..#.",
				"..#.",
				"####",
			},
			expectedLine:  0,
			expectedFound: false,
		},
		{
			name: "Example 1 - column, ignore one row",
			input: []string{
				"#.##..#",
				"..##...",
				"##..###",
				"#....#.",
				".#..#.#",
				".#..#.#",
				"#....#.",
				"##..###",
				"..##...",
			},
			expectedLine:  5,
			expectedFound: true,
		},
		{
			name: "Example 2 - row, ignore one row",
			input: []string{
				"#...##..#",
				"#....#..#",
				"..##..###",
				"#####.##.",
				"#####.##.",
				"..##..###",
				"#....#..#",
			},
			expectedLine:  4,
			expectedFound: true,
		},
	}

	for _, test := range tests {
		t.Run(string(test.name), func(t *testing.T) {
			if test.skip {
				t.Skipf("Skipping %s", string(test.name))
			}

			t.Log(test.name)
			gotLine, gotFound := hasReflection(test.input)
			assert.Equal(t, test.expectedLine, gotLine)
			assert.Equal(t, test.expectedFound, gotFound)
		})
	}
}

func TestParse(t *testing.T) {
	p := NewPart1(aoc.ReadAocInput("../../../inputs/13_test.txt"))
	expected := []pattern{
		{
			rows: []string{
				"#.##..##.",
				"..#.##.#.",
				"##......#",
				"##......#",
				"..#.##.#.",
				"..##..##.",
				"#.#.##.#.",
			},
			cols: []string{
				"#.##..#",
				"..##...",
				"##..###",
				"#....#.",
				".#..#.#",
				".#..#.#",
				"#....#.",
				"##..###",
				"..##...",
			},
		},
		{
			rows: []string{
				"#...##..#",
				"#....#..#",
				"..##..###",
				"#####.##.",
				"#####.##.",
				"..##..###",
				"#....#..#",
			},
			cols: []string{
				"##.##.#",
				"...##..",
				"..####.",
				"..####.",
				"#..##..",
				"##....#",
				"..####.",
				"..####.",
				"###..##",
			},
		},
	}
	got := parse(p)
	assert.Equal(t, expected, got)
}
