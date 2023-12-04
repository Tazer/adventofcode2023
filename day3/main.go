package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
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

	e := NewEngine(inputs)

	res := e.SumParts()

	log.Printf("Part 1: %d", res)

	res2 := e.GearRatios()
	log.Printf("Part 2: %d", res2)

}

type Engine struct {
	schema      map[int]map[int]string
	fullNumbers []FullNumber
}

func NewEngine(input []string) *Engine {
	return &Engine{
		schema: parseSchema(input),
	}
}

func parseSchema(input []string) map[int]map[int]string {
	schema := make(map[int]map[int]string)
	for i, line := range input {
		schema[i] = make(map[int]string)
		for j, char := range line {
			schema[i][j] = string(char)
		}
	}
	return schema
}

func (e *Engine) SumParts() int {

	sum := 0

	for r := 0; r < len(e.schema); r++ {
		for c := 0; c < len(e.schema[r]); c++ {
			item := e.schema[r][c]

			nSum, isNumber, move := e.lookAround(item, r, c)

			c += move
			if !isNumber {
				continue
			}

			sum += nSum

		}
	}

	return sum
}

func (e *Engine) GearRatios() int {

	sum := 0

	for r := 0; r < len(e.schema); r++ {
		for c := 0; c < len(e.schema[r]); c++ {
			item := e.schema[r][c]

			if item == "*" {
				res, add := e.lookForGearRatio(r, c)
				if add {
					sum += res
				}
			}
		}
	}

	return sum
}

func (e *Engine) lookForGearRatio(r, c int) (int, bool) {

	number1 := 0

	number2 := 0

	for _, f := range e.fullNumbers {
		for _, n := range f.Numbers {
			// check up
			if n.Row == r-1 && n.Col == c {
				if number1 == 0 {
					number1 = f.GetFullNumber()
					break
				} else {
					number2 = f.GetFullNumber()
					break
				}
			}

			// check down
			if n.Row == r+1 && n.Col == c {
				if number1 == 0 {
					number1 = f.GetFullNumber()
					break
				} else {
					number2 = f.GetFullNumber()
					break
				}
			}

			// check left
			if n.Row == r && n.Col == c-1 {
				if number1 == 0 {
					number1 = f.GetFullNumber()
					break
				} else {
					number2 = f.GetFullNumber()
					break
				}
			}

			// check right
			if n.Row == r && n.Col == c+1 {
				if number1 == 0 {
					number1 = f.GetFullNumber()
					break
				} else {
					number2 = f.GetFullNumber()
					break
				}
			}

			// check up left
			if n.Row == r-1 && n.Col == c-1 {
				if number1 == 0 {
					number1 = f.GetFullNumber()
					break
				} else {
					number2 = f.GetFullNumber()
					break
				}
			}

			// check up right
			if n.Row == r-1 && n.Col == c+1 {
				if number1 == 0 {
					number1 = f.GetFullNumber()
					break
				} else {
					number2 = f.GetFullNumber()
					break
				}
			}

			// check down left
			if n.Row == r+1 && n.Col == c-1 {
				if number1 == 0 {
					number1 = f.GetFullNumber()
					break
				} else {
					number2 = f.GetFullNumber()
					break
				}
			}

			// check down right
			if n.Row == r+1 && n.Col == c+1 {
				if number1 == 0 {
					number1 = f.GetFullNumber()
					break
				} else {
					number2 = f.GetFullNumber()
					break
				}
			}
		}
	}

	if number1 != 0 && number2 != 0 {
		return number1 * number2, true
	}

	return 0, false
}

func (e *Engine) lookAround(item string, r, c int) (int, bool, int) {
	// find the full number
	_, err := strconv.Atoi(item)

	if err != nil {
		return 0, false, 0
	}

	f := FullNumber{
		Numbers: []Number{
			{
				Number: item,
				Row:    r,
				Col:    c,
			},
		},
	}

	for col := c + 1; col < len(e.schema[r]); col++ {
		item := e.schema[r][col]

		_, err := strconv.Atoi(item)

		if err != nil {
			break
		}

		f.Numbers = append(f.Numbers, Number{
			Number: item,
			Row:    r,
			Col:    col,
		})
	}

	e.fullNumbers = append(e.fullNumbers, f)

	// see if there is any symbols that are not . around.
	sum, hasThingsClose := e.lookAtFullNumber(f)

	return sum, hasThingsClose, len(f.Numbers) - 1

}

func (e *Engine) lookAtFullNumber(f FullNumber) (int, bool) {

	found := false

	for _, n := range f.Numbers {
		// check up
		up := e.schema[n.Row-1][n.Col]
		if e.check(up) {
			found = true
			break
		}

		// check down
		down := e.schema[n.Row+1][n.Col]
		if e.check(down) {
			found = true
			break
		}

		// check left
		left := e.schema[n.Row][n.Col-1]
		if e.check(left) {
			found = true
			break
		}

		// check right
		right := e.schema[n.Row][n.Col+1]
		if e.check(right) {
			found = true
			break
		}

		// check up left
		upLeft := e.schema[n.Row-1][n.Col-1]
		if e.check(upLeft) {
			found = true
			break
		}

		// check up right
		upRight := e.schema[n.Row-1][n.Col+1]
		if e.check(upRight) {
			found = true
			break
		}

		// check down left
		downLeft := e.schema[n.Row+1][n.Col-1]
		if e.check(downLeft) {
			found = true
			break
		}

		// check down right
		downRight := e.schema[n.Row+1][n.Col+1]
		if e.check(downRight) {
			found = true
			break
		}

	}

	if found {
		return f.GetFullNumber(), true
	}

	return 0, false

}

func (e *Engine) check(v string) bool {

	if v == "" {
		return false
	}

	_, err := strconv.Atoi(v)

	if err == nil {
		return false
	}

	if v == "." {
		return false
	}

	return true
}

func (f *FullNumber) GetFullNumber() int {
	num := ""
	for _, n := range f.Numbers {
		num += n.Number
	}

	res, _ := strconv.Atoi(num)

	return res
}

type FullNumber struct {
	Numbers []Number
}

type Number struct {
	Number string
	Row    int
	Col    int
}
