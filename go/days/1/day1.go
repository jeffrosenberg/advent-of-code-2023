package main

import (
	"strconv"

	"github.com/jeffrosenberg/advent-of-code-2023/go/pkg/readaoc"
)

func main() {
	value := 0
	lines := readaoc.ReadAocInput("inputs/1.txt")
	for i := 0; i < len(lines); i++ {
		val, err := extractCalibrationValue(lines[i])
		if err != nil {
			panic(err)
		}
		value += val
	}
	println(value)
}

func extractCalibrationValue(input string) (int, error) {
	value := 0
	chars := []rune(input)

	// first number - iterate forwards
	for i := 0; i < len(chars); i++ {
		if isDigit(chars[i]) {
			digit, err := strconv.Atoi(string(chars[i]))
			if err != nil {
				return 0, err
			}
			value += digit * 10
			break
		}
	}

	// second number - iterate backwards
	for i := len(chars) - 1; i >= 0; i-- {
		if isDigit(chars[i]) {
			digit, err := strconv.Atoi(string(chars[i]))
			if err != nil {
				return 0, err
			}
			value += digit
			break
		}
	}

	return value, nil
}

func isDigit(char rune) bool {
	byt := byte(char)
	return byt >= byte('0') && byt <= byte('9')
}
