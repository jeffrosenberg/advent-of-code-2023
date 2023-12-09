package day8

import (
	"regexp"

	"github.com/jeffrosenberg/advent-of-code-2023/go/pkg/aoc"
)

const START = "AAA"
const GOAL = "ZZZ"

type Solver interface { // Extends aoc.Solver
	AddValue(int)
	aoc.Solver
}

type node struct {
	left  string
	right string
}

type Part1 struct {
	lines        []string
	value        int
	instructions string
	nodes        map[string]node
}

type Part2 struct {
	lines []string
	value int
}

func NewPart1(lines []string) *Part1 {
	p := Part1{
		lines: lines,
		value: 0,
		nodes: map[string]node{},
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
	p.instructions, p.nodes = parse(p)
	p.value = traverse(p.nodes, p.instructions)
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

func traverse(nodes map[string]node, instructions string) (steps int) {
	instr := []rune(instructions)
	len_instr := len(instr)
	var currentNode string = START

	for currentNode != GOAL {
		theNode := nodes[currentNode]
		instruction := instr[steps%len_instr]
		if instruction == 'L' {
			currentNode = theNode.left
		} else {
			currentNode = theNode.right
		}
		steps++
	}

	return
}

func parse(solver Solver) (instructions string, nodes map[string]node) {
	nodes = map[string]node{}
	for i, line := range solver.Lines() {
		if i == 0 {
			instructions = line
		} else if len(line) == 0 {
			continue
		} else {
			k, n := parseLine(line)
			nodes[k] = n
		}
	}
	return
}

func parseLine(line string) (string, node) {
	// Given the structure of this input, regex seems like a less annoying way to parse
	re := regexp.MustCompile(`^(\w{3})\s=\s\((\w{3}),\s(\w{3})\)$`)
	match := re.FindStringSubmatch(line)
	if len(match) != 4 {
		panic("Failed to match regex")
	}
	return match[1], node{left: match[2], right: match[3]}
}
