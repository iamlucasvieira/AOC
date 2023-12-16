package main

import (
	"fmt"
	"github.com/iamlucasvieira/aoc/utils"
)

type grid = utils.Grid[*tile]
type point = utils.Point
type cache = map[string]int

var (
	left  = point{X: -1}
	right = point{X: 1}
	up    = point{Y: -1}
	down  = point{Y: 1}
)

type tile struct {
	kind      rune
	energised bool
}

// nEnergised returns the number of energised tiles in the grid
func nEnergised(g grid) int {
	var count int
	for item := range g.Iterator() {
		if item.Value.energised {
			count++
		}
	}
	return count
}

func energyString(g grid) string {
	var output string

	for _, row := range g {
		for _, item := range row {
			if item.energised {
				output += "#"
			} else {
				output += string(item.kind)
			}
		}
		output += "\n"
	}
	return output
}

// parse returns the matrix with the problem grid
func parse(input []string) grid {
	var output [][]*tile
	for _, line := range input {
		if len(line) == 0 {
			continue
		}
		var row = make([]*tile, 0, len(line))
		for _, char := range line {
			row = append(row, &tile{char, false})
		}
		output = append(output, row)
	}
	return output
}

func moveBeam(start, direction point, g grid, c cache) {

	current := start

	if g.Contains(current) && g.Get(current).kind == '.' {
		g.Get(current).energised = true
	}

	for {

		// Check if already visited the point while travelling in the same direction
		_, ok := c[current.String()+direction.String()]
		if ok {
			return
		}
		c[current.String()+direction.String()] = 1

		nextPoint := current.Add(direction)
		if !g.Contains(nextPoint) {
			return
		}
		nextPointValue := g.Get(nextPoint)
		nextPointValue.energised = true
		switch nextPointValue.kind {
		case '.':
		case '|':
			if direction == left || direction == right {
				// Split the beam into two moving in opposite directions from the next point
				moveBeam(nextPoint, up, g, c)
				moveBeam(nextPoint, down, g, c)
				return
			}
		case '-':
			if direction == up || direction == down {
				// Split the beam into two moving in opposite directions from the next point
				moveBeam(nextPoint, left, g, c)
				moveBeam(nextPoint, right, g, c)
				return
			}
		case '/':
			switch direction {
			case up:
				direction = right
			case right:
				direction = up
			case down:
				direction = left
			case left:
				direction = down
			}
		case '\\':
			switch direction {
			case up:
				direction = left
			case left:
				direction = up
			case down:
				direction = right
			case right:
				direction = down
			}
		}
		current = nextPoint
	}
}

// maxEnergy returns the maximum energy that can be obtained from the grid.
// The function initialises a beam from every tile on
// First row: heading down
// Last row: heading up
// First column: heading right
// Last column: heading left
func maxEnergy(d []string) int {
	var maxEnergy int
	table := parse(d)
	width := table.Width()
	height := table.Height()

	var move = func(start point, direction, nextPoint point, maxI int) {
		for i := 0; i < maxI; i++ {
			g := parse(d)
			startPoint := start.Add(nextPoint.Mul(point{X: i, Y: i}))
			moveBeam(startPoint, direction, g, make(cache))
			maxEnergy = max(maxEnergy, nEnergised(g))
		}

	}
	// First row
	move(point{}, down, point{X: 1}, width)

	// Last row
	move(point{Y: height - 1}, up, point{X: 1}, width)

	// First column
	move(point{}, right, point{Y: 1}, height)

	// Last column
	move(point{X: width - 1}, left, point{Y: 1}, height)

	return maxEnergy
}

func part1() {
	fmt.Println("Part 1:")
	g := parse(utils.ReadFile("input.txt"))
	moveBeam(point{X: 0, Y: 0}, right, g, make(cache))
	n := nEnergised(g)
	fmt.Printf("Number of energised tiles: %d\n", n)
}

func part2() {
	fmt.Println("Part 2:")
	d := utils.ReadFile("input2.txt")
	n := maxEnergy(d)
	fmt.Printf("Maximum energy: %d\n", n)
}

func main() {
	part1()
	part2()
}
