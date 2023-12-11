package day10

import (
	"fmt"

	"github.com/jeffrosenberg/advent-of-code-2023/go/pkg/aoc"
)

type Solver interface { // Extends aoc.Solver
	getPipe(key string) (pipe, bool)
	getPipes() map[string]pipe
	getStart() pipe
	aoc.Solver
}

type pipe struct {
	x           int
	y           int
	orientation rune
}

func getCoords(x int, y int) string {
	return fmt.Sprintf("%d.%d", x, y)
}

func (p *pipe) string() string {
	return getCoords(p.x, p.y)
}

func (p *pipe) getNext(from string) (options []string) {
	options = make([]string, 0, 3)
	var possibilities []string
	fromValidDestination := false

	north := getCoords(p.x, p.y+1)
	south := getCoords(p.x, p.y-1)
	east := getCoords(p.x+1, p.y)
	west := getCoords(p.x-1, p.y)

	switch p.orientation {
	case '|':
		possibilities = []string{south, north}
	case '-':
		possibilities = []string{west, east}
	case 'L':
		possibilities = []string{north, east}
	case 'J':
		possibilities = []string{north, west}
	case '7':
		possibilities = []string{south, west}
	case 'F':
		possibilities = []string{south, east}
	default:
		possibilities = []string{west, north, east, south}
	}

	for _, option := range possibilities {
		if option == from {
			fromValidDestination = true
			continue
		}
		options = append(options, option)
	}

	if !fromValidDestination && from != "" {
		return []string{}
	}

	return
}

type Part1 struct {
	lines []string
	value int
	pipes map[string]pipe
	start pipe
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

func (p *Part1) getPipes() map[string]pipe {
	return p.pipes
}

func (p *Part1) getPipe(key string) (pipe, bool) {
	result, ok := p.pipes[key]
	return result, ok
}

func (p *Part1) getStart() pipe {
	return p.start
}

func (p *Part1) Value() int {
	return p.value
}

func (p *Part1) Solve() {
	p.Parse()
	loop, _ := calculateLoop(p, []pipe{p.getStart()}, "")
	p.value = len(loop) / 2 // truncate the .5
}

func (p *Part1) Parse() {
	pipes, startingPipe := parse(p)
	p.pipes = pipes
	p.start = *startingPipe
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

func calculateLoop(solver Solver, from []pipe, prev string) (pipes []pipe, success bool) {
	var prevKey string
	// allPipes := solver.getPipes()
	this := from[len(from)-1]
	if len(from) < 2 {
		prevKey = ""
	} else {
		prevKey = from[len(from)-2].string()
	}
	options := this.getNext(prevKey)

	for _, o := range options {
		if next, ok := solver.getPipe(o); ok {
			if next == solver.getStart() {
				return append(from, next), true
			}
			if result, succ := calculateLoop(solver, append(from, next), this.string()); succ {
				return result, succ
			}
		}
	}

	return
}

func parse(solver Solver) (output map[string]pipe, startingPipe *pipe) {
	output = map[string]pipe{}
	lines := solver.Lines()
	for i, line := range lines {
		pipes, start := parseLine(line, len(lines)-1-i) // Thanks zero-based indexing :-/
		for _, p := range pipes {
			output[p.string()] = p
		}
		if nil != start {
			startingPipe = start
		}
	}
	return
}

func parseLine(line string, lineNum int) (output []pipe, start *pipe) {
	output = make([]pipe, 0, len(line))
	for i, char := range line {
		if char != '.' {
			newPipe := pipe{
				x:           i,
				y:           lineNum,
				orientation: char,
			}
			output = append(output, newPipe)
			if char == 'S' {
				start = &newPipe
			}
		}
	}
	return
}
