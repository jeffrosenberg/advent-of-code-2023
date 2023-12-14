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
		smudges       int
		expectedLine  int
		expectedFound bool
		skip          bool
	}{
		// No "smudges" - part 1
		{
			name:    "Simple true",
			smudges: 0,
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
			name:    "Simple false",
			smudges: 0,
			input: []string{
				".##.",
				"..#.",
				"....",
				".##.",
			},
			expectedLine:  0,
			expectedFound: false,
		},
		{
			name:    "Match one line, then no match",
			smudges: 0,
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
			name:    "Example 1 - column, ignore one row",
			smudges: 0,
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
			name:    "Example 2 - row, ignore one row",
			smudges: 0,
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
		// 1 "smudge" - part 2
		{
			name:    "Simple true with smudge",
			smudges: 1,
			input: []string{
				"###.",
				"..#.",
				"..#.",
				".##.",
			},
			expectedLine:  2,
			expectedFound: true,
		},
		{
			name:    "Simple false with smudge",
			smudges: 1,
			input: []string{
				"###.",
				"..#.",
				"..#.",
				".###",
			},
			expectedLine:  0,
			expectedFound: false,
		},
		{
			name:    "No smudges when one required",
			smudges: 1,
			input: []string{
				"###.",
				"..#.",
				"..#.",
				"###.",
			},
			expectedLine:  0,
			expectedFound: false,
		},
	}

	for _, test := range tests {
		t.Run(string(test.name), func(t *testing.T) {
			if test.skip {
				t.Skipf("Skipping %s", string(test.name))
			}

			t.Log(test.name)
			gotLine, gotFound := hasReflection(test.input, test.smudges)
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
