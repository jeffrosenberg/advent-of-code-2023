package main

import (
	day1 "github.com/jeffrosenberg/advent-of-code-2023/go/aoc/day1"
	"github.com/jeffrosenberg/advent-of-code-2023/go/pkg/readaoc"
)

// TODO: Make Answer() an interface?
// If the structure of the puzzles remains the same
// (input []lines, output int), then go for it

func main() {
	lines := readaoc.ReadAocInput("inputs/1.txt")
	value := day1.Answer(lines)
	println(value)
}
