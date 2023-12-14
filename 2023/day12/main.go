package main

import (
	"fmt"
	"github.com/iamlucasvieira/aoc/utils"
	"slices"
	"strconv"
	"strings"
)

// record is a struct that contains the string and integer representation of the functioning of hot springs
type record struct {
	strings  []string
	integers []int
}

// String returns a string representation of a record strings and int slices
func (r record) String() string {
	return fmt.Sprintf("%v %v", r.strings, r.integers)
}

func sliceSum(slice []int) int {
	var sum int
	for _, num := range slice {
		sum += num
	}
	return sum
}

func countMatches(cache map[string]int, r record, size int) int {

	// Check if the record is already in the cache
	if val, ok := cache[r.String()]; ok {
		return val
	}

	// Check if the record integers is empty (Therefore, when no more groups to match)
	if len(r.integers) == 0 {
		// If empty, the strings should not contain any damaged #
		if slices.Contains(r.strings, "#") {
			return 0
		}
		return 1
	}

	// Get fist and future groups of integers
	firstGroup := r.integers[0]
	futureGroups := r.integers[1:]

	// Get length of future groups = length of each group + "." that separates them
	futureGroupsLength := len(futureGroups) + sliceSum(futureGroups)

	// Get maximum start index of the first group
	maximumStart := size - futureGroupsLength - firstGroup

	var matches int = 0

	// Loop through all possible start indexes of the first group
	for i := 0; i <= maximumStart; i++ {
		// Build candidate string
		c := buildCandidateString(i, firstGroup)

		// Check if candidate matches the first group
		if possiblePattern(c, r.strings) {
			candidateLength := len(c)
			nextStart := candidateLength

			// If this is the final group, neglect the "." at the end
			if len(futureGroups) == 0 {
				nextStart--
			}

			newStrings := make([]string, len(r.strings)-nextStart)
			copy(newStrings, r.strings[nextStart:])
			nextPattern := record{newStrings, futureGroups}
			matches += countMatches(cache, nextPattern, size-nextStart)
		}

	}

	// Add to cache
	cache[r.String()] = matches

	return matches
}

func buildCandidateString(nDot, nHash int) []string {
	c := strings.Repeat(".", nDot) + strings.Repeat("#", nHash) + "."
	return strings.Split(c, "")
}

// possiblePatterns returns true if a candidate string matches the pattern string
func possiblePattern(candidate []string, pattern []string) bool {
	for i, char := range candidate {

		if i == len(pattern) {
			return true
		}
		if pattern[i] == "?" {
			continue
		}
		if pattern[i] != char {
			return false
		}
	}

	return true
}

func parseData(data []string) ([]record, error) {
	var records []record

	for _, line := range data {
		// Split line into two parts by " "
		if len(line) == 0 {
			continue
		}
		lineParts := strings.Split(line, " ")

		if len(lineParts) != 2 {
			return nil, fmt.Errorf("lineParts should have 2 parts, got %d", len(lineParts))
		}

		// Turn first part into list of strings
		var recordStrings []string
		recordStrings = strings.Split(lineParts[0], "")

		// Turn second part into list of integers
		var recordIntegers []int
		for _, num := range strings.Split(lineParts[1], ",") {

			// Convert string to int
			numInt, err := strconv.Atoi(num)
			if err != nil {
				return nil, fmt.Errorf("strconv.Atoi() returned error: %v", err)
			}
			recordIntegers = append(recordIntegers, numInt)
		}

		// Append record to records
		records = append(records, record{recordStrings, recordIntegers})
	}
	return records, nil
}

func sumOfMatches(records []record) int {
	var sum int
	cache := make(map[string]int)
	for _, r := range records {
		sum += countMatches(cache, r, len(r.strings))
	}
	return sum
}

// unfold turns the list of strings into n copies of itself separated by "?" and integers into n copies of itself
func unfold(r record, n int) record {
	var newStrings []string
	var newIntegers []int

	for i := 0; i < n; i++ {
		newStrings = append(newStrings, r.strings...)
		newStrings = append(newStrings, "?")
		newIntegers = append(newIntegers, r.integers...)
	}

	// Return excluding the last "?"
	return record{newStrings[:len(newStrings)-1], newIntegers}
}

// unfoldAll turns the list of strings into n copies of itself separated by "?" and integers into n copies of itself
func unfoldAll(records []record, n int) []record {
	var newRecords []record
	for _, r := range records {
		newRecords = append(newRecords, unfold(r, n))
	}
	return newRecords
}

func part1() {
	fmt.Println("Part 1:")
	records, err := parseData(utils.ReadFile("input.txt"))
	if err != nil {
		panic(err)
	}
	sum := sumOfMatches(records)
	fmt.Printf("Sum of matches: %d\n", sum)
}

func part2() {
	fmt.Println("Part 2:")
	records, err := parseData(utils.ReadFile("input2.txt"))
	if err != nil {
		panic(err)
	}
	records = unfoldAll(records, 5)
	sum := sumOfMatches(records)
	fmt.Printf("Sum of matches: %d\n", sum)
}
func main() {
	part1()
	part2()
}
