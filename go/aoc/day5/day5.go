package day5

import (
	"fmt"
	"sort"
	"strings"

	"github.com/jeffrosenberg/advent-of-code-2023/go/pkg/aoc"
)

type day5 struct {
	lines []string
	value int
	maps  []map[int]int
	keys  [][]int // cache map keys for sorting
}

type Part1 struct {
	seeds []int
	d     *day5
}

type seedRange struct {
	start int
	end   int
}

type Part2 struct {
	seeds []seedRange
	d     *day5
}

func NewPart1(lines []string) *Part1 {
	p := Part1{
		seeds: []int{},
		d: &day5{
			lines: lines,
			value: 0,
			maps: []map[int]int{
				{0: 0},
			},
			keys: [][]int{{0}},
		},
	}
	return &p
}

func (d *day5) traverse(seed int) int {
	return d.traverse_step(seed, 0)
}

func (d *day5) traverse_step(in int, idx int) int {
	if idx >= len(d.maps) {
		return in
	}
	currentMap := d.maps[idx]
	currentKeys := d.keys[idx]
	dest := getDestination(in, currentMap, currentKeys)
	return d.traverse_step(dest, idx+1)
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

func (p *Part1) parseSeedLine() {
	seedLine := p.d.lines[0]
	if _, nums, ok := strings.Cut(seedLine, ": "); ok {
		for _, num := range strings.Split(nums, " ") {
			if parsedNum, isInt := aoc.ConvertInt(num); isInt {
				p.seeds = append(p.seeds, parsedNum)
			}
		}
	}
}

func (d *day5) parseMaps() {
	currentMap := -1 // HACK: We're going to add 1 each time, so start at -1
	for i := 1; i < len(d.lines); i++ {
		line := d.lines[i]
		if len(line) == 0 {
			continue
		} else if strings.Contains(line, "map:") {
			d.maps = append(d.maps, map[int]int{})
			d.keys = append(d.keys, []int{})
			currentMap++
			continue
		} else {
			cachedKeys := buildMap(line, d.maps[currentMap], d.keys[currentMap])
			d.keys[currentMap] = cachedKeys
		}
	}
}

func (d *day5) Lines() []string {
	return d.lines
}

func (p *Part1) Lines() []string {
	return p.d.Lines()
}
func (p *Part1) Value() int {
	return p.d.value
}

func (p *Part1) Solve() {
	p.parseSeedLine()
	p.d.parseMaps()
	for _, seed := range p.seeds {
		seedVal := p.d.traverse(seed)
		if seedVal < p.d.value || p.d.value == 0 {
			p.d.value = seedVal
		}
	}
}

func NewPart2(lines []string) *Part2 {
	p := Part2{
		seeds: []seedRange{},
		d: &day5{
			lines: lines,
			value: 0,
			maps: []map[int]int{
				{0: 0},
			},
			keys: [][]int{{0}},
		},
	}
	return &p
}

func (p *Part2) parseSeedLine() {
	seedLine := p.d.lines[0]
	if _, nums, ok := strings.Cut(seedLine, ": "); ok {
		var startNum int
		for i, num := range strings.Split(nums, " ") {
			if i%2 == 0 {
				if parsedNum, isInt := aoc.ConvertInt(num); isInt {
					startNum = parsedNum
				}
			} else {
				if rng, isInt := aoc.ConvertInt(num); isInt {
					p.seeds = append(p.seeds, seedRange{start: startNum, end: startNum + rng})
				}
			}
		}
	}
}

func (p *Part2) Lines() []string {
	return p.d.Lines()
}
func (p *Part2) Value() int {
	return p.d.value
}

// This is crazy slow.
// Given the large numbers involved, I don't know how to make it fast,
// although presumably there's a way to do it...
// Most likely I wasn't supposed to traverse all of these, but figure out a shortcut?
func (p *Part2) Solve() {
	p.parseSeedLine()
	p.d.parseMaps()
	for _, rng := range p.seeds {
		println(fmt.Sprintf("Processing %d - %d range...", rng.start, rng.end))
		for seed := rng.start; seed < rng.end; seed++ {
			seedVal := p.d.traverse(seed)
			if seedVal < p.d.value || p.d.value == 0 {
				p.d.value = seedVal
			}
		}
	}
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
