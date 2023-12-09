package day8

import (
	"regexp"

	"github.com/jeffrosenberg/advent-of-code-2023/go/pkg/aoc"
)

type Solver interface { // Extends aoc.Solver
	isStartNode(string) bool
	isEndNode(string) bool
	Instructions() []rune
	Nodes() map[string]node
	StartNodes() []string
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
	startNodes   []string
}

type Part2 struct {
	lines        []string
	value        int
	instructions string
	nodes        map[string]node
	startNodes   []string
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

func (p *Part1) isStartNode(key string) bool {
	return key == "AAA"
}

func (p *Part1) isEndNode(key string) bool {
	return key == "ZZZ"
}

func (p *Part1) Instructions() []rune {
	return []rune(p.instructions)
}

func (p *Part1) Nodes() map[string]node {
	return p.nodes
}

func (p *Part1) StartNodes() []string {
	return p.startNodes
}

func (p *Part1) Solve() {
	p.instructions, p.nodes, p.startNodes = parse(p)
	p.value = traverse(p)
}

func NewPart2(lines []string) *Part2 {
	p := Part2{
		lines: lines,
		value: 0,
		nodes: map[string]node{},
	}
	return &p
}

func (p *Part2) Lines() []string {
	return p.lines
}

func (p *Part2) Value() int {
	return p.value
}

func (p *Part2) isStartNode(key string) bool {
	return key[2] == byte('A')
}

func (p *Part2) isEndNode(key string) bool {
	return key[2] == byte('Z')
}

func (p *Part2) Instructions() []rune {
	return []rune(p.instructions)
}

func (p *Part2) Nodes() map[string]node {
	return p.nodes
}

func (p *Part2) StartNodes() []string {
	return p.startNodes
}

func (p *Part2) Solve() {
	p.instructions, p.nodes, p.startNodes = parse(p)
}

func traverse(solver Solver) (steps int) {
	instr := solver.Instructions()
	len_instr := len(instr)
	nodes := solver.Nodes()
	currentNode := solver.StartNodes()[0]

	for !solver.isEndNode(currentNode) {
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

func parse(solver Solver) (instructions string, nodes map[string]node, startNodes []string) {
	nodes = map[string]node{}
	for i, line := range solver.Lines() {
		if i == 0 {
			instructions = line
		} else if len(line) == 0 {
			continue
		} else {
			k, n := parseLine(line)
			nodes[k] = n
			if solver.isStartNode(k) {
				startNodes = append(startNodes, k)
			}
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
