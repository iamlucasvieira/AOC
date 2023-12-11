package main

import (
	"fmt"
	"github.com/iamlucasvieira/aoc/utils"
	"slices"
)

// point is a struct that represents a point in a 2D plane.
type point struct {
	x int
	y int
}

type grid [][]string

// get is a method that returns the value of a point in a grid.
func (g grid) get(p point) string {
	if p.y < 0 || p.y >= len(g) {
		return ""
	}
	if p.x < 0 || p.x >= len(g[p.y]) {
		return ""
	}
	return g[p.y][p.x]
}

// add is a method that adds a point to another point.
func (p point) add(p2 point) point {
	return point{
		x: p.x + p2.x,
		y: p.y + p2.y,
	}
}

// equal is a method that checks if a point is equal to another point.
func (p point) equal(p2 point) bool {
	return p.x == p2.x && p.y == p2.y
}

// Piece is a point that contains a value and connection to 2 other point.
type piece struct {
	point
	connections []point
}

// hasConnection is a method that checks if a piece has a connection to another point.
func (p piece) hasConnection(p2 point) bool {
	for _, connection := range p.connections {
		if connection.equal(p2) {
			return true
		}
	}
	return false
}

// policy is a map that contains the connection points for a piece.
var policy = map[string][]point{
	"L": {
		{0, -1},
		{1, 0},
	},
	"7": {
		{-1, 0},
		{0, 1},
	},
	"J": {
		{-1, 0},
		{0, -1},
	},
	"F": {
		{0, 1},
		{1, 0},
	},
	"|": {
		{0, 1},
		{0, -1},
	},
	"-": {
		{-1, 0},
		{1, 0},
	},
}

// getNext is a function that returns the next point in the path. It receives the point from and the piece name and position.
func getNext(from point, name string, position point) (point, error) {
	p, err := makePiece(name, position)
	if err != nil {
		return point{}, err
	}

	if from.equal(p.connections[0]) {
		return p.connections[1], nil
	}
	if from.equal(p.connections[1]) {
		return p.connections[0], nil
	}

	return point{}, fmt.Errorf("point %v is not connected to piece %s", from, name)
}

// makePiece is a function that creates a piece from a point.
func makePiece(name string, p point) (piece, error) {

	pieceConnections, ok := policy[name]

	if !ok {
		return piece{}, fmt.Errorf("piece %s not found in policy", name)
	}

	if len(pieceConnections) != 2 {
		return piece{}, fmt.Errorf("piece %s must have 2 connections", name)
	}

	connection1 := p.add(pieceConnections[0])
	connection2 := p.add(pieceConnections[1])

	return piece{
		point: p,
		connections: []point{
			connection1,
			connection2,
		},
	}, nil
}

// parseData is a function that parses the data into a grid and finds the start point.
func parseData(lines []string) (grid, point, error) {

	var grid grid
	var start point
	for yIdx, line := range lines {
		if line == "" {
			continue
		}

		var row []string
		for xIdx, char := range line {
			if char == 'S' {
				start = point{xIdx, yIdx}
			}
			row = append(row, string(char))
		}

		grid = append(grid, row)

	}
	return grid, start, nil
}

// Find the first piece connected to the start point.
func findFirstStep(g grid, start point) (point, error) {
	pointsAround := []point{
		start.add(point{0, 1}),
		start.add(point{0, -1}),
		start.add(point{1, 0}),
		start.add(point{-1, 0}),
	}

	for _, p := range pointsAround {
		pieceName := g.get(p)
		if pieceName != "" && pieceName != "." {
			pieceAtPoint, err := makePiece(pieceName, p)
			if err != nil {
				return point{}, err
			}
			if pieceAtPoint.hasConnection(start) {
				return p, nil
			}
		}
	}
	return point{}, fmt.Errorf("no piece found connected to start point %v", start)
}

// path is a function that returns the path from a start point.
func path(g grid, start point) ([]point, error) {

	var currentPoint, previousPoint, nextPoint point
	var err error

	// The first current point is the first piece connected to the start point.
	currentPoint, err = findFirstStep(g, start)
	if err != nil {
		return nil, err
	}
	previousPoint = start

	var path = []point{previousPoint, currentPoint}
	for !currentPoint.equal(start) {
		nextPoint, err = getNext(previousPoint, g.get(currentPoint), currentPoint)
		if err != nil {
			return nil, err
		}
		path = append(path, nextPoint)
		previousPoint = currentPoint
		currentPoint = nextPoint
	}
	return path, nil
}

// largestDistance is a function that returns the largest distance from a start point (path is a cycle).
func largestDistance(g grid, start point) (int, error) {
	path, err := path(g, start)
	if err != nil {
		return 0, err
	}

	return (len(path) - 1) / 2, nil
}

func pointInPolygon(p point, g grid, polygon []point) bool {
	var nCrossing int

	// Check if point is in grid
	if value := g.get(p); value == "" {
		return false
	}

	for x := p.x + 1; x < len(g[p.y]); x++ {
		nextP := point{x, p.y}

		// Check if it encounters a piece of the polygon
		for i := 1; i < len(polygon)-1; i++ {
			previous, current, next := polygon[i-1], polygon[i], polygon[i+1]
			if !current.equal(nextP) {
				continue
			}

			if !(current.y == previous.y+1) && !(current.y == next.y+1) {
				continue
			}

			nCrossing++

		}
	}
	// If even number of crossings, point is outside
	return nCrossing%2 == 1
}

func part1() {
	fmt.Println("Part 1:")
	data, start, err := parseData(utils.ReadFile("input.txt"))

	if err != nil {
		fmt.Printf("Error parsing data: %v", err)
		return
	}

	var distance int
	distance, err = largestDistance(data, start)
	fmt.Printf("Largest distance: %d\n", distance)
}

func tilesEnclosed(g grid, polygon []point) int {
	var enclosed int
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {

			contains := slices.ContainsFunc(polygon, func(p point) bool {
				return p.equal(point{x, y})
			})

			if contains {
				continue
			}

			p := point{x, y}

			if pointInPolygon(p, g, polygon) {
				enclosed++
			}
		}
	}
	return enclosed
}

func part2() {
	fmt.Println("Part 2:")
	data, start, err := parseData(utils.ReadFile("input2.txt"))

	if err != nil {
		fmt.Printf("Error parsing data: %v", err)
		return
	}

	var polygon []point
	polygon, err = path(data, start)

	if err != nil {
		fmt.Printf("Error finding path: %v", err)
		return
	}

	enclosed := tilesEnclosed(data, polygon)

	fmt.Printf("Tiles enclosed: %d\n", enclosed)
}
func main() {
	part1()
	part2()
}
