package main

import (
	"fmt"
	"github.com/iamlucasvieira/aoc/utils"
	"strconv"
	"strings"
)

type sequence struct {
	start, end int
}

// parse parses the input into a slice of sequences.
func parse(input []string) ([][2]sequence, error) {
	var data [][2]sequence
	for _, line := range input {
		if line == "" {
			continue
		}
		var seq = [2]sequence{}

		parts := strings.Split(line, ",")

		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid input: %s", line)
		}

		elf1 := strings.Split(parts[0], "-")
		elf2 := strings.Split(parts[1], "-")

		for i, elf := range [2][]string{elf1, elf2} {
			if len(elf) != 2 {
				return nil, fmt.Errorf("invalid input: %s", line)
			}
			start, err := strconv.Atoi(elf[0])
			if err != nil {
				return nil, fmt.Errorf("invalid input: %s", line)
			}

			var end int
			end, err = strconv.Atoi(elf[1])
			if err != nil {
				return nil, fmt.Errorf("invalid input: %s", line)
			}
			seq[i] = sequence{start, end}

		}
		data = append(data, seq)
	}
	return data, nil
}

// isWithin returns true if any of the sequence contains its pair.
func isWithin(values [2]sequence) bool {
	within := func(a, b sequence) bool {
		return a.start >= b.start && a.end <= b.end
	}
	return within(values[0], values[1]) || within(values[1], values[0])
}

// nWithin returns the number of sequences that contain their pair.
func nWithin(values [][2]sequence) int {
	var count int
	for _, v := range values {
		if isWithin(v) {
			count++
		}
	}
	return count
}

// hasOverlap returns true if any of the sequence overlaps its pair.
func hasOverlap(values [2]sequence) bool {
	overlap := func(a, b sequence) bool {
		return a.start >= b.start && a.start <= b.end
	}
	return overlap(values[0], values[1]) || overlap(values[1], values[0])
}

// nOverlap returns the number of sequences that overlap their pair.
func nOverlap(values [][2]sequence) int {
	var count int
	for _, v := range values {
		if hasOverlap(v) {
			count++
		}
	}
	return count
}
func part1() int {
	fmt.Println("Part 1:")
	data, err := parse(utils.ReadFile("input.txt"))

	if err != nil {
		panic(err)
	}
	sum := nWithin(data)
	fmt.Printf("Sum: %d\n", sum)
	return sum
}

func part2() int {
	fmt.Println("Part 2:")
	data, err := parse(utils.ReadFile("input2.txt"))

	if err != nil {
		panic(err)
	}
	sum := nOverlap(data)
	fmt.Printf("Sum: %d\n", sum)
	return sum
}

func main() {
	part1()
	part2()
}
