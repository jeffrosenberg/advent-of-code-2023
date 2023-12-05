package main

import (
	day4 "github.com/jeffrosenberg/advent-of-code-2023/go/aoc/day4"
	"github.com/jeffrosenberg/advent-of-code-2023/go/pkg/readaoc"
)

type Puzzle interface {
	Lines() []string
	Solve()
	Value() int
}

func main() {
	var puzzle Puzzle = day4.NewPart1(readaoc.ReadAocInput("inputs/4.txt"))
	puzzle.Solve()
	println(puzzle.Value())
}
