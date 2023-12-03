package main

import (
	aoc "github.com/jeffrosenberg/advent-of-code-2023/go/aoc/day3"
	"github.com/jeffrosenberg/advent-of-code-2023/go/pkg/readaoc"
)

// TODO: Make Answer() an interface?
// If the structure of the puzzles remains the same
// (input []lines, output int), then go for it

func main() {
	lines := readaoc.ReadAocInput("inputs/3.txt")
	value := aoc.Answer(lines)
	println(value)
}
