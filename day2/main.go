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

	w := NewWorld(inputs, 12, 13, 14)

	log.Printf("World: %d", w.Play())
	log.Printf("World2: %d", w.Play2())
}

func NewWorld(inputs []string, maxRed, maxGreen, maxBlue int) *World {
	return &World{
		maxRed:   maxRed,
		maxBlue:  maxBlue,
		maxGreen: maxGreen,
		Games:    parseGames(inputs),
	}
}

func (w *World) Play() int {
	possibleGames := []Game{}
	for _, g := range w.Games {
		isPossible := true
		for _, r := range g.Rounds {
			for _, p := range r.Plays {
				switch p.Color {
				case "red":
					if p.Number > w.maxRed {
						isPossible = false
						break
					}
				case "blue":
					if p.Number > w.maxBlue {
						isPossible = false
						break
					}
				case "green":
					if p.Number > w.maxGreen {
						isPossible = false
						break
					}
				}
			}
		}
		if isPossible {
			possibleGames = append(possibleGames, g)
		}
	}

	gIDsum := 0
	for _, g := range possibleGames {
		gIDsum += g.ID
	}
	return gIDsum

}

func (w *World) Play2() int {
	games := []int{}
	for _, g := range w.Games {
		redHigh := 0
		blueHigh := 0
		greenHigh := 0
		for _, r := range g.Rounds {
			for _, p := range r.Plays {
				switch p.Color {
				case "red":
					if p.Number > redHigh {
						redHigh = p.Number
					}
				case "blue":
					if p.Number > blueHigh {
						blueHigh = p.Number
					}
				case "green":
					if p.Number > greenHigh {
						greenHigh = p.Number
					}
				}
			}
		}
		games = append(games, redHigh*blueHigh*greenHigh)
	}

	gIDsum := 0
	for _, g := range games {
		gIDsum += g
	}
	return gIDsum

}

type World struct {
	maxRed   int
	maxBlue  int
	maxGreen int
	Games    []Game
}

type Game struct {
	ID     int
	Rounds []Round
}

type Play struct {
	Number int
	Color  string
}

type Round struct {
	Plays []Play
}

func parseGames(inputs []string) []Game {
	games := []Game{}

	for _, v := range inputs {
		games = append(games, parseGame(v))
	}

	return games
}

func parseGame(input string) Game {
	g := Game{}
	gSplit := strings.Split(input, ":")

	gSplitID := strings.Split(gSplit[0], " ")[1]

	g.ID, _ = strconv.Atoi(gSplitID)

	g.Rounds = parseRounds(gSplit[1])

	return g
}

func parseRounds(input string) []Round {
	rounds := []Round{}

	rSplit := strings.Split(input, ";")

	for _, v := range rSplit {
		rounds = append(rounds, parseRound(v))
	}

	return rounds
}

func parseRound(input string) Round {
	plays := []Play{}
	round := Round{}

	pSplit := strings.Split(input, ",")

	for _, v := range pSplit {
		plays = append(plays, parsePlay(v))
	}

	round.Plays = plays

	return round
}

func parsePlay(input string) Play {
	p := Play{}

	pSplit := strings.Split(strings.TrimLeft(input, " "), " ")

	p.Number, _ = strconv.Atoi(pSplit[0])
	p.Color = pSplit[1]

	return p
}
