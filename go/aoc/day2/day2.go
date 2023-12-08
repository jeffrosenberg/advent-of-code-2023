package day2

import (
	"bufio"
	"errors"
	"strconv"
	"strings"

	"github.com/jeffrosenberg/advent-of-code-2023/go/pkg/aoc"
)

type Solver interface { // Extends aoc.Solver
	GetGameId(*bufio.Scanner)
	UpdateValue(func()) int
	aoc.Solver
}

type Part1 struct {
	lines  []string
	value  int
	gameId int
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

func (p *Part1) Solve() {
	for _, line := range p.lines {
		scanner := bufio.NewScanner(strings.NewReader(line))
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

		p.value += gameId
	}
}

func (p *Part1) GetGameId(s *bufio.Scanner) {
	id, err := getGameId(s)
	if err != nil {
		panic(err)
	}
	p.gameId = id
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

func (p *Part2) Value() int {
	return p.value
}

func (p *Part2) Solve() {
	for _, line := range p.lines {
		scanner := bufio.NewScanner(strings.NewReader(line))
		scanner.Split(bufio.ScanWords)

		// We don't care about Game ID anymore,
		// but we still need to parse that part of the input
		getGameId(scanner)

		// Cache the max observed value for each
		var max = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for {
			color, num, err := getCubeDraw(scanner)
			if err != nil {
				panic(err)
			}
			if color == "" && num == 0 {
				break
			}
			max[color] = aoc.Max(max[color], num)
		}

		p.value += (max["red"] * max["green"] * max["blue"])
	}
}

func (p *Part2) GetGameId(s *bufio.Scanner) {
	// We don't care about Game ID anymore,
	// but we still need to parse that part of the input
	_, err := getGameId(s)
	if err != nil {
		panic(err)
	}
}

// func Answer(lines []string) (value int) {
// 	p := NewPart1(lines)
// 	for i := 0; i < len(lines); i++ {
// 		scanner := bufio.NewScanner(strings.NewReader(lines[i]))
// 		scanner.Split(bufio.ScanWords)

// 		// We don't care about Game ID anymore,
// 		// but we still need to parse that part of the input
// 		p.GetGameId(scanner)

// 		// Cache the max observed value for each
// 		var max = map[string]int{
// 			"red":   0,
// 			"green": 0,
// 			"blue":  0,
// 		}

// 		for {
// 			color, num, err := getCubeDraw(scanner)
// 			if err != nil {
// 				panic(err)
// 			}
// 			if color == "" && num == 0 {
// 				break
// 			}
// 			max[color] = int(math.Max(float64(max[color]), float64(num)))
// 		}

// 		value += (max["red"] * max["green"] * max["blue"])
// 	}
// 	return
// }

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
	var limits = map[string]int{
		// only 12 red cubes, 13 green cubes, and 14 blue cubes
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	return quantity <= limits[color]
}
