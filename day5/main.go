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

	log.Printf("Part 1: %d", g.LowestLocationNumber())
}

type Garden struct {
	Seeds                 []int
	SeedsToSoil           []map[int]int
	SoilToFertilizer      []map[int]int
	FertilizerToWater     []map[int]int
	WaterToLight          []map[int]int
	LightToTemperature    []map[int]int
	TemperatureToHumidity []map[int]int
	HumidityToLocation    []map[int]int
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

func generateMap(line string) map[int]int {
	smap := map[int]int{}

	sline := strings.Split(line, " ")

	iteration, _ := strconv.Atoi(sline[2])
	destination, _ := strconv.Atoi(sline[0])
	source, _ := strconv.Atoi(sline[1])

	for i := 0; i < iteration; i++ {
		smap[source+i] = destination + i
	}

	return smap
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

func (g *Garden) LowestLocationNumber() int {
	locations := []int{}
	for _, s := range g.Seeds {
		// seed to soil
		soil := s
		for _, soilmap := range g.SeedsToSoil {
			if _, ok := soilmap[s]; ok {
				soil = soilmap[s]
			}
		}

		// soil to fertilizer
		fertil := soil
		for _, fertilmap := range g.SoilToFertilizer {
			if _, ok := fertilmap[soil]; ok {
				fertil = fertilmap[soil]
			}
		}

		// fertilizer to water
		water := fertil
		for _, watermap := range g.FertilizerToWater {
			if _, ok := watermap[fertil]; ok {
				water = watermap[fertil]
			}
		}

		// water to light
		light := water
		for _, lightmap := range g.WaterToLight {
			if _, ok := lightmap[water]; ok {
				light = lightmap[water]
			}
		}

		// light to temperature
		temp := light
		for _, tempmap := range g.LightToTemperature {
			if _, ok := tempmap[light]; ok {
				temp = tempmap[light]
			}
		}

		// temperature to humidity
		humid := temp
		for _, humidmap := range g.TemperatureToHumidity {
			if _, ok := humidmap[temp]; ok {
				humid = humidmap[temp]
			}
		}

		// humidity to location
		loc := humid
		for _, locmap := range g.HumidityToLocation {
			if _, ok := locmap[humid]; ok {
				loc = locmap[humid]
			}
		}
		locations = append(locations, loc)
	}

	lowest := 0

	for _, l := range locations {
		if l < lowest || lowest == 0 {
			lowest = l
		}
	}

	return lowest
}
