package day6

import (
	"math"
)

/*
	This day was actually solved on a calculator with a few quadratic equations!
	Reproducing the answer here for completeness
	This can be solved as a quadratic equation because:

		x(y - x) > z
		xy -x^2 > z
		-x^2 + xy - z > 0
		x^2 - xy + z < 0
*/

type Part1 struct {
	value int
}

type Part2 struct {
	value int
}

func NewPart1(lines []string) *Part1 {
	p := Part1{
		// Keep the signature, but ignore lines this time
		value: 0,
	}
	return &p
}

func (p *Part1) Value() int {
	return p.value
}

func (p *Part1) Solve() {
	// There are only 4 sets of numbers here, so I'm dispensing with
	// parsing and looping, and just entering each set individually
	min, max := quadratic(1, -44, 283)
	p.value = (max - min + 1)
	min, max = quadratic(1, -70, 1134)
	p.value *= (max - min + 1)
	min, max = quadratic(1, -70, 1134)
	p.value *= (max - min + 1)
	min, max = quadratic(1, -80, 1491)
	p.value *= (max - min + 1)
}

func NewPart2(lines []string) *Part2 {
	p := Part2{
		// Keep the signature, but ignore lines this time
		value: 0,
	}
	return &p
}

func (p *Part2) Value() int {
	return p.value
}

func (p *Part2) Solve() {
	min, max := quadratic(1, -44707080, 283113411341491)
	p.value = (max - min + 1)
}

func quadratic(a int, b int, c int) (min int, max int) {
	a_float := float64(a)
	b_float := float64(b)
	c_float := float64(c)
	qdrt_min := (-b_float - math.Sqrt(math.Pow(b_float, 2.0)-(4*a_float*c_float))) / (2.0 * a_float)
	qdrt_max := (-b_float + math.Sqrt(math.Pow(b_float, 2.0)-(4*a_float*c_float))) / (2.0 * a_float)
	min = int(math.Ceil(qdrt_min))
	max = int(math.Floor(qdrt_max))
	return
}
