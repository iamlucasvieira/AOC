package main

import (
	"fmt"
	"github.com/iamlucasvieira/aoc/utils"
	"strings"
)

type lrPair [2]string
type lrPairMap map[string]lrPair

// parseInstructions parses the input and returns the instructions and the
// L becomes 0 and R becomes 1.
func parseInstructions(input string) ([]int, error) {
	lrMap := map[string]int{
		"L": 0,
		"R": 1,
	}

	// Create an array of instructions with length of the input.
	instructions := make([]int, len(input))

	for idx, instructionStr := range input {
		instruction, ok := lrMap[string(instructionStr)]

		if !ok {
			return nil, fmt.Errorf("invalid instruction: %q", instructionStr)
		}

		instructions[idx] = instruction
	}

	return instructions, nil
}

// parseOption parses the input and returns the option name and the left and right values.
// AAA = (BBB, CCC) should return  "AAA", ("BBB", "CCC"), nil
func parseOption(input string) (string, lrPair, error) {

	// Remove spaces and \n from input
	input = strings.Replace(input, " ", "", -1)
	input = strings.Replace(input, "\n", "", -1)

	// Split input by "="
	parts := strings.Split(input, "=")

	if len(parts) != 2 {
		return "", lrPair{}, fmt.Errorf("invalid input: %q", input)
	}

	name := parts[0]
	values := parts[1]

	// Remove "(" and ")" from values with Trim
	values = strings.Trim(values, "()")

	// Split values by ","
	valuesParts := strings.Split(values, ",")

	if len(valuesParts) != 2 {
		return "", lrPair{}, fmt.Errorf("invalid values: %q", values)
	}

	return name, lrPair(valuesParts), nil
}

// parseData parses the input and returns the instructions, the map of options, and the start position.
func parseData(lines []string) ([]int, lrPairMap, error) {

	// First non-empty line is the instructions
	// All other lines that contain = are options

	var instructions []int
	var options = make(lrPairMap)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		if len(instructions) == 0 {
			var err error
			instructions, err = parseInstructions(line)
			if err != nil {
				return nil, nil, err
			}
			continue
		}

		if strings.Contains(line, "=") {
			name, values, err := parseOption(line)
			if err != nil {
				return nil, nil, err
			}

			options[name] = values
		}
	}
	return instructions, options, nil
}

func stepsTo(instructions []int, options lrPairMap, start string, endFunc func(string) bool) (int, error) {
	var steps int
	var currentOption = start

	// Loop through instructions multiple times until we reach Z
	for {
		if endFunc(currentOption) {
			return steps, nil
		}

		// Get the current instruction. Cycle through the instruction indexes.
		instruction := instructions[steps%len(instructions)]

		// Get the current option
		currentOptionPair := options[currentOption]

		// Get the next option
		currentOption = currentOptionPair[instruction]

		steps++
	}
}

func stepsToZZZ(instructions []int, options lrPairMap) (int, error) {
	endFunc := func(currentOption string) bool {
		return currentOption == "ZZZ"
	}

	return stepsTo(instructions, options, "AAA", endFunc)
}

func stepsToZ(instructions []int, options lrPairMap) (int, error) {
	var steps []int

	endFunc := func(currentOption string) bool {
		if !strings.HasSuffix(currentOption, "Z") {
			return false
		}
		return true
	}

	for k := range options {
		if strings.HasSuffix(k, "A") {
			step, err := stepsTo(instructions, options, k, endFunc)
			if err != nil {
				return 0, err
			}
			steps = append(steps, step)

		}
	}

	return utils.LCM(steps...), nil
}

func part1() {
	fmt.Println("Part 1:")
	instructions, options, err := parseData(utils.ReadFile("input.txt"))
	if err != nil {
		panic(err)
	}

	var steps int
	steps, err = stepsToZZZ(instructions, options)
	if err != nil {
		panic(err)
	}

	fmt.Printf("It takes %d steps to reach Z\n", steps)
}

func part2() {
	fmt.Println("Part 2:")
	instructions, options, err := parseData(utils.ReadFile("input2.txt"))
	if err != nil {
		panic(err)
	}

	var steps int
	steps, err = stepsToZ(instructions, options)
	if err != nil {
		panic(err)
	}

	fmt.Printf("It takes %d steps to reach Z\n", steps)
}
func main() {
	part1()
	part2()
}
