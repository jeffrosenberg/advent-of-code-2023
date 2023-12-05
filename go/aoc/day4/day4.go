package day4

import (
	"strconv"
	"strings"
)

type Line struct {
	cardName string
	winners  map[int]bool
	given    []int
	matches  int
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

func (p *Part1) Value() int {
	return p.value
}

func (p *Part1) AddValue(val int) {
	p.value += val
}

func (p *Part1) Solve() {
	for _, line := range p.lines {
		parsedLine := parse(line)
		parsedLine.CalculateMatches()
		if parsedLine.matches == 0 {
			continue
		}
		p.AddValue(power(2, (parsedLine.matches - 1)))
	}
}

func NewPart2(lines []string) *Part2 {
	p := Part2{
		lines: lines,
		value: 0,
	}
	return &p
}

func (p *Part2) Value() int {
	return p.value
}

func (p *Part2) AddValue(val int) {
	p.value += val
}

func (p *Part2) Solve() {

}

func (line *Line) CalculateMatches() {
	for _, g := range line.given {
		if _, match := line.winners[g]; match {
			line.winners[g] = true
			line.matches++
		}
	}
}

func parse(line string) Line {
	output := Line{
		winners: map[int]bool{},
		given:   []int{},
	}
	if name, contents, success := strings.Cut(line, ":"); success {
		output.cardName = name
		if winners, given, success := strings.Cut(contents, "|"); success {
			for _, token := range strings.Split(winners, " ") {
				if winner, isInt := convertInt(token); isInt {
					output.winners[winner] = false // Mark all false until a match is made
				}
			}
			for _, token := range strings.Split(given, " ") {
				if given, isInt := convertInt(token); isInt {
					output.given = append(output.given, given)
				}
			}
			return output
		}
	}
	panic("Line in unexpected format!")
}

func convertInt(val string) (int, bool) {
	if val == "" {
		return 0, false
	}
	output, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return output, true
}

func power(n, m int) int { // a power function that accepts and returns ints
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}
