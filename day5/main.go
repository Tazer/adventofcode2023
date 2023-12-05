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

	g := NewGarden(inputs)

	log.Printf("Part 1: %d", g.LowestLocationNumber(false))
	log.Printf("Part 2: %d", g.LowestLocationNumber(true))
}

type Garden struct {
	Seeds                 []int
	Seeds2                []int
	SeedsToSoil           []Lookup
	SoilToFertilizer      []Lookup
	FertilizerToWater     []Lookup
	WaterToLight          []Lookup
	LightToTemperature    []Lookup
	TemperatureToHumidity []Lookup
	HumidityToLocation    []Lookup
}

func NewGarden(input []string) *Garden {

	return parseGarden(input)
}

func parseGarden(input []string) *Garden {
	g := &Garden{}

	currentMap := ""

	for _, line := range input {
		if strings.Contains(line, "seeds:") {
			g.Seeds = parseSeeds(line)
			g.Seeds2 = parseSeeds2(line)
			continue
		}

		if line == "" {
			continue
		}

		switch line {
		case "seed-to-soil map:":
			currentMap = "seeds-to-soil"
			continue
		case "soil-to-fertilizer map:":
			currentMap = "soil-to-fertilizer"
			continue
		case "fertilizer-to-water map:":
			currentMap = "fertilizer-to-water"
			continue
		case "water-to-light map:":
			currentMap = "water-to-light"
			continue
		case "light-to-temperature map:":
			currentMap = "light-to-temperature"
			continue
		case "temperature-to-humidity map:":
			currentMap = "temperature-to-humidity"
			continue
		case "humidity-to-location map:":
			currentMap = "humidity-to-location"
			continue
		}

		switch currentMap {
		case "seeds-to-soil":
			g.SeedsToSoil = append(g.SeedsToSoil, generateMap(line))
		case "soil-to-fertilizer":
			g.SoilToFertilizer = append(g.SoilToFertilizer, generateMap(line))
		case "fertilizer-to-water":
			g.FertilizerToWater = append(g.FertilizerToWater, generateMap(line))
		case "water-to-light":
			g.WaterToLight = append(g.WaterToLight, generateMap(line))
		case "light-to-temperature":
			g.LightToTemperature = append(g.LightToTemperature, generateMap(line))
		case "temperature-to-humidity":
			g.TemperatureToHumidity = append(g.TemperatureToHumidity, generateMap(line))
		case "humidity-to-location":
			g.HumidityToLocation = append(g.HumidityToLocation, generateMap(line))
		}

	}
	return g
}

type Lookup struct {
	Min  int
	Max  int
	Diff int
}

func generateMap(line string) Lookup {
	sline := strings.Split(line, " ")

	iteration, err := strconv.Atoi(sline[2])
	if err != nil {
		log.Fatal(err)
	}
	destination, _ := strconv.Atoi(sline[0])
	if err != nil {
		log.Fatal(err)
	}
	source, _ := strconv.Atoi(sline[1])
	if err != nil {
		log.Fatal(err)
	}

	diff := destination - source

	l := Lookup{
		Min:  source,
		Max:  source + iteration - 1,
		Diff: diff,
	}

	return l
}

func diffa(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func parseSeeds(line string) []int {
	seeds := []int{}
	for _, s := range strings.Split(line, " ") {
		if strings.Contains(s, "seeds:") {
			continue
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		seeds = append(seeds, i)
	}
	return seeds
}

func parseSeeds2(line string) []int {
	seeds := []int{}

	sSplit := strings.Split(line, " ")

	for i := 0; i < len(sSplit); {
		if sSplit[i] == "seeds:" {
			i++
			continue
		}

		iVal, err := strconv.Atoi(sSplit[i+1])
		if err != nil {
			log.Fatal(err)
		}

		start, err := strconv.Atoi(sSplit[i])
		if err != nil {
			log.Fatal(err)
		}

		for y := start; y < start+iVal; y++ {
			seeds = append(seeds, y)
		}
		i += 2

	}

	return seeds
}

func (g *Garden) GetSeeds(part2 bool) []int {
	if part2 {
		return g.Seeds2
	}
	return g.Seeds
}

func (g *Garden) LowestLocationNumber(part2 bool) int {
	lowest := 0

	for i, s := range g.GetSeeds(part2) {
		// seed to soil
		soil := s
		for _, soilmap := range g.SeedsToSoil {
			if s >= soilmap.Min && s <= soilmap.Max {
				soil = s + soilmap.Diff
			}
		}

		// soil to fertilizer
		fertil := soil
		for _, fertilmap := range g.SoilToFertilizer {
			if soil >= fertilmap.Min && soil <= fertilmap.Max {
				fertil = soil + fertilmap.Diff
			}
		}

		// fertilizer to water
		water := fertil
		for _, watermap := range g.FertilizerToWater {
			if fertil >= watermap.Min && fertil <= watermap.Max {
				water = fertil + watermap.Diff
			}
		}

		// water to light
		light := water
		for _, lightmap := range g.WaterToLight {
			if water >= lightmap.Min && water <= lightmap.Max {
				light = water + lightmap.Diff
			}
		}

		// light to temperature
		temp := light
		for _, tempmap := range g.LightToTemperature {
			if light >= tempmap.Min && light <= tempmap.Max {
				temp = light + tempmap.Diff
			}
		}

		// temperature to humidity
		humid := temp
		for _, humidmap := range g.TemperatureToHumidity {
			if temp >= humidmap.Min && temp <= humidmap.Max {
				humid = temp + humidmap.Diff
			}
		}

		// humidity to location
		loc := humid
		for _, locmap := range g.HumidityToLocation {
			if humid >= locmap.Min && humid <= locmap.Max {
				loc = humid + locmap.Diff
			}
		}
		if loc < lowest || lowest == 0 {
			lowest = loc
		}
		if i%10000000 == 0 {
			log.Printf("processed: %d", i)
		}
	}

	return lowest
}
