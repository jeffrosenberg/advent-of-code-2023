package day7

import (
	"sort"
	"strconv"
	"strings"

	"github.com/jeffrosenberg/advent-of-code-2023/go/pkg/aoc"
)

type hand struct {
	cards    string
	bid      int
	handType rune   // see consts below
	strength string // a string used to sort by strength
}

/*
Sorting rules from the puzzle:

Hands are primarily ordered based on type; for example,
every full house is stronger than any three of a kind.

If two hands have the same type, a second ordering rule takes effect.
Start by comparing the first card in each hand.
If these cards are different, the hand with the stronger first card
is considered stronger. If the first card in each hand have the same label,
however, then move on to considering the second card in each hand.
If they differ, the hand with the higher second card wins;
otherwise, continue with the third card in each hand, then the fourth,
then the fifth.

I'm going to solve this by calculating "strength" as a string, as follows:
	- The first character will represent the type, thus going first in sort order
	- Each subsequent character will represent a card
	- Where card values are non-numeric (10-A cards, hand types),
		I'll use an arbitrary number that will sort properly
*/

type solver interface { // Extends aoc.Solver
	addHand(hand)
	calculateHandType(map[rune]int) rune
	faceValues() map[rune]rune
	aoc.Solver
}

const (
	HAND_HIGHCARD  = '1'
	HAND_PAIR      = '2'
	HAND_TWOPAIR   = '3'
	HAND_THREEKIND = '4'
	HAND_FULLHOUSE = '5'
	HAND_FOURKIND  = '6'
	HAND_FIVEKIND  = '7'
)

func (h *hand) calculateStrength(faceMap map[rune]rune, calcFunc func(map[rune]int) rune) {
	var handBuilder strings.Builder
	cardsInHand := make(map[rune]int, 5)

	for _, c := range h.cards {
		// Card strength
		if val, ok := faceMap[c]; ok {
			handBuilder.WriteRune(val)
		} else {
			handBuilder.WriteRune(c)
		}

		// Hand type
		cardsInHand[c] += 1
	}

	h.handType = calcFunc(cardsInHand)
	h.strength = string(h.handType)
	h.strength += handBuilder.String()
}

type Part1 struct {
	lines []string
	value int
	hands []hand
}

type Part2 struct {
	lines []string
	value int
	hands []hand
}

func NewPart1(lines []string) *Part1 {
	p := Part1{
		lines: lines,
		hands: []hand{},
		value: 0,
	}
	return &p
}

func (p *Part1) addHand(h hand) {
	p.hands = append(p.hands, h)
}

func (p *Part1) faceValues() map[rune]rune {
	return map[rune]rune{
		'T': 'A',
		'J': 'B',
		'Q': 'C',
		'K': 'D',
		'A': 'E',
	}
}

func (p *Part1) calculateHandType(cardsInHand map[rune]int) rune {
	qtys := []int{}
	for _, qty := range cardsInHand {
		qtys = append(qtys, qty)
	}
	sort.Ints(qtys)
	return calculateHandType(qtys)
}

func (p *Part1) Lines() []string {
	return p.lines
}

func (p *Part1) Value() int {
	return p.value
}

func (p *Part1) Solve() {
	parse(p)
	sortByStrength(p.hands)

	// Once sorted, loop through and multiple bid * rank
	for i, hand := range p.hands {
		p.value += ((i + 1) * hand.bid)
	}
}

func NewPart2(lines []string) *Part2 {
	p := Part2{
		lines: lines,
		value: 0,
	}
	return &p
}

func (p *Part2) addHand(h hand) {
	p.hands = append(p.hands, h)
}

func (p *Part2) faceValues() map[rune]rune {
	return map[rune]rune{
		'T': 'A',
		'J': '1', // Change from part 1
		'Q': 'C',
		'K': 'D',
		'A': 'E',
	}
}

func (p *Part2) calculateHandType(cardsInHand map[rune]int) rune {
	qtys := []int{}
	var jokers int
	for card, qty := range cardsInHand {
		if card == 'J' {
			// Change from part 1
			// Save any jokers and add them to the greatest remaining qty
			jokers = qty
		} else {
			qtys = append(qtys, qty)
		}
	}
	sort.Ints(qtys)
	if jokers == 5 {
		qtys = append(qtys, 5)
	} else if jokers > 0 {
		qtys[len(qtys)-1] += jokers
	}

	return calculateHandType(qtys)
}

func (p *Part2) Lines() []string {
	return p.lines
}
func (p *Part2) Value() int {
	return p.value
}

func (p *Part2) Solve() {
	parse(p)
	sortByStrength(p.hands)

	// Once sorted, loop through and multiple bid * rank
	for i, hand := range p.hands {
		p.value += ((i + 1) * hand.bid)
	}
}

func parse(solver solver) {
	for _, line := range solver.Lines() {
		if cards, bid, ok := strings.Cut(line, " "); ok {
			bidInt, err := strconv.Atoi(bid)
			if err != nil {
				panic(err)
			}
			newHand := hand{
				bid:   bidInt,
				cards: cards,
			}

			newHand.calculateStrength(solver.faceValues(), solver.calculateHandType)
			solver.addHand(newHand)
		}
	}
}

func calculateHandType(qtys []int) rune {
	switch len(qtys) {
	case 5:
		return HAND_HIGHCARD
	case 4:
		return HAND_PAIR
	case 3:
		if qtys[2] == 2 {
			return HAND_TWOPAIR
		} else {
			return HAND_THREEKIND
		}
	case 2:
		if qtys[1] == 3 {
			return HAND_FULLHOUSE
		} else {
			return HAND_FOURKIND
		}
	case 1:
		return HAND_FIVEKIND
	}

	panic("Didn't match hand type")
}

func sortByStrength(hands []hand) {
	sort.Slice(hands, func(i, j int) bool { return hands[i].strength < hands[j].strength })
}
