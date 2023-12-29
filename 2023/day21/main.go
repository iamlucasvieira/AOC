package main

import (
	"fmt"
	"math"
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
	var steps = 26501365

	g, err := parse(utils.ReadFile("input2.txt"))
	if err != nil {
		panic(err)
	}

	// Assumption that the grid is square
	if g.Width() != g.Height() {
		panic("Grid is not square")
	}

	// Assumption that steps is equal
	width := g.Width()
	start := findStart(g)

	// Assumption that the start point is in the middle of the grid
	if start.X != width/2 || start.Y != width/2 {
		panic("Start point is not in the middle of the grid")
	}

	// Assumption that the steps is equal to w * n + w/2
	if steps%width != width/2 {
		panic("Steps is not equal to w * n + w/2")
	}

	// Half Width of the diamond formed by the repeated grid
	var diamondWidth = steps/width - 1

	// Number of even and odd grids
	var oddGrids = int(math.Pow(float64(diamondWidth/2*2+1), 2))
	var evenGrids = int(math.Pow(float64((diamondWidth+1)/2*2), 2))

	// Number of points reached in even and odd grids
	oddPoints := possibleEnd(start, g, width*2+1) // Just needs to be a large enough number
	evenPoints := possibleEnd(start, g, width*2)

	// Number of points reached in the corners
	cornerTop := possibleEnd(point{X: width - 1, Y: start.Y}, g, width-1)
	cornerRight := possibleEnd(point{X: start.X, Y: 0}, g, width-1)
	cornerBottom := possibleEnd(point{X: 0, Y: start.Y}, g, width-1)
	cornerLeft := possibleEnd(point{X: start.X, Y: width - 1}, g, width-1)

	// Small Segments
	segmentTopRight := possibleEnd(point{X: width - 1, Y: 0}, g, width/2-1)
	segmentBottomRight := possibleEnd(point{X: 0, Y: 0}, g, width/2-1)
	segmentTopLeft := possibleEnd(point{X: width - 1, Y: width - 1}, g, width/2-1)
	segmentBottomLeft := possibleEnd(point{X: 0, Y: width - 1}, g, width/2-1)

	// Large Segments
	segmentLargeTopRight := possibleEnd(point{X: width - 1, Y: 0}, g, width*3/2-1)
	segmentLargeBottomRight := possibleEnd(point{X: 0, Y: 0}, g, width*3/2-1)
	segmentLargeTopLeft := possibleEnd(point{X: width - 1, Y: width - 1}, g, width*3/2-1)
	segmentLargeBottomLeft := possibleEnd(point{X: 0, Y: width - 1}, g, width*3/2-1)

	// Number of small segments
	var smallSegments = (diamondWidth + 1)

	// Number of Large Segments
	var largeSegments = diamondWidth

	// Number of points reached in the sides
	allPoints := 0
	allPoints += oddGrids * len(oddPoints)
	allPoints += evenGrids * len(evenPoints)
	allPoints += len(cornerTop)
	allPoints += len(cornerRight)
	allPoints += len(cornerBottom)
	allPoints += len(cornerLeft)
	allPoints += len(segmentTopRight) * smallSegments
	allPoints += len(segmentBottomRight) * smallSegments
	allPoints += len(segmentTopLeft) * smallSegments
	allPoints += len(segmentBottomLeft) * smallSegments
	allPoints += len(segmentLargeTopRight) * largeSegments
	allPoints += len(segmentLargeBottomRight) * largeSegments
	allPoints += len(segmentLargeTopLeft) * largeSegments
	allPoints += len(segmentLargeBottomLeft) * largeSegments
	fmt.Printf("There are %d points reached after %d steps.\n", allPoints, steps)

}

func main() {
	part1()
	part2()
}
