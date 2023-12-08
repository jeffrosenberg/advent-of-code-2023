package aoc

import (
	"bufio"
	"os"
	"strconv"
)

type Solver interface {
	Solve()
	Value() int
}

func ReadAocInput(path string) (lines []string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return
}

func ConvertInt(val string) (int, bool) {
	if val == "" {
		return 0, false
	}
	output, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return output, true
}

func Min(a, b int) int {
	// This is a built-in in Go 1.21,
	// but my OS currently ships Go 1.20 :-(

	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	// This is a built-in in Go 1.21,
	// but my OS currently ships Go 1.20 :-(

	if a > b {
		return a
	}
	return b
}
