package day3

import (
	"fmt"
	"strconv"
)

type Day3Aoc interface {
	Lines() []string
	Value() int
	AddValue(int)
	Calculate(int, int, int) bool
}

type NumberToken struct {
	Number int
	Start  int
	End    int
}

type Part1 struct {
	lines []string
	value int
}

type Part2 struct {
	lines         []string
	possibleGears map[string][]int
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

func NewPart2(lines []string) *Part2 {
	p := Part2{
		lines:         lines,
		possibleGears: map[string][]int{},
	}
	return &p
}

func (p *Part2) Lines() []string {
	return p.lines
}

func (p *Part2) Value() int {
	// For part 2, need to calculate value from possibleGears
	val := 0
	for _, v := range p.possibleGears {
		if len(v) == 2 {
			val += v[0] * v[1]
		}
	}
	return val
}

func (p *Part2) AddValue(val int) {
	// no-op to fit within the interface
}

func (p *Part2) AddGear(key string, value int) {
	if gears, ok := p.possibleGears[key]; ok {
		p.possibleGears[key] = append(gears, value)
	} else {
		p.possibleGears[key] = []int{value}
	}
}
func (p *Part2) PossibleGears() map[string][]int {
	return p.possibleGears
}

func Answer(aoc Day3Aoc) int {
	for line := 0; line < len(aoc.Lines()); line++ {
		tokens := parseLine(aoc.Lines()[line])
		for i := 0; i < len(tokens); i++ {
			calculateSymbols(aoc, tokens[i], line)
		}
	}
	return aoc.Value()
}

func parseLine(line string) (tokens []NumberToken) {
	tokens = []NumberToken{}
	chars := []rune(line)
	pos := 0
	for {
		if token, found := seekNumber(chars, &pos); found {
			tokens = append(tokens, token)
		} else {
			return
		}
	}
}

func (p *Part1) Calculate(i int, j int, val int) bool {
	if isSymbol(p.lines[i][j]) {
		p.AddValue(val)
		return true // Finding any symbol is sufficient for part 1, return true
	}
	return false
}

func (p *Part2) Calculate(i int, j int, val int) bool {
	if p.lines[i][j] == byte('*') {
		p.AddGear(fmt.Sprintf("%d.%d", i, j), val)
	}
	return false // Always return false, we want to find all possible options for part 2
}

func calculateSymbols(aoc Day3Aoc, input NumberToken, line int) bool {
	startCol := input.Start - 1
	if startCol < 0 {
		startCol = 0
	}
	endCol := input.End + 1

	// Iterate through all characters we might care about:
	// One line above, one line below, one char to the left, one char to the right
	for i := line - 1; i <= line+1; i++ {
		if i < 0 || i >= len(aoc.Lines()) {
			continue
		}
		if endCol == len(aoc.Lines()[i]) {
			endCol--
		}
		for j := startCol; j <= endCol; j++ {
			if aoc.Calculate(i, j, input.Number) {
				return true
			}
		}
	}
	return false
}

func seekNumber(chars []rune, pos *int) (result NumberToken, numberFound bool) {
	// caching variables
	var i int
	var digits []rune
	var positions []int

	for i = *pos; i < len(chars); i++ {
		char := chars[i]
		if isDigit(char) {
			digits = append(digits, char)
			positions = append(positions, i)
		} else if len(digits) > 0 {
			break
		}
	}

	if len(digits) > 0 {
		*pos = i
		number, err := strconv.Atoi(string(digits))
		if err != nil {
			panic(err) // Highly unlikely in these contrived examples!
		}
		result = NumberToken{
			Number: number,
			Start:  positions[0],
			End:    positions[len(positions)-1],
		}
		numberFound = true
	}
	return
}

func isDigit(char rune) bool {
	return char >= rune('0') && char <= rune('9')
}

func isSymbol(char byte) bool {
	return ((char < byte('0') || char > byte('9')) && char != byte('.'))
}
