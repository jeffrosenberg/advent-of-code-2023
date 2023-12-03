package main

import (
	day2 "github.com/jeffrosenberg/advent-of-code-2023/go/aoc/day2"
	"github.com/jeffrosenberg/advent-of-code-2023/go/pkg/readaoc"
)

// TODO: Make Answer() an interface?
// If the structure of the puzzles remains the same
// (input []lines, output int), then go for it

func main() {
	lines := readaoc.ReadAocInput("inputs/2.txt")
	value := day2.Answer(lines)
	println(value)
}
