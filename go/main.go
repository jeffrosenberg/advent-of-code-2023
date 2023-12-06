package main

import (
	"fmt"
	"os"

	day1 "github.com/jeffrosenberg/advent-of-code-2023/go/aoc/day1"
	day2 "github.com/jeffrosenberg/advent-of-code-2023/go/aoc/day2"
	day3 "github.com/jeffrosenberg/advent-of-code-2023/go/aoc/day3"
	day4 "github.com/jeffrosenberg/advent-of-code-2023/go/aoc/day4"
	day5 "github.com/jeffrosenberg/advent-of-code-2023/go/aoc/day5"
	"github.com/jeffrosenberg/advent-of-code-2023/go/pkg/aoc"
)

const DAY string = "5"
const PART string = "1"

func main() {
	args := os.Args[1:]
	solver := getSolver(args)
	solver.Solve()
	println(solver.Value())
}

func getSolver(args []string) aoc.Solver {
	var day, part string
	if len(args) > 0 {
		day = args[0]
		if len(args) > 1 {
			part = args[1]
		}
	}
	if day == "" {
		day = DAY
	}
	if part == "" {
		part = PART
	}
	path := fmt.Sprintf("inputs/%s.txt", day)

	// TODO: I'm sure this can be much more elegant,
	// but I don't know how in Go!
	switch {
	case day == "1" && part == "1":
		return day1.NewPart1(aoc.ReadAocInput(path))
	case day == "1" && part == "2":
		return day1.NewPart2(aoc.ReadAocInput(path))
	case day == "2" && part == "1":
		return day2.NewPart1(aoc.ReadAocInput(path))
	case day == "2" && part == "2":
		return day2.NewPart2(aoc.ReadAocInput(path))
	case day == "3" && part == "1":
		return day3.NewPart1(aoc.ReadAocInput(path))
	case day == "3" && part == "2":
		return day3.NewPart2(aoc.ReadAocInput(path))
	case day == "4" && part == "1":
		return day4.NewPart1(aoc.ReadAocInput(path))
	case day == "4" && part == "2":
		return day4.NewPart2(aoc.ReadAocInput(path))
	case day == "5" && part == "1":
		return day5.NewPart1(aoc.ReadAocInput(path))
	case day == "5" && part == "2":
		return day5.NewPart2(aoc.ReadAocInput(path))
	}

	return day5.NewPart1(aoc.ReadAocInput(path))
}
