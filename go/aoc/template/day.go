package dayn

import (
	"github.com/jeffrosenberg/advent-of-code-2023/go/pkg/aoc"
)

type Solver interface { // Extends aoc.Solver
	AddValue(int)
	aoc.Solver
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

func (p *Part1) AddValue(val int) {
	p.value += val
}

func (p *Part1) Solve() {
	// TODO
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

func (p *Part2) AddValue(val int) {
	p.value += val
}

func (p *Part2) Solve() {
	// TODO
}

func supportingFunc(input string) bool {
	return false
}
