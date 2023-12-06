package main

import (
	"fmt"
	"github.com/iamlucasvieira/aoc/utils"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type condition struct {
	destination, source, width int
}

func (c condition) isSatisfied(value int) (int, bool) {
	satisfied := value >= c.source && value < c.source+c.width
	convertedValue := value - c.source + c.destination
	return convertedValue, satisfied
}

type conversion []condition

func (c *conversion) add(destination, source, width int) {
	*c = append(*c, condition{destination, source, width})
}

// adaptRanges adapt the range seeds to fit the conversion condition limits.
func (c *conversion) adaptRanges(rangeSeeds [][]int) [][]int {
	adaptedRangeSeeds := make([][]int, 0)
	remainingRangeSeeds := make([][]int, 0)

	for _, rangeSeed := range rangeSeeds {
		start := rangeSeed[0]
		end := rangeSeed[1]
		isWithinRange := false

		for _, condition := range *c {

			// Condition 1: start is inside the range
			if start >= condition.source && start < condition.source+condition.width {

				isWithinRange = true
				// Condition 1.1: end is inside the range
				if end >= condition.source && end < condition.source+condition.width {
					adaptedRangeSeeds = append(adaptedRangeSeeds, []int{start, end})
					break
				} else {
					// Condition 1.2: end is outside the range
					adaptedRangeSeeds = append(adaptedRangeSeeds, []int{start, condition.source + condition.width - 1})
					remainingRangeSeeds = append(remainingRangeSeeds, []int{condition.source + condition.width, end})
				}

			}
		}
		// Condition 2: start is outside the range
		if !isWithinRange {
			// Find the closest condition that initializing after the start
			closestEnd := end
			for _, condition := range *c {
				if condition.source > start && condition.source < closestEnd {
					closestEnd = condition.source
				}
			}
			// Condition 2.1: end is inside the range
			adaptedRangeSeeds = append(adaptedRangeSeeds, []int{start, closestEnd})

			// Condition 2.2: end is outside the range
			if closestEnd != end {
				remainingRangeSeeds = append(remainingRangeSeeds, []int{closestEnd + 1, end})
			}
		}
	}

	if len(remainingRangeSeeds) > 0 {
		adaptedRemainingSeeds := c.adaptRanges(remainingRangeSeeds)
		adaptedRangeSeeds = append(adaptedRangeSeeds, adaptedRemainingSeeds...)
	}

	return adaptedRangeSeeds
}

type instruction struct {
	seedToSoil            conversion
	soilToFertilizer      conversion
	fertilizerToWater     conversion
	waterToLight          conversion
	lightToTemperature    conversion
	temperatureToHumidity conversion
	humidityToLocation    conversion
}

func (i instruction) convert(v int, c conversion) int {
	for _, condition := range c {
		if convertedValue, satisfied := condition.isSatisfied(v); satisfied {
			return convertedValue
		}
	}
	return v
}

func (i instruction) adaptAndConvertRanges(v [][]int, c conversion) [][]int {
	// Adapt seed ranges for first conversion
	adaptedRanges := c.adaptRanges(v)

	// Get ranges for soil
	convertedRanges := make([][]int, 0)

	for _, adaptedRange := range adaptedRanges {
		convertedRanges = append(convertedRanges, i.convertRange(adaptedRange, c))
	}

	return convertedRanges
}

// convertRange converts a list of values using the conversion.
func (i instruction) convertRange(v []int, c conversion) []int {
	converted := make([]int, len(v))
	for idx, value := range v {
		converted[idx] = i.convert(value, c)
	}
	return converted
}

func (i instruction) convertSeedToLocation(seed int) int {
	soil := i.convert(seed, i.seedToSoil)
	fertilizer := i.convert(soil, i.soilToFertilizer)
	water := i.convert(fertilizer, i.fertilizerToWater)
	light := i.convert(water, i.waterToLight)
	temperature := i.convert(light, i.lightToTemperature)
	humidity := i.convert(temperature, i.temperatureToHumidity)
	location := i.convert(humidity, i.humidityToLocation)
	return location
}

func (i instruction) seedRangesToLocationRanges(seedRanges [][]int) [][]int {
	soil := i.adaptAndConvertRanges(seedRanges, i.seedToSoil)
	fertilizer := i.adaptAndConvertRanges(soil, i.soilToFertilizer)
	water := i.adaptAndConvertRanges(fertilizer, i.fertilizerToWater)
	light := i.adaptAndConvertRanges(water, i.waterToLight)
	temperature := i.adaptAndConvertRanges(light, i.lightToTemperature)
	humidity := i.adaptAndConvertRanges(temperature, i.temperatureToHumidity)
	location := i.adaptAndConvertRanges(humidity, i.humidityToLocation)
	return location
}

// parseInstructions parses the instruction data. Returns a list of seeds, instruction, and error.
func parseInstructions(data []string) ([]int, instruction, error) {

	var seeds []int

	conversionData := make([]conversion, 7)
	var idx = -1

	// Used to get all numbers from a string
	numberRegex := regexp.MustCompile(`\d+`)

	for _, line := range data {
		if strings.HasSuffix(line, "map:") {
			idx++
			conversionData[idx] = make(conversion, 0)
		} else if line == "" {
			continue
		} else if strings.HasPrefix(line, "seeds:") {
			// Get all seeds using regex
			for _, seed := range numberRegex.FindAllString(line, -1) {
				// Covert seed to int
				intSeed, err := strconv.Atoi(seed)
				if err != nil {
					return nil, instruction{}, err
				}
				seeds = append(seeds, intSeed)
			}

		} else {
			// Get all numbers using regex
			numbers := numberRegex.FindAllString(line, -1)
			if len(numbers) != 3 {
				return nil, instruction{}, fmt.Errorf("expected 3 numbers, got %v", len(numbers))
			}
			// Convert numbers to int
			intNumbers := make([]int, 3)
			for i, n := range numbers {
				intN, err := strconv.Atoi(n)
				if err != nil {
					return nil, instruction{}, err
				}
				intNumbers[i] = intN
			}

			// Add conversion
			conversionData[idx].add(intNumbers[0], intNumbers[1], intNumbers[2])
		}
	}

	// Create instruction
	instruction := instruction{
		seedToSoil:            conversionData[0],
		soilToFertilizer:      conversionData[1],
		fertilizerToWater:     conversionData[2],
		waterToLight:          conversionData[3],
		lightToTemperature:    conversionData[4],
		temperatureToHumidity: conversionData[5],
		humidityToLocation:    conversionData[6],
	}

	return seeds, instruction, nil
}

// parseInstructionsRangeSeeds parses the instruction data. Returns a list of seeds, instruction, and error.
// Seeds are pairs. First number is the start value and second number is the end value.
func parseInstructionsRangeSeeds(data []string) ([][]int, instruction, error) {
	seeds, instruction, err := parseInstructions(data)

	if err != nil {
		return nil, instruction, err
	}

	// Seeds are pairs. First number is the start value and second number is the end value.

	if len(seeds)%2 != 0 {
		return nil, instruction, fmt.Errorf("expected even number of seeds, got %v", len(seeds))
	}

	seedRanges := make([][]int, 0)

	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		end := start + seeds[i+1] - 1
		seedRanges = append(seedRanges, []int{start, end})
	}
	return seedRanges, instruction, err
}

func closestLocation(seeds []int, instruction instruction) int {
	// Get the first seed's location
	location := instruction.convertSeedToLocation(seeds[0])

	if len(seeds) == 1 {
		return location
	}

	// Find the closest location
	for _, seed := range seeds[1:] {
		seedLocation := instruction.convertSeedToLocation(seed)
		if seedLocation < location {
			location = seedLocation
		}
	}

	return location
}

func closestLocationRangeSeeds(seeds [][]int, instruction instruction) int {
	// Get the first seed's location
	locations := instruction.seedRangesToLocationRanges(seeds)
	closestLocation := locations[0][0]

	// find the closest location
	for _, location := range locations[1:] {
		if currentMin := slices.Min(location); currentMin < closestLocation {
			closestLocation = currentMin
		}
	}
	return closestLocation
}

func part1() {
	fmt.Println("Part 1:")
	data := utils.ReadFile("input.txt")
	seeds, instruction, err := parseInstructions(data)
	if err != nil {
		panic(err)
	}
	c := closestLocation(seeds, instruction)
	fmt.Printf("Closest location: %v\n", c)
}

func part2() {
	fmt.Println("Part 2:")
	data := utils.ReadFile("input2.txt")
	seeds, instruction, err := parseInstructionsRangeSeeds(data)
	if err != nil {
		panic(err)
	}
	c := closestLocationRangeSeeds(seeds, instruction)
	fmt.Printf("Closest location: %v\n", c)
}

func main() {
	part1()
	part2()
}
