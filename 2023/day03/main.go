package main

import (
	"fmt"
	"github.com/iamlucasvieira/aoc/utils"
	"sort"
	"strconv"
)

type point struct {
	x, y int
}

type grid map[point]string
type charChecker func(rune) bool

// parseGridTemplate parses the input using a checker function
func parseGridTemplate(input []string, checker charChecker) (grid, []point) {
	grid := make(grid)
	var chars []point

	for i, line := range input {
		for j, char := range line {
			// Check if char is a number
			if char >= '0' && char <= '9' {
				grid[point{j, i}] = string(char)
			} else if checker(char) {
				chars = append(chars, point{j, i})
			}
		}
	}

	return grid, chars
}

// parseGrid parses the input
// Returns a map with all numbers and their locations
// Returns a slice with all the locations of characters different from '.'
func parseGrid(input []string) (grid, []point) {
	checker := func(char rune) bool {
		return char != '.'
	}
	return parseGridTemplate(input, checker)
}

// parseGridGear parses the input. Returns a grid with all numbers and a slice with all the locations of '*'
func parseGridGear(input []string) (grid, []point) {
	checker := func(char rune) bool {
		return char == '*'
	}
	return parseGridTemplate(input, checker)
}

// getNeighbours of a point in the grid including diagonals
func getNeighbours(p point) []point {
	return []point{
		{p.x - 1, p.y - 1},
		{p.x, p.y - 1},
		{p.x + 1, p.y - 1},
		{p.x - 1, p.y},
		{p.x + 1, p.y},
		{p.x - 1, p.y + 1},
		{p.x, p.y + 1},
		{p.x + 1, p.y + 1},
	}
}

// getXNeighbours returns the neighbours of a point in the grid in the x axis
func getXNeighbours(p point) []point {
	return []point{
		{p.x - 1, p.y},
		{p.x + 1, p.y},
	}
}

// buildNumber connects numbers in a grid with its x-axis neighbours
func buildNumber(startPoint point, g grid) grid {
	var numberGrid grid = make(grid)

	numberGrid[startPoint] = g[startPoint]

	// Helper function to append points in a direction until a condition is met
	appendDirection := func(step int) {
		for i := 1; ; i++ {
			nextPoint := point{startPoint.x + i*step, startPoint.y}
			if value, ok := g[nextPoint]; ok {
				numberGrid[nextPoint] = value
			} else {
				break
			}
		}
	}

	appendDirection(-1) // Move left
	appendDirection(1)  // Move right

	return numberGrid
}

// numberMapToInt converts a map of points to an int value according to their values
func numberMapToInt(n grid) int {
	var result string

	// Get y value of the first point
	var y int
	for k := range n {
		y = k.y
		break
	}

	// Sort by x value in the map
	var keys []int
	for k := range n {
		keys = append(keys, k.x)
	}

	sort.Ints(keys)

	for _, k := range keys {
		result += n[point{k, y}]
	}

	// Convert to int
	intValue, err := strconv.Atoi(result)
	if err != nil {
		panic(err)
	}

	return intValue
}

// sumValidNumbers sums all the numbers in the grid that are neighbours to a character
func sumValidNumbers(g grid, chars []point) int {
	var sum int

	visitedNumbers := make(grid)

	for _, char := range chars {
		// Get neighbours of the char
		neighbours := getNeighbours(char)

		// Check if neighbours are numbers and not visited
		for _, neighbour := range neighbours {
			_, isNumber := g[neighbour]
			_, isVisited := visitedNumbers[neighbour]

			if isNumber && !isVisited {
				// Build number
				number := buildNumber(neighbour, g)

				// Convert to int
				numberInt := numberMapToInt(number)

				// Add to sum
				sum += numberInt

				// Mark as visited
				for k := range number {
					visitedNumbers[k] = number[k]
				}
			}

		}
	}

	return sum
}

// sumValidNumbersGear sums all the product of numbers in the grid that are neighbours to a "*
func sumValidNumbersGear(g grid, chars []point) int {
	var sum int

	visitedNumbers := make(grid)

	for _, char := range chars {
		// Get neighbours of the char
		neighbours := getNeighbours(char)
		product := 1
		nNumberNeighbours := 0
		// Check if neighbours are numbers and not visited
		for _, neighbour := range neighbours {
			_, isNumber := g[neighbour]
			_, isVisited := visitedNumbers[neighbour]

			if isNumber && !isVisited {
				// Build number
				number := buildNumber(neighbour, g)

				// Convert to int
				numberInt := numberMapToInt(number)

				// Multiply to product
				product *= numberInt
				nNumberNeighbours++

				// Mark as visited
				for k := range number {
					visitedNumbers[k] = number[k]
				}
			}
		}
		// Add product to sum if there are more than 1 neighbours to a "*"
		if nNumberNeighbours > 1 {
			sum += product
		}
	}
	return sum
}

func part1() {
	fmt.Println("Part 1:")
	data := utils.ReadFile("input.txt")
	g, c := parseGrid(data)
	sum := sumValidNumbers(g, c)
	fmt.Printf("Sum: %d\n", sum)
}

func part2() {
	fmt.Println("Part 2:")
	data := utils.ReadFile("input2.txt")
	g, c := parseGridGear(data)
	sum := sumValidNumbersGear(g, c)
	fmt.Printf("Sum: %d\n", sum)
}

func main() {
	part1()
	part2()
}
