package day11

import (
	"math"
	"sort"

	"github.com/jeffrosenberg/advent-of-code-2023/go/pkg/aoc"
)

type Solver interface { // Extends aoc.Solver
	aoc.Solver
}

type point struct {
	x int
	y int
}

func getPathLength(from point, to point) int {
	y := math.Abs(float64(from.y) - float64(to.y))
	x := math.Abs(float64(from.x) - float64(to.x))
	return int(x + y)
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
	points := parseAndExpand(p)
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			p.value += getPathLength(points[i], points[j])
		}
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
	// TODO
}

// TODO: I bet performance on this stinks, revisit after initial solution
// 1. Make it work
// 2. Make it right
// 3. Make it fast
func parseAndExpand(solver Solver) (output []point) {
	output = []point{}
	parsed := []*point{}
	columns := map[int]int{}
	emptyLines := []int{}
	emptyColumns := []int{}
	lines := solver.Lines()

	// Parse
	for i, line := range lines {
		y := len(lines) - 1 - i // Thanks zero-based indexing :-/
		points, cols := parseLine(line, y)

		if len(points) == 0 {
			emptyLines = append(emptyLines, y)
			continue
		}

		parsed = append(parsed, points...)
		for _, col := range cols {
			columns[col] += 1
		}
	}

	for i := 0; i < len(lines[0]); i++ {
		if columns[i] == 0 {
			emptyColumns = append(emptyColumns, i)
		}
	}

	// Expand - keeping part of one function to make it easier to act on the same variables
	// TODO: Refactor into its own function?
	// Should already be sorted in reverse order
	for _, y := range emptyLines {
		for _, pt := range parsed {
			// Since these are pointers, we should be able to update them directly
			if pt.y > y {
				pt.y += 1
			}
		}
	}

	sort.Ints(emptyColumns)
	for i := len(emptyColumns) - 1; i >= 0; i-- {
		x := emptyColumns[i]
		for _, pt := range parsed {
			// Since these are pointers, we should be able to update them directly
			if pt.x > x {
				pt.x += 1
			}
		}
	}

	for _, pt := range parsed {
		output = append(output, *pt)
	}
	return output
}

func parseLine(line string, y int) (points []*point, columns []int) { // TODO: Set return type
	points = []*point{}
	columns = []int{}

	for x, char := range line {
		if char == '#' {
			newPoint := point{
				x: x,
				y: y,
			}
			points = append(points, &newPoint)
			columns = append(columns, x)
		}
	}

	return
}
