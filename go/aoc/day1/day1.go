package day1

import (
	"strconv"
)

func Answer(lines []string) (value int) {
	for i := 0; i < len(lines); i++ {
		val, err := extractCalibrationValue(lines[i])
		if err != nil {
			panic(err)
		}
		value += val
	}
	return
}

func extractCalibrationValue(input string) (int, error) {
	value := 0
	var digits []int

	// iterate through and store digits
	for i := 0; i < len(input); i++ {
		if digit, isDigit := isDigit((input)[i]); isDigit {
			digits = append(digits, digit)
		}
		if digit, isDigit := isSpelledDigit(input, i); isDigit {
			digits = append(digits, digit)
		}
	}

	// assemble calibration value
	value += digits[0] * 10
	value += digits[len(digits)-1]

	return value, nil
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
