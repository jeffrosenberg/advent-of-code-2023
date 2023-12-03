package day3

import (
	"strconv"
)

type NumberToken struct {
	Number int
	Start  int
	End    int
}

func Answer(lines []string) (value int) {
	for line := 0; line < len(lines); line++ {
		tokens := parseLine(lines[line])
		for i := 0; i < len(tokens); i++ {
			if hasAdjacentSymbols(tokens[i], line, lines) {
				// println(fmt.Sprintf("Line %d : Number %d", line, tokens[i].Number))
				value += tokens[i].Number
			}
		}
	}
	return
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

func hasAdjacentSymbols(input NumberToken, line int, lines []string) bool {
	startCol := input.Start - 1
	if startCol < 0 {
		startCol = 0
	}
	endCol := input.End + 1

	// Iterate through all characters we might care about:
	// One line above, one line below, one char to the left, one char to the right
	for i := line - 1; i <= line+1; i++ {
		if i < 0 || i >= len(lines) {
			continue
		}
		if endCol == len(lines[i]) {
			endCol--
		}
		for j := startCol; j <= endCol; j++ {
			if isSymbol(lines[i][j]) {
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
