package main

import (
	"github.com/iamlucasvieira/aoc/utils"
	"slices"
	"sort"
	"strconv"
	"testing"
)

var mockData = []string{
	"seeds: 79 14 55 13",
	"",
	"seed-to-soil map:",
	"50 98 2",
	"52 50 48",
	"",
	"soil-to-fertilizer map:",
	"0 15 37",
	"37 52 2",
	"39 0 15",
	"",
	"fertilizer-to-water map:",
	"49 53 8",
	"0 11 42",
	"42 0 7",
	"57 7 4",
	"",
	"water-to-light map:",
	"88 18 7",
	"18 25 70",
	"",
	"light-to-temperature map:",
	"45 77 23",
	"81 45 19",
	"68 64 13",
	"",
	"temperature-to-humidity map:",
	"0 69 1",
	"1 0 69",
	"",
	"humidity-to-location map:",
	"60 56 37",
	"56 93 4",
}

func TestParseInstructions(t *testing.T) {
	seeds, instruction, err := parseInstructions(mockData)

	expectedSeeds := []int{79, 14, 55, 13}

	// Sort the seeds
	sort.Ints(seeds)
	sort.Ints(expectedSeeds)

	if err != nil {
		t.Errorf("Error parsing instructions: %v", err)
	}

	if !slices.Equal(seeds, expectedSeeds) {
		t.Errorf("Expected seeds %v, got %v", expectedSeeds, seeds)
	}

	if instructionLength := len(instruction.seedToSoil); instructionLength != 2 {
		t.Errorf("Expected seedToSoil length 2, got %v", instructionLength)
	}

	if instructionLength := len(instruction.soilToFertilizer); instructionLength != 3 {
		t.Errorf("Expected soilToFertilizer length 3, got %v", instructionLength)
	}

}

func TestConvert(t *testing.T) {
	_, instruction, _ := parseInstructions(mockData)

	if value1 := instruction.convert(99, instruction.seedToSoil); value1 != 51 {
		t.Errorf("Expected convert(99, seedToSoil) to be 51, got %v", value1)
	}

	if value2 := instruction.convert(1000, instruction.soilToFertilizer); value2 != 1000 {
		t.Errorf("Expected convert(1000, soilToFertilizer) to be 1000, got %v", value2)
	}
}

func TestConvertSeedToLocation(t *testing.T) {
	testCases := []utils.TestCase[int, int]{
		{Input: 79, Expected: 82},
		{Input: 14, Expected: 43},
		{Input: 55, Expected: 86},
		{Input: 13, Expected: 35},
	}

	_, instruction, _ := parseInstructions(mockData)

	for _, tc := range testCases {
		// convert int to string
		inputStr := strconv.Itoa(tc.Input)
		t.Run(inputStr, func(t *testing.T) {
			location := instruction.convertSeedToLocation(tc.Input)
			if location != tc.Expected {
				t.Errorf("Expected location %v, got %v", tc.Expected, location)
			}
		})
	}
}

func TestClosestLocation(t *testing.T) {
	seeds, instruction, _ := parseInstructions(mockData)
	expectedClosestLocation := 35

	closestLocation := closestLocation(seeds, instruction)

	if closestLocation != expectedClosestLocation {
		t.Errorf("Expected closest location %v, got %v", expectedClosestLocation, closestLocation)
	}

}

func TestParseInstructionsRangeSeeds(t *testing.T) {
	seeds, _, _ := parseInstructionsRangeSeeds(mockData)

	if len(seeds) != 2 {
		t.Errorf("Expected seeds length 2, got %v", len(seeds))
	}

	if seeds[0][0] != 79 || seeds[0][1] != 92 {
		t.Errorf("Expected seeds[0] to be [79, 92], got %v", seeds[0])
	}

	if seeds[1][0] != 55 || seeds[1][1] != 67 {
		t.Errorf("Expected seeds[1] to be [55, 67], got %v", seeds[1])
	}
}

func TestClosestLocationRangeSeeds(t *testing.T) {
	seeds, instruction, _ := parseInstructionsRangeSeeds(mockData)
	expectedClosestLocation := 46

	closestLocation := closestLocationRangeSeeds(seeds, instruction)

	if closestLocation != expectedClosestLocation {
		t.Errorf("Expected closest location %v, got %v", expectedClosestLocation, closestLocation)
	}
}

func TestConversionAdaptRanges(t *testing.T) {
	_, instruction, _ := parseInstructionsRangeSeeds(mockData)

	seedToSoil := instruction.seedToSoil

	if adapted := seedToSoil.adaptRanges([][]int{{50, 97}}); len(adapted) != 1 {
		t.Errorf("Expected adapted length 1, got %v", len(adapted))
	}
	if adapted := seedToSoil.adaptRanges([][]int{{79, 99}}); len(adapted) != 2 {
		t.Errorf("Expected adapted length 2, got %v", len(adapted))
	}

	if adapted := seedToSoil.adaptRanges([][]int{{79, 101}}); len(adapted) != 3 {
		t.Errorf("Expected adapted length 3, got %v", len(adapted))
	}

	if adapted := seedToSoil.adaptRanges([][]int{{10, 102}}); len(adapted) != 4 {
		t.Errorf("Expected adapted length 4, got %v", len(adapted))
	}
}

func TestConvertRange(t *testing.T) {
	_, instruction, _ := parseInstructionsRangeSeeds(mockData)

	if value1 := instruction.convertRange([]int{0, 10}, instruction.seedToSoil); value1[0] != 0 || value1[1] != 10 {
		t.Errorf("Expected convertRange([0, 10], seedToSoil) to be [0, 10], got %v", value1)
	}

	if value2 := instruction.convertRange([]int{98, 99}, instruction.seedToSoil); value2[0] != 50 || value2[1] != 51 {
		t.Errorf("Expected convertRange([98, 99], seedToSoil) to be [50, 51], got %v", value2)
	}
}

func TestAdaptAndConvertRange(t *testing.T) {
	_, instruction, _ := parseInstructionsRangeSeeds(mockData)

	ranges := [][]int{{79, 99}}

	convertedRanges := instruction.adaptAndConvertRanges(ranges, instruction.seedToSoil)

	if len(convertedRanges) != 2 {
		t.Errorf("Expected convertedRanges length 2, got %v", len(convertedRanges))
	}

	if convertedRanges[0][0] != 81 || convertedRanges[0][1] != 99 {
		t.Errorf("Expected convertedRanges[0] to be [81, 99], got %v", convertedRanges[0])
	}

	if convertedRanges[1][0] != 50 || convertedRanges[1][1] != 51 {
		t.Errorf("Expected convertedRanges[1] to be [50, 51], got %v", convertedRanges[1])
	}

}
