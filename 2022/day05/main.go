package main

import (
	"fmt"
	"github.com/iamlucasvieira/aoc/utils"
	"strconv"
	"strings"
)

type stack = utils.Stack[string]

type rule struct {
	n, from, to int
}

// moveMultiple pops n boxes from stack and push them in the original order to another stack
func moveMultiple(n int, from, to *stack) error {
	var boxes []string
	for i := 0; i < n; i++ {
		v, ok := from.Pop()
		if !ok {
			return fmt.Errorf("moveMultiple: No box to move from stack")
		}
		boxes = append(boxes, v)
	}

	for i := len(boxes) - 1; i >= 0; i-- {
		to.Push(boxes[i])
	}
	return nil
}

// parse parses the input file and returns the stacks and rules
func parse(lines []string) ([]stack, []rule, error) {
	var rules []rule
	var boxes [][]string
	var stacks []stack
	var firstLine bool = true
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		if line[0] == 'm' {
			parts := strings.Split(line, " ")
			// move 1 from 2 to 1
			if len(parts) != 6 {
				return nil, nil, fmt.Errorf("parse: Wrong number of parts in %s", line)
			}

			nStrings := []string{parts[1], parts[3], parts[5]}
			nIntegers := make([]int, 3)
			for i, n := range nStrings {
				nInt, err := strconv.Atoi(n)
				if err != nil {
					return nil, nil, fmt.Errorf("parse: %s is not an int", n)
				}
				nIntegers[i] = nInt
			}

			rules = append(rules, rule{
				n:    nIntegers[0],
				from: nIntegers[1] - 1, // 1-based to 0-based
				to:   nIntegers[2] - 1, // 1-based to 0-based
			})

		} else {

			// In first line that contains boxes, we populate stacks with empty stacks
			if firstLine {
				firstLine = false

				for i := 0; i < len(line); i += 4 {
					stacks = append(stacks, stack{})
					boxes = append(boxes, []string{})
				}
			}

			// In every line that contains boxes, we populate stacks with boxes
			for i := 0; i < len(line); i += 4 {
				box := line[i+1]
				if box != ' ' && strings.Contains(line, "[") {
					boxes[i/4] = append(boxes[i/4], string(box))
				}
			}
		}
	}

	// Populate stacks with boxes
	for i := range stacks {
		for j := range boxes[i] {
			stacks[i].Push(boxes[i][len(boxes[i])-j-1])
		}
	}

	return stacks, rules, nil
}

// moveBoxes moves boxes according to the rules
func moveBoxes(stacks *[]stack, rules []rule) error {
	for _, rule := range rules {
		for i := 0; i < rule.n; i++ {
			var v, ok = (*stacks)[rule.from].Pop()
			if !ok {
				return fmt.Errorf("moveBoxes: No box to move from stack %d", rule.from)
			}
			(*stacks)[rule.to].Push(v)
		}
	}
	return nil
}

func moveBoxesWithMultiples(stacks *[]stack, rules []rule) error {
	for _, rule := range rules {
		err := moveMultiple(rule.n, &(*stacks)[rule.from], &(*stacks)[rule.to])
		if err != nil {
			return err
		}
	}
	return nil
}

// topMessage returns a string with the initial of each box on top of each stack
func topMessage(stacks []stack) string {
	var message string
	for _, stack := range stacks {
		v, ok := stack.Peek()
		if !ok {
			continue
		} else {
			message += v
		}
	}
	return message
}

func part1() string {
	fmt.Println("Part 1:")
	s, r, err := parse(utils.ReadFile("input.txt"))
	if err != nil {
		fmt.Println(err)
		return ""
	}

	err = moveBoxes(&s, r)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	m := topMessage(s)
	fmt.Printf("The message is: %s\n", m)
	return m
}

func part2() string {
	fmt.Println("Part 2:")
	s, r, err := parse(utils.ReadFile("input2.txt"))
	if err != nil {
		fmt.Println(err)
		return ""
	}

	err = moveBoxesWithMultiples(&s, r)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	m := topMessage(s)
	fmt.Printf("The message is: %s\n", m)
	return m
}

func main() {
	part1()
	part2()
}
