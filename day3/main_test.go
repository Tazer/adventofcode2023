package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	e := NewEngine([]string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	})

	res := e.SumParts()

	assert.Equal(t, 4361, res)

	res2 := e.GearRatios()

	assert.Equal(t, 467835, res2)
}
