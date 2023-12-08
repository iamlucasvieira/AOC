package main

import (
	"fmt"
	"github.com/iamlucasvieira/aoc/utils"
	"slices"
	"sort"
	"strconv"
)

// parseData parses the data into slices
func parseData(data []string) ([]int, error) {
	var calories = []int{0}

	idx := 0
	for _, line := range data {

		if line == "" {
			calories = append(calories, 0)
			idx++
			continue
		}

		calorie, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		calories[idx] += calorie

	}
	return calories, nil
}

// topThreeSum returns the sum of the top three slices
func topThreeSum(calories []int) int {
	if len(calories) == 0 {
		return 0 // No elements to sum
	}

	// Sorting the slice in ascending order
	sort.Ints(calories)

	// Summing up the top three elements
	sum := 0
	for i := 0; i < 3 && i < len(calories); i++ {
		sum += calories[len(calories)-1-i]
	}
	return sum
}

func part1() {
	fmt.Println("Part 1:")
	data, err := parseData(utils.ReadFile("input.txt"))
	if err != nil {
		panic(err)
	}
	maxCalories := slices.Max(data)
	fmt.Printf("Max calories: %d\n", maxCalories)
}

func part2() {
	fmt.Println("Part 2:")
	data, err := parseData(utils.ReadFile("input.txt"))
	if err != nil {
		panic(err)
	}
	maxThree := topThreeSum(data)
	fmt.Printf("Sum of highest three: %d\n", maxThree)
}
func main() {
	part1()
	part2()
}
