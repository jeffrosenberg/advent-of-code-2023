package day1

import (
	"strconv"
)

type Part1 struct {
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
	for _, line := range p.lines {
		p.value += extractCalibrationValue(line, false)
	}
}

type Part2 struct {
	lines []string
	value int
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
	for _, line := range p.lines {
		p.value += extractCalibrationValue(line, true)
	}
}

func extractCalibrationValue(input string, includeSpelled bool) int {
	value := 0
	var digits []int

	// iterate through and store digits
	for i := 0; i < len(input); i++ {
		if digit, isDigit := isDigit((input)[i]); isDigit {
			digits = append(digits, digit)
		}
		if includeSpelled {
			if digit, isDigit := isSpelledDigit(input, i); isDigit {
				digits = append(digits, digit)
			}
		}
	}

	// assemble calibration value
	value += digits[0] * 10
	value += digits[len(digits)-1]

	return value
}

func isDigit(char byte) (int, bool) {
	if char >= byte('1') && char <= byte('9') {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err) // Highly unlikely in these contrived examples!
		}
		return digit, true
	}
	return 0, false
}

func isSpelledDigit(input string, index int) (digit int, result bool) {
	chars := input[index:]
	numChars := len(input) - index

	// this is brute-forced and not super elegant,
	// but I'm not convinced making it fancier is any better
	switch {
	case numChars < 3:
		return 0, false
	case chars[:3] == "one":
		return 1, true
	case chars[:3] == "two":
		return 2, true
	case numChars >= 5 && chars[:5] == "three":
		return 3, true
	case numChars >= 4 && chars[:4] == "four":
		return 4, true
	case numChars >= 4 && chars[:4] == "five":
		return 5, true
	case chars[:3] == "six":
		return 6, true
	case numChars >= 5 && chars[:5] == "seven":
		return 7, true
	case numChars >= 5 && chars[:5] == "eight":
		return 8, true
	case numChars >= 4 && chars[:4] == "nine":
		return 9, true
	}
	return 0, false
}
