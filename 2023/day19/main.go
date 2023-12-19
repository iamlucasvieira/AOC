package main

import (
	"fmt"
	"github.com/iamlucasvieira/aoc/utils"
	"strconv"
	"strings"
)

type partsMap map[string]int
type partsMapRange map[string]ranges

// update updates the parts map based on the action
func (p partsMapRange) update(key string, value ranges) partsMapRange {
	var mapCopy = make(partsMapRange)

	for k, v := range p {
		mapCopy[k] = v
	}
	// make a copy of the map
	mapCopy[key] = value
	return mapCopy
}

type ranges struct {
	min int
	max int
}

// update updates the ranges based on the action
func (r ranges) update(a action) ranges {
	if a.operation == ">" && a.value > r.min {
		r.min = a.value + 1
	} else if a.operation == "<" && a.value < r.max {
		r.max = a.value - 1
	}
	return r
}

// Returns two ranges that are the result of splitting the ranges based on the action
// The first satisfies the action and the second does not
func (r ranges) split(a action) (ranges, ranges) {
	if a.operation == ">" {
		if a.value > r.max {
			return ranges{}, r
		} else if a.value < r.min {
			return r, ranges{}
		} else {
			return ranges{min: a.value + 1, max: r.max}, ranges{min: r.min, max: a.value}
		}
	} else if a.operation == "<" {
		if a.value > r.max {
			return r, ranges{}
		} else if a.value < r.min {
			return ranges{}, r
		} else {
			return ranges{min: r.min, max: a.value - 1}, ranges{min: a.value, max: r.max}
		}
	}
	return ranges{}, ranges{}
}

// isValid checks if the ranges are valid
func (r ranges) isValid() bool {
	return r.min <= r.max
}

// len returns the length of the ranges
func (r ranges) len() int {
	return r.max - r.min + 1
}

// isEmpty checks if the ranges are empty
func (r ranges) isEmpty() bool {
	return r.min == 0 && r.max == 0
}

// newPartsMapRanges creates ranges for each part in the parts map
func newPartsMapRanges(defaultMin, defaultMax int) partsMapRange {
	return partsMapRange{
		"x": {min: defaultMin, max: defaultMax},
		"m": {min: defaultMin, max: defaultMax},
		"a": {min: defaultMin, max: defaultMax},
		"s": {min: defaultMin, max: defaultMax},
	}
}

// newParts creates a map of parts amounts from a string of part
// The string must be in the format of: "{x=787,m=2655,a=1222,s=2876}"
func newParts(s string) (partsMap, error) {
	// Remove the curly braces
	s = strings.Trim(s, "{}")

	// Split the string into parts
	split := strings.Split(s, ",")

	parts := make(map[string]int)

	for _, part := range split {
		partSplit := strings.Split(part, "=")
		if len(partSplit) != 2 {
			return nil, fmt.Errorf("NewParts: Wrong format in %s", s)
		}

		partName := partSplit[0]
		partAmount, err := strconv.Atoi(partSplit[1])

		if err != nil {
			return nil, fmt.Errorf("NewParts: Wrong format in %s", s)
		}

		parts[partName] = partAmount
	}

	// Check if map contains x, m, a, s
	for _, part := range []string{"x", "m", "a", "s"} {
		if _, ok := parts[part]; !ok {
			return nil, fmt.Errorf("NewParts: Missing part %s in %s", part, s)
		}
	}

	return parts, nil
}

// rule is a struct that represents a rule in the input
type rules map[string][]action

// action is a struct that represents an action in the input
type action struct {
	part        string
	operation   string
	value       int
	destination string
}

func (a action) isEnd() bool {
	return a.operation == "" && a.value == 0 && a.part == ""
}

func (a action) evaluate(p partsMap) bool {

	if a.isEnd() {
		return false
	}

	part := p[a.part]

	switch a.operation {
	case ">":
		return part > a.value
	case "<":
		return part < a.value
	}
	return false
}

// newAction creates a new action from a string
// The string must be in the format of: "part>value:destination" or "destination" or "part<value:destination"
func newAction(s string) (action, error) {
	split := strings.Split(s, ":")
	if len(split) == 1 {
		return action{destination: s}, nil
	}

	destination := split[1]
	var operation string
	if strings.Contains(split[0], "<") {
		operation = "<"
	} else if strings.Contains(split[0], ">") {
		operation = ">"
	} else {
		return action{}, fmt.Errorf("NewAction: Wrong operation type in %s, must be '>' or '<'", s)
	}

	split = strings.Split(split[0], operation)

	if len(split) != 2 {
		return action{}, fmt.Errorf("NewAction: Operation must split into two parts in %s", s)
	}

	part := split[0]
	value, err := strconv.Atoi(split[1])

	if err != nil {
		return action{}, fmt.Errorf("NewAction: The value after the operation in %s must be an int", s)
	}

	return action{
		part:        part,
		operation:   operation,
		value:       value,
		destination: destination,
	}, nil
}

// parse parses the input into a slice of rules and a map of parts
func parse(lines []string) (rules, []partsMap, error) {
	var r = make(rules)
	var parts []partsMap

	for _, line := range lines {

		if len(line) == 0 {
			continue
		}

		if line[0] != '{' {
			// Get the rule id that goes up to the first colon
			split := strings.Split(line, "{")
			if len(split) != 2 {
				return nil, nil, fmt.Errorf("parse: Wrong format in %s", line)
			}

			ruleId := split[0]
			actionsStr := strings.Trim(split[1], "}")
			var actions []action

			for _, actionStr := range strings.Split(actionsStr, ",") {
				action, err := newAction(actionStr)
				if err != nil {
					return nil, nil, fmt.Errorf("parse: %s", err)
				}
				actions = append(actions, action)
			}

			r[ruleId] = actions

		} else {
			p, err := newParts(line)
			if err != nil {
				return nil, nil, fmt.Errorf("parse: %s", err)
			}
			parts = append(parts, p)
		}
	}
	return r, parts, nil
}

func numPartsAccepted(r rules, parts []partsMap) int {
	var numAccepted int

	for _, p := range parts {

		// Start with the 'in' rule
		var rule = "in"

		for rule != "A" && rule != "R" {
			for _, action := range r[rule] {

				if action.isEnd() || action.evaluate(p) {
					rule = action.destination
					break
				}

			}
		}

		if rule == "A" {
			numAccepted += p["x"] + p["m"] + p["a"] + p["s"]
		}
	}

	return numAccepted
}

func sumPossibleAcceptableParts(ruleName string, idx int, r rules, parts partsMapRange) int {
	// If any range in the map is invalid or empty, return 0
	for _, r := range parts {
		if !r.isValid() || r.isEmpty() {
			return 0
		}
	}

	if ruleName == "A" {
		return parts["x"].len() * parts["m"].len() * parts["a"].len() * parts["s"].len()
	} else if ruleName == "R" {
		return 0 // No parts are acceptable
	}

	// The action to be evaluated
	currentAction := r[ruleName][idx]

	// The next destination
	nextDestination := currentAction.destination

	// When the action is just to move to the next rule
	if currentAction.isEnd() {
		return sumPossibleAcceptableParts(nextDestination, 0, r, parts)
	}

	// The part that is being tested
	partTested := currentAction.part

	// Get the ranges that satisfy the current action
	satisfied, unsatisfied := parts[partTested].split(currentAction)

	// For the satisfied we move to the next destination
	satisfiedSum := sumPossibleAcceptableParts(
		nextDestination,
		0,
		r,
		parts.update(partTested, satisfied),
	)

	// For the unsatisfied we move to the next action
	unsatisfiedSum := sumPossibleAcceptableParts(
		ruleName,
		idx+1,
		r,
		parts.update(partTested, unsatisfied),
	)

	return satisfiedSum + unsatisfiedSum
}

func part1() int {
	r, p, err := parse(utils.ReadFile("input.txt"))
	if err != nil {
		panic(err)
	}
	sum := numPartsAccepted(r, p)
	fmt.Printf("Part 1: %d\n", sum)
	return sum
}

func part2() int {
	r, _, err := parse(utils.ReadFile("input2.txt"))
	if err != nil {
		panic(err)
	}

	parts := newPartsMapRanges(1, 4000)

	sum := sumPossibleAcceptableParts("in", 0, r, parts)
	fmt.Printf("Part 2: %d\n", sum)
	return sum
}
func main() {
	part1()
	part2()
}
