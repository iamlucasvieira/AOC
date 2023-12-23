package main

import (
	"fmt"
	"strings"

	"github.com/iamlucasvieira/aoc/utils"
)

type (
	grid  = utils.Grid[string]
	point = utils.Point
)

// neighbours is a function that returns the neighbours of a point in a grid.
func neighbours(p point, g grid) []point {
	neighbours := []point{
		{X: p.X - 1, Y: p.Y},
		{X: p.X + 1, Y: p.Y},
		{X: p.X, Y: p.Y - 1},
		{X: p.X, Y: p.Y + 1},
	}

	var validNeighbours []point

	for _, n := range neighbours {

		if !g.Contains(n) {
			continue
		}

		if g.Get(n) != "#" {
			validNeighbours = append(validNeighbours, n)
		}
	}

	return validNeighbours
}

// neighboursInfinity is a function that returns the neighbours of a point in a grid.
// If a point is outside the grid, the grid is repeated infinitely.
func neighboursInfinity(p point, g grid) []point {
	if p.X < 0 {
		// Continues from right side
		p.X = g.Width() - 1
	}
	return []point{}
}

func parse(lines []string) (grid, error) {
	var g grid
	for _, line := range lines {
		if line == "" {
			continue
		}

		g = append(g, strings.Split(line, ""))

	}

	return g, nil
}

// findStart returns the start point of a grid.
func findStart(g grid) point {
	for y, row := range g {
		for x, col := range row {
			if col == "S" {
				return point{X: x, Y: y}
			}
		}
	}
	return point{}
}

// possibleEnd returns a list of possible end points after moving n steps from a point.
func possibleEnd(start point, g grid, n int) []point {
	type mapPoint map[point]bool
	var queue = make(mapPoint)

	queue[start] = true

	for i := 0; i < n; i++ {
		// Visit all points in the queue

		var newQueue = make(mapPoint)

		for k := range queue {
			for _, n := range neighbours(k, g) {
				newQueue[n] = true
			}

		}

		queue = newQueue
	}

	var endPoints = make([]point, 0, len(queue))
	for k := range queue {
		endPoints = append(endPoints, k)
	}

	return endPoints
}

func part1() int {
	fmt.Println("Part 1:")
	g, err := parse(utils.ReadFile("input.txt"))
	if err != nil {
		panic(err)
	}

	start := findStart(g)
	endPoints := possibleEnd(start, g, 64)

	fmt.Printf("There are %d possible end points after 64 steps.\n", len(endPoints))
	return len(endPoints)
}

func part2() {
	fmt.Println("Part 2:")
	// Assumption that in this input the distance formula f() forms a quadratic function§

	g, err := parse(utils.ReadFile("input2.txt"))
	if err != nil {
		panic(err)
	}

	width := g.Width()
	start := findStart(g)

	E := len(possibleEnd(start, g, 3*width))
	O := len(possibleEnd(start, g, 3*width+1))

	sa := (3*width - 3) / 2
	sb := (width - 3) / 2

	p00 := len(possibleEnd(point{X: 0, Y: 0}, g, sa))
	p0h := len(possibleEnd(point{X: 0, Y: width - 1}, g, sa))
	ph0 := len(possibleEnd(point{X: width - 1, Y: 0}, g, sa))
	pwh := len(possibleEnd(point{X: width - 1, Y: width - 1}, g, sa))

	A := p00 + p0h + ph0 + pwh

	p00b := len(possibleEnd(point{X: 0, Y: 0}, g, sb))
	p0hb := len(possibleEnd(point{X: 0, Y: width - 1}, g, sb))
	ph0b := len(possibleEnd(point{X: width - 1, Y: 0}, g, sb))
	pwhb := len(possibleEnd(point{X: width - 1, Y: width - 1}, g, sb))

	B := p00b + p0hb + ph0b + pwhb

	p0sc := len(possibleEnd(point{X: 0, Y: start.Y}, g, width))
	pswc := len(possibleEnd(point{X: start.X, Y: width - 1}, g, width))
	pwsc := len(possibleEnd(point{X: width - 1, Y: start.Y}, g, width))
	pshc := len(possibleEnd(point{X: start.X, Y: width - 1}, g, width))

	T := p0sc + pswc + pwsc + pshc

	F := func(x int) int {
		return (x-1)*(x-1)*O + x*x*E + (x-1)*A + x*B + T
	}

	result := F((16733044 - start.X) / width)

	fmt.Printf("The number of steps to reach the vault is %d.\n", result)

}

func main() {
	part1()
	part2()
}
