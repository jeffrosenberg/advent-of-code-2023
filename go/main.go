package main

import (
	day3 "github.com/jeffrosenberg/advent-of-code-2023/go/aoc/day3"
	"github.com/jeffrosenberg/advent-of-code-2023/go/pkg/readaoc"
)

// TODO: Make Answer() an interface?
// If the structure of the puzzles remains the same
// (input []lines, output int), then go for it

func main() {
	var puzzle day3.Day3Aoc = day3.NewPart2(readaoc.ReadAocInput("inputs/3.txt"))
	value := day3.Answer(puzzle)
	println(value)
}
