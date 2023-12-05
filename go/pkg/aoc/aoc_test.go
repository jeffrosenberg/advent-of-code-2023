package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadAocInput(t *testing.T) {
	lines := ReadAocInput("../../../inputs/1.txt")
	assert.Equal(t, 1000, len(lines))
	assert.Equal(t, "2xjzgsjzfhzhm1", lines[0])
	assert.Equal(t, "6ninexrxsvlmmzrsevenjhzzggfcxqrvfjtnjctveight9", lines[422])
	assert.Equal(t, "9lgmxktj1frxl", lines[999])
}
