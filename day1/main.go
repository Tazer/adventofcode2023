package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputs := []string{}

	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	cal := New(inputs)
	res := cal.Part1()

	log.Printf("Part 1: %d", res)

	res2 := cal.Part2()

	log.Printf("Part 2: %d", res2)
}

func New(input []string) *Calibration {
	return &Calibration{
		input: input,
	}
}

type Calibration struct {
	input []string
}

// func (c *Calibration) ReplaceInput() {
// 	for i, _ := range c.input {
// 		c.input[i] = strings.ReplaceAll(c.input[i], "one", "1")
// 		c.input[i] = strings.ReplaceAll(c.input[i], "two", "2")
// 		c.input[i] = strings.ReplaceAll(c.input[i], "three", "3")
// 		c.input[i] = strings.ReplaceAll(c.input[i], "four", "4")
// 		c.input[i] = strings.ReplaceAll(c.input[i], "five", "5")
// 		c.input[i] = strings.ReplaceAll(c.input[i], "six", "6")
// 		c.input[i] = strings.ReplaceAll(c.input[i], "seven", "7")
// 		c.input[i] = strings.ReplaceAll(c.input[i], "eight", "8")
// 		c.input[i] = strings.ReplaceAll(c.input[i], "nine", "9")
// 	}
// }

func (c *Calibration) Part1() int {

	val := 0

	for _, v := range c.input {
		if v == "" {
			continue
		}

		first := ""
		last := ""
		for _, r := range v {
			_, err := strconv.Atoi(string(r))

			if err != nil {
				continue
			}

			if first == "" {
				first = string(r)
				continue
			}

			last = string(r)
		}

		if last == "" {
			last = first
		}

		res := first + last

		iVal, err := strconv.Atoi(res)

		if err != nil {
			continue
		}

		val += iVal

	}

	return val
}

func (c *Calibration) Part2() int {

	val := 0

	for _, v := range c.input {
		if v == "" {
			continue
		}

		first := ""

		word := ""

		for _, r := range v {
			_, err := strconv.Atoi(string(r))

			if err != nil {
				word += string(r)
				wordRes := c.checkWord(word, false)
				if wordRes != 0 {
					first = strconv.Itoa(wordRes)
					break
				}
				continue
			}

			if first == "" {
				first = string(r)
				break
			}
		}
		word = ""
		last := ""

		for i := len(v) - 1; i >= 0; i-- {
			r := v[i]
			_, err := strconv.Atoi(string(r))

			if err != nil {
				word += string(r)
				wordRes := c.checkWord(word, true)
				if wordRes != 0 {
					last = strconv.Itoa(wordRes)
					break
				}
				continue
			}

			if last == "" {
				last = string(r)
				break
			}
		}

		if last == "" {
			last = first
		}

		res := first + last

		iVal, err := strconv.Atoi(res)

		if err != nil {
			continue
		}

		val += iVal

	}

	return val
}

func (c *Calibration) checkWord(word string, reverse bool) int {

	if reverse {
		newWord := ""
		for i := len(word) - 1; i >= 0; i-- {
			newWord += string(word[i])
		}
		word = newWord
	}

	if strings.Contains(word, "one") {
		return 1
	}

	if strings.Contains(word, "two") {
		return 2
	}
	if strings.Contains(word, "three") {
		return 3
	}
	if strings.Contains(word, "four") {
		return 4
	}

	if strings.Contains(word, "five") {
		return 5
	}

	if strings.Contains(word, "six") {
		return 6
	}

	if strings.Contains(word, "seven") {
		return 7
	}

	if strings.Contains(word, "eight") {
		return 8
	}

	if strings.Contains(word, "nine") {
		return 9
	}

	return 0

}
