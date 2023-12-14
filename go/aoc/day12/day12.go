package day12

import (
	"strings"

	"github.com/jeffrosenberg/advent-of-code-2023/go/pkg/aoc"
)

const (
	CONDITION_OPERATIONAL = iota
	CONDITION_UNKNOWN     = iota
	CONDITION_DAMAGED     = iota
)

type Solver interface { // Extends aoc.Solver
	aoc.Solver
}

type group struct {
	position  int
	length    int
	condition int
	complete  bool
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
	inputs, groups, wanted := parse(p)
	for i := 0; i < len(inputs); i++ {
		p.value += calculateArrangements(inputs[i], groups[i], wanted[i])
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

func calculateArrangements(input string, groups []group, wanted []int) int {
	if arrangementMeetsConditions(parseConditions(input, true), wanted) {
		return 1
	}
	return 0
}

func arrangementMeetsConditions(groups []group, wanted []int) bool {
	if len(groups) != len(wanted) {
		return false
	}
	for i, w := range wanted {
		if groups[i].length == w {
			continue
		} else {
			return false
		}
	}

	return true
}

func parse(solver Solver) (groupInputs []string, groups [][]group, wanted [][]int) {
	groups = [][]group{}
	wanted = [][]int{}

	for _, line := range solver.Lines() {
		i, g, a := parseLine(line)
		groupInputs = append(groupInputs, i)
		groups = append(groups, g)
		wanted = append(wanted, a)
	}
	return
}

func parseLine(line string) (groupInput string, groups []group, wanted []int) {
	groupInput, wantedInput, ok := strings.Cut(line, " ")
	if !ok {
		panic("unable to parse line")
	}

	groups = parseConditions(groupInput, false)
	wanted = parseWanted(wantedInput)

	return
}

func parseConditions(input string, matchesOnly bool) (groups []group) {
	var thisGroup group
	for i := 0; i < len(input); i++ {
		condition := parseCharacter(input[i])
		if i == 0 {
			thisGroup = group{
				position:  i,
				length:    1,
				condition: condition,
			}
		} else if condition == thisGroup.condition {
			thisGroup.length = thisGroup.length + 1
		} else {
			if !matchesOnly || condition == CONDITION_DAMAGED {
				groups = append(groups, thisGroup)
			}
			thisGroup = group{
				position:  i,
				length:    1,
				condition: condition,
			}
		}
	}
	if !matchesOnly || thisGroup.condition == CONDITION_DAMAGED {
		groups = append(groups, thisGroup)
	}
	return
}

func parseWanted(input string) (wanted []int) {
	arrStr := strings.Split(input, ",")
	for _, input := range arrStr {
		arrInt, _ := aoc.ConvertInt(input)
		wanted = append(wanted, arrInt)
	}
	return
}

func parseCharacter(char byte) int {
	switch char {
	case byte('.'):
		return CONDITION_OPERATIONAL
	case byte('#'):
		return CONDITION_DAMAGED
	default:
		return CONDITION_UNKNOWN
	}
}
