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
	for _, pattern := range patterns {
		p.value += calculateReflection(pattern)
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
	for _, pattern := range patterns {
		p.value += calculateReflection(pattern)
	}
}

func calculateReflection(pt pattern) int {
	if line, found := hasReflection(pt.cols); found {
		return line
	}
	if line, found := hasReflection(pt.rows); found {
		return line * 100
	}
	return 0
}

func hasReflection(lines []string) (linePosition int, found bool) {
	// TODO: This could be optimized by starting towards the center
	l := len(lines)
	for i := 1; i < l; i++ {
		// Definitely no match at this index; short circuit
		if lines[i-1] != lines[i] {
			continue
		}

		// Possible match, explore further
		matched := true
		reflectionSize := aoc.Min(i, len(lines[i:]))
		for j := 1; j < reflectionSize; j++ {
			if lines[i-1-j] != lines[i+j] {
				matched = false
				break
			}
		}
		if matched {
			return i, true
		}
	}
	return
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
