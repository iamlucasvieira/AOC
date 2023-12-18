package main

import (
	"fmt"
	"github.com/iamlucasvieira/aoc/utils"
	"strconv"
	"strings"
)

type (
	point = utils.Point
)

var directions = map[string]point{
	"U": {0, -1},
	"D": {0, 1},
	"L": {-1, 0},
	"R": {1, 0},
}

type colorPoint struct {
	point
	color string
}

type command struct {
	direction point
	steps     int
	color     string
}

func parse(lines []string) ([]command, error) {
	var commands []command
	for _, line := range lines {
		if line == "" {
			continue
		}

		// split line
		parts := strings.Split(line, " ")

		if len(parts) != 3 {
			return nil, fmt.Errorf("invalid line: %v", line)
		}
		direction, ok := directions[parts[0]]

		if !ok {
			return nil, fmt.Errorf("invalid direction: %v", parts[0])
		}

		steps, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid steps: %v", parts[1])
		}

		color := parts[2]
		// Remove the parenthesis
		color = color[1 : len(color)-1]

		commands = append(commands, command{direction, steps, color})
	}
	return commands, nil
}

func buildPath(commands []command) []colorPoint {
	var path []colorPoint
	var current colorPoint

	for _, command := range commands {
		for i := 0; i < command.steps; i++ {
			current.point = current.point.Add(command.direction)
			current.color = command.color
			path = append(path, current)
		}
	}
	return path
}

// polygonArea returns the area of a polygon
func polygonArea(points []colorPoint) int {
	var area int
	nPoints := len(points)
	for i := 0; i <= nPoints-1; i++ {
		area += points[i].X*points[(i+1)%nPoints].Y - points[(i+1)%nPoints].X*points[i].Y
	}

	if area < 0 {
		area = -area
	}
	return area / 2
}

// pickleTheorem returns the number of points inside a polygon and the border
func pickleTheorem(points []colorPoint) int {
	var area = polygonArea(points)
	return area + len(points)/2 + 1
}

func part1() int {
	fmt.Println("Part 1:")
	commands, err := parse(utils.ReadFile("input.txt"))
	if err != nil {
		panic(err)
	}

	path := buildPath(commands)
	count := pickleTheorem(path)
	fmt.Printf("The number of points is %d\n", count)
	return count
}

// commandsFromHex converts the commands from hex to decimal
func commandsFromHex(commands []command) ([]command, error) {
	var newCommands []command
	for c := range commands {
		// First 5 digits is the distance and final digit is the direction: 0->R, 1->D, 2->L, 3->U

		if len(commands[c].color) != 7 {
			return nil, fmt.Errorf("invalid color: %v", commands[c].color)
		}

		distanceStr := commands[c].color[1:6]
		directionStr := commands[c].color[6:7]

		// Convert hex to int
		distance, err := strconv.ParseInt(distanceStr, 16, 64)

		if err != nil {
			return nil, fmt.Errorf("invalid distance: %v", distanceStr)
		}

		var direction point

		switch directionStr {
		case "0":
			direction = directions["R"]
		case "1":
			direction = directions["D"]
		case "2":
			direction = directions["L"]
		case "3":
			direction = directions["U"]
		}

		newCommands = append(newCommands, command{direction, int(distance), commands[c].color})
	}
	return newCommands, nil
}

func part2() int {
	fmt.Println("Part 2:")
	commands, err := parse(utils.ReadFile("input2.txt"))
	if err != nil {
		panic(err)
	}

	commands, err = commandsFromHex(commands)
	if err != nil {
		panic(err)
	}

	path := buildPath(commands)
	count := pickleTheorem(path)
	fmt.Printf("The number of points is %d\n", count)
	return count
}

func main() {
	part1()
	part2()
}
