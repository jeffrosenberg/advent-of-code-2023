package day9

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jeffrosenberg/advent-of-code-2023/go/pkg/aoc"
)

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
	histories := parse(p)
	for _, hist := range histories {
		p.value += extrapolate(hist)
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
	histories := parse(p)
	for _, hist := range histories {
		p.value += extrapolateBackwards(hist)
	}
}

func extrapolate(history []int) (next int) {
	diffs := []int{}
	nextInterval := 0

	for i := 1; i < len(history); i++ {
		diffs = append(diffs, history[i]-history[i-1])
	}
	if !(diffs[0] == 0 && diffs[len(diffs)-1] == 0) {
		nextInterval = extrapolate(diffs)
	}
	return history[len(history)-1] + nextInterval
}

func extrapolateBackwards(history []int) (prev int) {
	diffs := []int{}
	prevInterval := 0

	for i := 1; i < len(history); i++ {
		diffs = append(diffs, history[i]-history[i-1])
	}
	if !(diffs[0] == 0 && diffs[len(diffs)-1] == 0) {
		prevInterval = extrapolateBackwards(diffs)
	}
	return history[0] - prevInterval
}

func parse(solver aoc.Solver) (output [][]int) {
	output = [][]int{}
	for _, line := range solver.Lines() {
		output = append(output, parseLine(line))
	}
	return output
}

func parseLine(line string) (output []int) {
	output = []int{}
	for _, str := range strings.Split(line, " ") {
		i, err := strconv.Atoi(str)
		if err != nil {
			panic(fmt.Sprintf("Unable to parse line: %s", line))
		}
		output = append(output, i)
	}
	return
}
