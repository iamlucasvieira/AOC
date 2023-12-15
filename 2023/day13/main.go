package main

import (
	"errors"
	"fmt"
	"github.com/iamlucasvieira/aoc/utils"
	"strings"
)

// parse takes a slice of strings and returns a slice of patterns.
func parse(input []string) ([][][]string, error) {
	var patterns [][][]string
	var pattern [][]string

	for _, line := range input {
		if line == "" {
			patterns = append(patterns, pattern)
			pattern = [][]string{}
			continue
		}

		pattern = append(pattern, strings.Split(line, ""))
	}

	if len(pattern) > 0 {
		patterns = append(patterns, pattern)
	}

	return patterns, nil
}

// isReflective takes a slice and return true if the slice is a reflection on axis x.
// x is the number of columns on the left side of the slice.
func isReflective(s []string, x int) bool {

	sliceLength := len(s)
	// If x is 0 or the length of the slice, it's not a reflection.
	if x == 0 || x == sliceLength {
		return false
	}

	left, right := s[:x], s[x:]

	// get the minimum length of the two slices. Only that length needs to be checked.
	minIdx := min(x, sliceLength-x)

	for i := 0; i < minIdx; i++ {
		// Check if the left and right side are equal. Left side is reversed.
		if left[len(left)-i-1] != right[i] {
			return false
		}
	}

	return true
}

func nonReflectivePoints(s []string, x int) (int, error) {
	var count int

	sliceLength := len(s)
	// If x is 0 or the length of the slice, it's not a reflection.
	if x == 0 || x == sliceLength {
		return 0, errors.New("x is 0 or the length of the slice")
	}

	left, right := s[:x], s[x:]

	// get the minimum length of the two slices. Only that length needs to be checked.
	minIdx := min(x, sliceLength-x)

	for i := 0; i < minIdx; i++ {
		// Check if the left and right side are equal. Left side is reversed.
		if left[len(left)-i-1] != right[i] {
			count++
		}
	}

	return count, nil
}

// getColumn returns column idx as a row.
func getColumn(s [][]string, idx int) []string {
	var row = make([]string, len(s))

	for i := range row {
		row[i] = s[i][idx]
	}
	return row
}

// getRow returns row idx as a column.
func getRow(s [][]string, idx int) []string {
	return s[idx]
}

// possibleReflections returns a slice of possible reflections for a pattern.
func possibleReflections(s []string) []int {
	var reflections []int

	for i := 1; i < len(s); i++ {
		if isReflective(s, i) {
			reflections = append(reflections, i)
		}
	}

	return reflections
}

func reflectionPoint(p [][]string, getSlice func([][]string, int) []string, lenColumns int) int {
	firstRow := getSlice(p, 0)
	possibleReflections := possibleReflections(firstRow)

	for _, r := range possibleReflections {
		isPossible := true
		for i := 1; i < lenColumns; i++ {
			if !isReflective(getSlice(p, i), r) {
				isPossible = false
				break
			}
		}
		if isPossible {
			return r
		} else {
			continue
		}
	}
	return -1
}

func reflectionPointFixing(p [][]string, getSlice func([][]string, int) []string, lenColumns int) int {

	firstRow := getSlice(p, 0)

	// Try every possible row reflection point.
	for i := 1; i < len(firstRow); i++ {
		brokenCount := 0
		// Check all columns.
		for j := 0; j < lenColumns; j++ {
			b, err := nonReflectivePoints(getSlice(p, j), i)
			if err != nil {
				panic(err)
			}
			brokenCount += b

			if brokenCount > 1 {
				break
			}
		}
		if brokenCount == 1 {
			return i
		}
	}
	return -1
}

func rowReflectionPointFixing(p [][]string) int {
	return reflectionPointFixing(p, getColumn, len(p[0]))
}

func columnReflectionPointFixing(p [][]string) int {
	return reflectionPointFixing(p, getRow, len(p))
}

// rowReflectionPoint return the number of rows before a reflection point.
func rowReflectionPoint(p [][]string) int {
	return reflectionPoint(p, getColumn, len(p[0]))
}

// columnReflectionPoint return the number of columns before a reflection point.
func columnReflectionPoint(p [][]string) int {
	return reflectionPoint(p, getRow, len(p))
}

// patternScoreTemplate returns the score of a pattern.
func patternScoreTemplate(p [][]string, rowFunc, colFunc func([][]string) int) (int, error) {
	colReflection := colFunc(p)
	if colReflection != -1 {
		return colReflection, nil
	}

	rowReflection := rowFunc(p)
	if rowReflection != -1 {
		return rowReflection * 100, nil
	}

	return 0, errors.New("no reflection point found")
}

// patternScore returns the score of a pattern.
func patternScore(p [][]string) (int, error) {
	return patternScoreTemplate(p, rowReflectionPoint, columnReflectionPoint)
}

// patternScoreFixing returns the score of a pattern allowing one broken reflection.
func patternScoreFixing(p [][]string) (int, error) {
	return patternScoreTemplate(p, rowReflectionPointFixing, columnReflectionPointFixing)
}

// sumScoreTemplate returns the sum of the scores of a slice of patterns.
func sumScoreTemplate(p [][][]string, scoreFunc func([][]string) (int, error)) int {
	var sum int

	for _, pattern := range p {
		score, err := scoreFunc(pattern)
		if err != nil {
			continue
		}
		sum += score
	}

	return sum
}

// sumScore returns the sum of the scores of a slice of patterns.
func sumScore(p [][][]string) int {
	return sumScoreTemplate(p, patternScore)
}

// sumScoreFixing returns the sum of the scores of a slice of patterns allowing one broken reflection.
func sumScoreFixing(p [][][]string) int {
	return sumScoreTemplate(p, patternScoreFixing)
}

func part1() {
	fmt.Println("Part 1:")
	data, err := parse(utils.ReadFile("input.txt"))

	if err != nil {
		panic(err)
	}

	score := sumScore(data)
	fmt.Printf("Score: %v\n", score)
}

func part2() {
	fmt.Println("Part 2:")
	data, err := parse(utils.ReadFile("input2.txt"))

	if err != nil {
		panic(err)
	}

	score := sumScoreFixing(data)
	fmt.Printf("Score: %v\n", score)
}

func main() {
	part1()
	part2()
}
