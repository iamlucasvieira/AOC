package main

import (
	"fmt"
	"github.com/iamlucasvieira/aoc/utils"
	"strconv"
	"strings"
)

// parse takes a slice of strings and returns a slice of strings, where each string represents a value
func parse(lines []string) []string {
	// Turn slice of strings into a single string
	var singleString string
	for _, line := range lines {
		singleString += line
	}

	// Split string into slice of strings
	return strings.Split(singleString, ",")
}

func parseLenses(lines []string) []lens {
	var lenses []lens
	for _, l := range parse(lines) {
		if strings.Contains(l, "=") {
			split := strings.Split(l, "=")
			label := split[0]
			focal, err := strconv.Atoi(split[1])
			if err != nil {
				panic(err)
			}
			lenses = append(lenses, lens{label, focal, "="})
		} else {
			split := strings.Split(l, "-")
			label := split[0]
			lenses = append(lenses, lens{label, -1, "-"})
		}
	}
	return lenses
}

// hash takes a string and returns an int
func hash(line string) int {
	var value int

	for _, char := range line {
		value += int(char)
		value *= 17
		value = value % 256
	}

	return value
}

// hashSum returns the sum of the hash of each string in a slice of strings
func hashSum(lines []string) int {
	var sum int
	for _, line := range lines {
		sum += hash(line)
	}
	return sum
}

type lens struct {
	label     string
	focal     int
	operation string
}

type boxes map[int][]lens

func (b boxes) index(boxIndex int, l lens) int {
	box := b[boxIndex]
	for i, lens := range box {
		if lens.label == l.label {
			return i
		}
	}
	return -1
}

// addLens adds a lens to a box
func addLens(b boxes, l lens) {
	toBox := hash(l.label)
	i := b.index(toBox, l)

	if i == -1 {
		b[toBox] = append(b[toBox], l)
	} else {
		b[toBox][i] = l
	}
}

// removeLens removes a lens from a box
func removeLens(b boxes, l lens) {
	toBox := hash(l.label)
	i := b.index(toBox, l)

	if i != -1 {
		b[toBox] = append(b[toBox][:i], b[toBox][i+1:]...)
	}
}

func processLenses(b boxes, lenses []lens) {
	for _, l := range lenses {
		if l.operation == "=" {
			addLens(b, l)
		} else {
			removeLens(b, l)
		}
	}
}

// score returns the score of a box
func score(b boxes) int {
	var sum int
	for boxIdx, box := range b {
		for lensIndex, l := range box {
			sum += (boxIdx + 1) * (lensIndex + 1) * l.focal
		}
	}
	return sum
}

func part1() {
	fmt.Println("Part 1:")
	data := parse(utils.ReadFile("input.txt"))
	sum := hashSum(data)
	fmt.Printf("The sum of the hash of each string is %d\n", sum)
}

func part2() {
	fmt.Println("Part 2:")
	lenses := parseLenses(utils.ReadFile("input2.txt"))
	b := make(boxes)
	processLenses(b, lenses)
	fmt.Printf("The score of the boxes is %d\n", score(b))
}

func main() {
	part1()
	part2()
}
