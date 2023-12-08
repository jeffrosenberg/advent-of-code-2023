package day7

import (
	"sort"
	"strconv"
	"strings"
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

const (
	HAND_HIGHCARD  = '1'
	HAND_PAIR      = '2'
	HAND_TWOPAIR   = '3'
	HAND_THREEKIND = '4'
	HAND_FULLHOUSE = '5'
	HAND_FOURKIND  = '6'
	HAND_FIVEKIND  = '7'
)

var faceMap map[rune]rune = map[rune]rune{
	'T': 'A',
	'J': 'B',
	'Q': 'C',
	'K': 'D',
	'A': 'E',
}

func (h *hand) calculateStrength() {
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

	h.handType = calculateHandType(cardsInHand)
	h.strength = string(h.handType)
	h.strength += handBuilder.String()
}

func calculateHandType(cardsInHand map[rune]int) rune {
	qtys := []int{}
	for _, qty := range cardsInHand {
		qtys = append(qtys, qty)
	}
	sort.Ints(qtys)

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

type Part1 struct {
	lines []string
	value int
	hands []hand
}

type Part2 struct {
	lines []string
	value int
}

func NewPart1(lines []string) *Part1 {
	p := Part1{
		lines: lines,
		hands: []hand{},
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
	for _, line := range p.lines {
		p.hands = append(p.hands, parse(line))
	}
	p.sortByStrength()

	// Once sorted, loop through and multiple bid * rank
	for i, hand := range p.hands {
		p.value += ((i + 1) * hand.bid)
	}
}

func (p *Part1) sortByStrength() {
	sort.Slice(p.hands, func(i, j int) bool { return p.hands[i].strength < p.hands[j].strength })
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

func parse(input string) hand {
	if cards, bid, ok := strings.Cut(input, " "); ok {
		bidInt, err := strconv.Atoi(bid)
		if err != nil {
			panic(err)
		}
		newHand := hand{
			bid:   bidInt,
			cards: cards,
		}

		newHand.calculateStrength()
		return newHand
	}

	// We should remain within the if block above
	panic("unable to parse line")
}
