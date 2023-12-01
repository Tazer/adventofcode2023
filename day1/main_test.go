package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1(t *testing.T) {
	cal := New([]string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	})

	res := cal.Part1()

	assert.Equal(t, 142, res)
}

func TestDay1Part2(t *testing.T) {
	cal := New([]string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	})

	res := cal.Part2()

	assert.Equal(t, 281, res)
}

func TestDay1Part2Test1(t *testing.T) {
	cal := New([]string{
		// "two1nine",
		"eightwothree",
		// "abcone2threexyz",
		// "xtwone3four",
		// "4nineeightseven2",
		// "zoneight234",
		// "7pqrstsixteen",
	})

	res := cal.Part2()

	assert.Equal(t, 83, res)
}

func TestParser(t *testing.T) {
	cal := New([]string{})

	res := cal.checkWord("one", false)
	assert.Equal(t, 1, res)

	res2 := cal.checkWord("eno", true)
	assert.Equal(t, 1, res2)

}
