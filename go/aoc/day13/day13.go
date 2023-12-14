package day13

import (
	"strings"

	"github.com/jeffrosenberg/advent-of-code-2023/go/pkg/aoc"
)

type Solver interface { // Extends aoc.Solver
	aoc.Solver
}

type pattern struct {
	rows []string
	cols []string
}

type Part1 struct {
	lines []string
	value int
}

type Part2 struct {
	lines []string
	value int
}

func NewPart1(lines []string) *Part1 {
	p := Part1{
		lines: lines,
		value: 0,
	}
	return &p
}

func (p *Part1) Lines() []string {
	return p.lines
}

func (p *Part1) Value() int {
	return p.value
}

func (p *Part1) Solve() {
	patterns := parse(p)
	smudges := 0
	for _, pattern := range patterns {
		p.value += calculateReflection(pattern, smudges)
	}
}

func NewPart2(lines []string) *Part2 {
	p := Part2{
		lines: lines,
		value: 0,
	}
	return &p
}

func (p *Part2) Lines() []string {
	return p.lines
}

func (p *Part2) Value() int {
	return p.value
}

func (p *Part2) Solve() {
	patterns := parse(p)
	smudges := 1
	for _, pattern := range patterns {
		p.value += calculateReflection(pattern, smudges)
	}
}

func calculateReflection(pt pattern, smudges int) int {
	if line, found := hasReflection(pt.cols, smudges); found {
		return line
	}
	if line, found := hasReflection(pt.rows, smudges); found {
		return line * 100
	}

	return 0
}

func hasReflection(lines []string, smudges int) (linePosition int, found bool) {
	// TODO: This could be optimized by starting towards the center
	l := len(lines)
	for i := 1; i < l; i++ {
		smudgesFound := 0
		matched := true
		reflectionSize := aoc.Min(i, len(lines[i:]))

		for j := 0; j < reflectionSize && smudgesFound <= smudges; j++ {
			line1 := lines[i-1-j]
			line2 := lines[i+j]
			if line1 != line2 {
				if isSmudge(line1, line2) {
					smudgesFound += 1
					// This still counts as a match, but we'll also match smudges
				} else {
					matched = false
					break
				}
			}
		}
		if matched && smudgesFound == smudges {
			return i, true
		}
	}
	return
}

func isSmudge(a string, b string) bool {
	if a == b {
		return false
	}

	// Test whether there's only a single character out of place
	missedChars := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			missedChars += 1
			if missedChars > 1 {
				return false
			}
		}
	}

	// This should be true if we've reached this point
	if missedChars == 1 {
		return true
	}

	return false // We shouldn't reach this line
}

func parse(solver Solver) (output []pattern) {
	output = []pattern{}
	lines := solver.Lines()
	thisPattern := pattern{
		rows: []string{},
		cols: []string{},
	}
	startLine := 0

	for i, line := range lines {
		if len(line) > 0 {
			thisPattern.rows = append(thisPattern.rows, line)
		} else {
			for j := 0; j < len(lines[startLine]); j++ {
				thisPattern.cols = append(thisPattern.cols, transposeColumn(lines[startLine:i], j))
			}
			output = append(output, thisPattern)
			thisPattern = pattern{
				rows: []string{},
				cols: []string{},
			}
			startLine = i + 1
		}
	}

	for j := 0; j < len(lines[startLine]); j++ {
		thisPattern.cols = append(thisPattern.cols, transposeColumn(lines[startLine:], j))
	}
	output = append(output, thisPattern)
	return output
}

func transposeColumn(input []string, col int) string {
	bldr := strings.Builder{}
	for i := 0; i < len(input); i++ {
		bldr.WriteByte(input[i][col])
	}
	return bldr.String()
}
