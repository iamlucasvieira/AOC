package main

import "testing"

var mockGrid = []string{
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
}

var mockGridGear = []string{
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
}

func TestParseGrid(t *testing.T) {
	grid, chars := parseGrid(mockGrid)

	if len(grid) != 28 {
		t.Errorf("Expected 8 numbers, got %d", len(grid))
	}

	if len(chars) != 6 {
		t.Errorf("Expected 6 chars, got %d", len(chars))
	}

	if grid[point{0, 0}] != "4" {
		t.Errorf("Expected 4, got %s", grid[point{0, 0}])
	}
}

func TestGetNeighbours(t *testing.T) {
	neighbours := getNeighbours(point{1, 1})

	if len(neighbours) != 8 {
		t.Errorf("Expected 8 neighbours, got %d", len(neighbours))
	}

}

func TestGetXNeighbours(t *testing.T) {
	neighbours := getXNeighbours(point{1, 1})

	if len(neighbours) != 2 {
		t.Errorf("Expected 2 neighbours, got %d", len(neighbours))
	}
}

func TestBuildNumber(t *testing.T) {
	grid, _ := parseGrid(mockGrid)

	number := buildNumber(point{0, 0}, grid)

	if len(number) != 3 {
		t.Errorf("Expected 3 numbers, got %d", len(number))
	}

	value, ok := number[point{0, 0}]
	if !ok || value != "4" {
		t.Errorf("Expected to find 4 in number grid")
	}

	value, ok = number[point{1, 0}]
	if !ok || value != "6" {
		t.Errorf("Expected to find 6 in number grid")
	}

	value, ok = number[point{2, 0}]
	if !ok || value != "7" {
		t.Errorf("Expected to find 7 in number grid")
	}

}

func TestNumberMapToInt(t *testing.T) {
	grid, _ := parseGrid(mockGrid)

	number := buildNumber(point{0, 0}, grid)

	numberInt := numberMapToInt(number)

	if numberInt != 467 {
		t.Errorf("Expected 467, got %d", numberInt)
	}
}

func TestSumValidNumbers(t *testing.T) {
	grid, characters := parseGrid(mockGrid)
	sum := sumValidNumbers(grid, characters)

	if sum != 4361 {
		t.Errorf("Expected %d, got %d", 4361, sum)
	}
}

func TestParseGridGear(t *testing.T) {
	grid, chars := parseGridGear(mockGridGear)

	if len(grid) != 28 {
		t.Errorf("Expected 28 number, got %d", len(grid))
	}

	if len(chars) != 3 {
		t.Errorf("Expected 3 chars, got %d", len(chars))
	}
}

func TestSumValidNumbersGear(t *testing.T) {
	grid, characters := parseGridGear(mockGridGear)
	sum := sumValidNumbersGear(grid, characters)

	if sum != 467835 {
		t.Errorf("Expected %d, got %d", 467835, sum)
	}
}
