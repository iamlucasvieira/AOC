package main

import (
	"fmt"
	"github.com/iamlucasvieira/aoc/utils"
	"strings"
)

// sharedItem returns the first character that appears in both halves of the input string.
func sharedItem(s string) string {
	var mid = len(s) / 2
	for i := 0; i < mid; i++ {
		var current = string(s[i])
		if strings.Contains(s[mid:], current) {
			return current
		}
	}
	return ""
}

func sharedItemStrings(s []string) string {
	if len(s) == 0 {
		return ""
	}

	if len(s) == 1 {
		return s[0]
	}

	// Create a copy of the slice
	sCopy := make([]string, len(s))
	copy(sCopy, s)

	var first, second = sCopy[0], sCopy[1]
	rest := sCopy[2:]

	var sharedItems string

	for _, item := range first {
		itemString := string(item)

		if strings.Contains(second, itemString) && !strings.Contains(sharedItems, itemString) {
			sharedItems += itemString
		}
	}

	rest = append(rest, sharedItems)
	return sharedItemStrings(rest)
}

// priority turns a string into a rune and returns its priority.
// a->z: 1->26
// A->Z: 27->52
func priority(s string) int {
	if len(s) != 1 {
		return -1
	}

	var r = rune(s[0])

	if r >= 'a' && r <= 'z' {
		return int(r) - int('a') + 1
	}

	if r >= 'A' && r <= 'Z' {
		return int(r) - int('A') + 27
	}

	return -1
}

// priorityOfSharedItems returns the sum of the priorities of the shared items.
func priorityOfSharedItems(s []string) int {
	var sum int

	for _, backpack := range s {
		currentPriority := priority(sharedItem(backpack))
		if currentPriority == -1 { // no shared item
			continue
		}
		sum += currentPriority
	}

	return sum
}

func priorityOfSharedItemsThree(s []string) int {
	var sum int
	// Check every three rows
	for i := 0; i < len(s); i += 3 {
		if len(s[i]) == 0 { // empty row
			continue
		}

		shared := sharedItemStrings(s[i : i+3])
		if len(shared) != 1 {
			continue
		}
		sum += priority(shared)
	}
	return sum
}

func part1() {
	fmt.Println("Part 1:")
	data := utils.ReadFile("input.txt")
	sum := priorityOfSharedItems(data)
	fmt.Printf("Sum of priorities: %d\n", sum)
}

func part2() {
	fmt.Println("Part 2:")
	data := utils.ReadFile("input2.txt")
	fmt.Printf("Shared items: %v\n", priorityOfSharedItemsThree(data))
}
func main() {
	part1()
	part2()
}
