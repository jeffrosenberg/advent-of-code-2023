package day2

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
)

var limits = map[string]int{
	// only 12 red cubes, 13 green cubes, and 14 blue cubes
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Answer(lines []string) (value int) {
	for i := 0; i < len(lines); i++ {
		scanner := bufio.NewScanner(strings.NewReader(lines[i]))
		scanner.Split(bufio.ScanWords)

		gameId, err := getGameId(scanner)
		if err != nil {
			panic(err)
		}

		for {
			color, num, err := getCubeDraw(scanner)
			if err != nil {
				panic(err)
			}
			if color == "" && num == 0 {
				break
			}
			if !isPossible(color, num) {
				gameId = 0 // Impossible input, don't add to final output
				break
			}
		}

		value += gameId
	}
	return
}

func getGameId(s *bufio.Scanner) (int, error) {
	// First token: game
	if !s.Scan() {
		return 0, errors.New("nothing to scan")
	}
	// Second token: game number
	if !s.Scan() {
		return 0, errors.New("nothing to scan")
	}
	numberToken, _ := strings.CutSuffix(s.Text(), ":")
	val, err := strconv.Atoi(numberToken)
	if err != nil {
		return 0, err
	}
	return val, nil
}

func getCubeDraw(s *bufio.Scanner) (string, int, error) {
	// First token: number
	if !s.Scan() {
		return "", 0, nil // EOF, return empty values
	}
	num, err := strconv.Atoi(s.Text())
	if err != nil {
		return "", 0, err
	}

	// Second token: color
	if !s.Scan() {
		return "", 0, errors.New("nothing to scan")
	}
	color, _ := strings.CutSuffix(s.Text(), ",")
	color, _ = strings.CutSuffix(color, ";")
	return color, num, nil
}

func isPossible(color string, quantity int) bool {
	return quantity <= limits[color]
}
