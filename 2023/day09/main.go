package main

import (
	"fmt"
	"github.com/iamlucasvieira/aoc/utils"
	"strconv"
	"strings"
)

// parseData parses the data into slices. Each line contains numbers separated by spaces.
func parseData(lines []string) ([][]int, error) {

	var numbersLists [][]int
	for _, line := range lines {
		if line == "" {
			continue
		}

		var numbers []int
		for _, number := range strings.Split(line, " ") {
			// Turn the string into an int
			n, err := strconv.Atoi(number)
			if err != nil {
				return nil, err
			}
			numbers = append(numbers, n)
		}

		numbersLists = append(numbersLists, numbers)

	}
	return numbersLists, nil
}

// diff returns the difference between elements in a slice.
func diff(n []int) []int {
	var diff []int
	for i := 1; i < len(n); i++ {
		diff = append(diff, n[i]-n[i-1])
	}
	return diff
}

// allZeros returns true if all elements in a slice are 0.
func allZeros(n []int) bool {
	for _, v := range n {
		if v != 0 {
			return false
		}
	}
	return true
}

// diffIterations returns the first or end element of the diff of a slice iteratively until all elements are 0.
func diffIterations(n []int, end bool) []int {

	var allFinals []int
	for !allZeros(n) {
		n = diff(n)
		var idx = 0
		if end {
			idx = len(n) - 1
		}
		allFinals = append(allFinals, n[idx])
	}
	return allFinals
}

// nextNumber returns the next number in the sequence.
func nextNumber(n []int) int {
	if len(n) == 0 {
		return 0
	}

	diffHistory := diffIterations(n, true)

	// Sum the diffHistory
	sum := 0
	for _, v := range diffHistory {
		sum += v
	}

	// Return the sum + the last number in the input
	return sum + n[len(n)-1]
}

// previousNumber returns the previous number in the sequence.
func previousNumber(n []int) int {
	if len(n) == 0 {
		return 0
	}

	diffHistory := diffIterations(n, false)

	// Sum the diffHistory
	sum := 0
	// Loop through the diffHistory in reverse
	for i := len(diffHistory) - 1; i >= 0; i-- {
		sum = diffHistory[i] - sum
	}

	return n[0] - sum
}

// sumListPredictions sums the predictions of a slice of slices.
func sumListPredictions(nLists [][]int, predictionFunc func([]int) int) int {
	sum := 0
	for _, n := range nLists {
		sum += predictionFunc(n)
	}
	return sum
}

// sumNextNumbers sums the next numbers in the sequence.
func sumNextNumbers(nLists [][]int) int {
	return sumListPredictions(nLists, nextNumber)
}

// sumPreviousNumbers sums the previous numbers in the sequence.
func sumPreviousNumbers(nLists [][]int) int {
	return sumListPredictions(nLists, previousNumber)
}

func part1() {
	fmt.Println("Part 1:")
	data, err := parseData(utils.ReadFile("input.txt"))
	if err != nil {
		panic(err)
	}
	sum := sumNextNumbers(data)
	fmt.Printf("Sum: %d\n", sum)
}

func part2() {
	fmt.Println("Part 2:")
	data, err := parseData(utils.ReadFile("input2.txt"))
	if err != nil {
		panic(err)
	}
	sum := sumPreviousNumbers(data)
	fmt.Printf("Sum: %d\n", sum)
}

func main() {
	part1()
	part2()
}
