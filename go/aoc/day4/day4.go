package day4

import (
	"strconv"
	"strings"
)

type Card struct {
	cardName string
	winners  map[int]bool
	given    []int
	matches  int
	copies   int
}

type Part1 struct {
	lines []string
	cards []Card
}

type Part2 struct {
	lines []string
	cards []Card
}

func NewPart1(lines []string) *Part1 {
	p := Part1{
		lines: lines,
	}
	return &p
}

func (p *Part1) Value() int {
	val := 0
	for _, card := range p.cards {
		if card.matches == 0 {
			continue
		}
		val += power(2, (card.matches - 1))
	}
	return val
}

func (p *Part1) Solve() {
	for _, line := range p.lines {
		card := parse(line)
		card.calculateMatches()
		p.cards = append(p.cards, card)
	}
}

func NewPart2(lines []string) *Part2 {
	p := Part2{
		lines: lines,
	}
	return &p
}

func (p *Part2) Value() int {
	var val int = 0
	for _, card := range p.cards {
		val += card.copies
	}
	return val
}

func (p *Part2) Solve() {
	for _, line := range p.lines {
		card := parse(line)
		card.calculateMatches()
		p.cards = append(p.cards, card)
	}
	p.generateCopies()
}

// TODO: better solution using recursion?
func (p *Part2) generateCopies() {
	for i := 0; i < len(p.cards); i++ {
		p.cards[i].copies++ // 1 copy to denote self
		if p.cards[i].matches > 0 {
			// Add copies to n subsequent cards
			for j := i + 1; j < (i+1)+p.cards[i].matches && j < len(p.cards); j++ {
				p.cards[j].copies += p.cards[i].copies
			}
		}
	}
}

func (card *Card) calculateMatches() {
	for _, g := range card.given {
		if _, match := card.winners[g]; match {
			card.winners[g] = true
			card.matches++
		}
	}
}

func parse(line string) Card {
	output := Card{
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
