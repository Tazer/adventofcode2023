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

	s := NewStack(inputs)

	log.Printf("Part 1: %d", s.Points())
	log.Printf("Part 2: %d", s.Points2())

}

func NewStack(input []string) *Stack {
	return &Stack{
		Cards: parseCards(input),
	}
}

func parseCards(input []string) []Card {
	cards := []Card{}
	for _, line := range input {

		c := parseCard(line)

		cards = append(cards, c)
	}
	return cards
}

func parseCard(line string) Card {

	card := Card{}
	firstSplit := strings.Split(line, ":")

	name := firstSplit[0]

	card.Name = name

	indexName := strings.ReplaceAll(strings.ReplaceAll(name, "  ", " "), "  ", " ")
	indexVal := strings.Split(indexName, " ")[1]
	card.Index, _ = strconv.Atoi(indexVal)
	card.Index--

	secondSplit := strings.Split(firstSplit[1], "|")
	winningNumberSplit := strings.Split(secondSplit[0], " ")
	playedNumberSplit := strings.Split(secondSplit[1], " ")

	winnings := []int{}
	for _, w := range winningNumberSplit {
		if w == "" {
			continue
		}
		n, _ := strconv.Atoi(w)
		winnings = append(winnings, n)
	}
	played := []int{}
	for _, p := range playedNumberSplit {
		if p == "" {
			continue
		}
		n, _ := strconv.Atoi(p)
		played = append(played, n)
	}

	card.Winnings = winnings
	card.Played = played

	return card
}

type Stack struct {
	Cards  []Card
	Cards2 []Card
}

func (s *Stack) Points() int {
	totalPoints := 0
	for _, c := range s.Cards {
		totalPoints += c.Points()
	}
	return totalPoints
}

func (s *Stack) Points2() int {
	s.Cards2 = append(s.Cards2, s.Cards...)
	for i, c := range s.Cards {
		if c.Points() == 0 {
			continue
		}

		match := c.Matches()
		if match > len(s.Cards2)-1 {
			match = len(s.Cards2) - 1
		}
		c.PointsWithSubCards(s.Cards2[i+1:i+1+match], s)
	}
	return len(s.Cards2)
}

type Card struct {
	Name     string
	Index    int
	Winnings []int
	Played   []int
}

func (c *Card) PointsWithSubCards(cards []Card, s *Stack) {
	for _, c := range cards {
		if c.Points() == 0 {
			continue
		}
		match := c.Matches()
		if match > len(s.Cards[c.Index+1:])-1 {
			match = len(s.Cards[c.Index+1:]) - 1
		}
		c.PointsWithSubCards(s.Cards[c.Index+1:c.Index+1+match], s)
	}
	s.Cards2 = append(s.Cards2, cards...)
}

func (c *Card) Points() int {
	match := 0
	for _, p := range c.Played {
		for _, w := range c.Winnings {
			if p == w {
				if match == 0 {
					match = 1
				} else {
					match = match * 2
				}

			}
		}
	}
	return match
}
func (c *Card) Matches() int {
	match := 0
	for _, p := range c.Played {
		for _, w := range c.Winnings {
			if p == w {
				match++
			}
		}
	}
	return match
}
