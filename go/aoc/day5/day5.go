package day5

import (
	"sort"
	"strings"

	"github.com/jeffrosenberg/advent-of-code-2023/go/pkg/aoc"
)

type Part1 struct {
	lines []string
	value int
	seeds []int
	maps  []map[int]int
	keys  [][]int // cache map keys for sorting
}

type Part2 struct {
	lines []string
	value int
}

func NewPart1(lines []string) *Part1 {
	p := Part1{
		lines: lines,
		value: 0,
		seeds: []int{},
		maps: []map[int]int{
			{0: 0},
		},
		keys: [][]int{{0}},
	}
	return &p
}

func (p *Part1) traverse(seed int) int {
	return p.traverse_step(seed, 0)
}

func (p *Part1) traverse_step(in int, idx int) int {
	if idx >= len(p.maps) {
		return in
	}
	currentMap := p.maps[idx]
	currentKeys := p.keys[idx]
	dest := getDestination(in, currentMap, currentKeys)
	return p.traverse_step(dest, idx+1)
}

func getDestination(in int, currentMap map[int]int, currentKeys []int) int {
	// Iterate through maps until we get to the range that includes our input
	var srcRange int
	sort.Ints(currentKeys)
	for _, key := range currentKeys {
		if key < in {
			srcRange = key
		} else if key == in {
			srcRange = key
			break
		} else {
			// keep the previous value for srcRange
			break
		}
	}

	adjustment := currentMap[srcRange]
	return in + adjustment
}

func (p *Part1) parse() {
	seedLine := p.lines[0]
	if _, nums, ok := strings.Cut(seedLine, ":"); ok {
		for _, num := range strings.Split(nums, " ") {
			if num == " " {
				continue
			}
			if parsedNum, isInt := aoc.ConvertInt(num); isInt {
				p.seeds = append(p.seeds, parsedNum)
			}
		}
	}

	currentMap := -1 // HACK: We're going to add 1 each time, so start at -1
	for i := 1; i < len(p.lines); i++ {
		line := p.lines[i]
		if len(line) == 0 {
			continue
		} else if strings.Contains(line, "map:") {
			p.maps = append(p.maps, map[int]int{})
			p.keys = append(p.keys, []int{})
			currentMap++
			continue
		} else {
			cachedKeys := buildMap(line, p.maps[currentMap], p.keys[currentMap])
			p.keys[currentMap] = cachedKeys
		}
	}
}

func (p *Part1) Value() int {
	return p.value
}

func (p *Part1) Solve() {
	p.parse()
	for _, seed := range p.seeds {
		seedVal := p.traverse(seed)
		if seedVal < p.value || p.value == 0 {
			p.value = seedVal
		}
	}
}

func NewPart2(lines []string) *Part2 {
	p := Part2{
		lines: lines,
		value: 0,
	}
	return &p
}

func (p *Part2) Lines() []string {
	return p.lines
}

func (p *Part2) Value() int {
	return p.value
}

func (p *Part2) AddValue(val int) {
	p.value += val
}

func (p *Part2) Solve() {
	// TODO
}

func buildMap(input string, targetMap map[int]int, targetKeys []int) []int {
	// Input always consists of three parts:
	// 1. destination start
	// 2. source start
	// 3. range length
	params := strings.Split(input, " ")
	dest, _ := aoc.ConvertInt(params[0])
	src, _ := aoc.ConvertInt(params[1])
	rng, _ := aoc.ConvertInt(params[2])

	// Trying to store full ranges in maps results in massive data structures
	// instead, just store the bottom of each range and the adjustment factor
	// (including the end of a range when the factor returns to 0)

	targetMap[src] = dest - src
	targetKeys = append(targetKeys, src)
	if _, exists := targetMap[src+rng]; !exists {
		targetMap[src+rng] = 0
		targetKeys = append(targetKeys, src+rng)
	}
	return targetKeys
}
