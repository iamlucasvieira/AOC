package main

import (
	"fmt"
	"strings"

	"github.com/iamlucasvieira/aoc/utils"
)

func parse(lines []string) []string {
	var result []string
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		result = append(result, line)
	}
	return result
}

func firstUniqueSequence(line string, n int) int {
	// Returns the first time a string contains n unique characters
	var tempLine string
	for i := 0; i < len(line); i++ {
		index := strings.Index(tempLine, string(line[i]))

		// If the character is not in the string, add it
		if index == -1 {
			tempLine += string(line[i])
		} else {
			tempLine = tempLine[index+1:] + string(line[i])
		}

		// If the string has n unique characters, return the index
		if len(tempLine) == n {
			return i + 1
		}
	}
	return -1
}

func part1() int {
	fmt.Println("Part 1:")
	lines := parse(utils.ReadFile("input.txt"))
	value := firstUniqueSequence(lines[0], 4)
	fmt.Println(value)
	return value
}

func part2() int {
	fmt.Println("Part 2:")
	lines := parse(utils.ReadFile("input.txt"))
	value := firstUniqueSequence(lines[0], 14)
	fmt.Println(value)
	return value
}

func main() {
	part1()
	part2()
}
