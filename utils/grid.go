package utils

import (
	"fmt"
	"math"
)

type Grid[T any] [][]T

// Contains is a method that checks if a point is contained in a Grid.
func (g Grid[T]) Contains(p Point) bool {
	if p.Y < 0 || p.Y >= len(g) {
		return false
	}
	if p.X < 0 || p.X >= len(g[p.Y]) {
		return false
	}
	return true
}

// ContainsXY is a method that checks if a point is contained in a Grid.
func (g Grid[T]) ContainsXY(x, y int) bool {
	return g.Contains(Point{x, y})
}

// Width is a method that returns the width of a Grid.
func (g Grid[T]) Width() int {
	if len(g) == 0 {
		return 0
	}
	return len(g[0])
}

// Height is a method that returns the height of a Grid.
func (g Grid[T]) Height() int {
	return len(g)
}

// Element is a struct that represents an element in a Grid.
type Element[T any] struct {
	Value T
	Index Point
}

// Iterator returns a channel that iterates over the Grid.
// Elements are sent through the channel along with their coordinates.
func (g Grid[T]) Iterator() <-chan Element[T] {
	ch := make(chan Element[T])

	go func() {
		defer close(ch)
		for y, row := range g {
			for x, item := range row {
				ch <- Element[T]{item, Point{x, y}}
			}
		}
	}()

	return ch
}

// Get is a method that returns the value of a point in a Grid.
func (g Grid[T]) Get(p Point) T {
	return g[p.Y][p.X]
}

// Set is a method that sets the value of a point in a Grid.
func (g Grid[T]) Set(p Point, value T) {
	g[p.Y][p.X] = value
}

// Point is a struct that represents a point in a 2D plane.
type Point struct {
	X, Y int
}

// String is a method that returns a string representation of a point.
func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

// Equal is a method that checks if a point is equal to another point.
func (p Point) Equal(p2 Point) bool {
	return p.X == p2.X && p.Y == p2.Y
}

// Add is a method that adds a point to another point.
func (p Point) Add(p2 Point) Point {
	return Point{
		X: p.X + p2.X,
		Y: p.Y + p2.Y,
	}
}

// Sub is a method that subtracts a point from another point.
func (p Point) Sub(p2 Point) Point {
	return Point{
		X: p.X - p2.X,
		Y: p.Y - p2.Y,
	}
}

// Mul is a method that multiplies a point by another point.
func (p Point) Mul(p2 Point) Point {
	return Point{
		X: p.X * p2.X,
		Y: p.Y * p2.Y,
	}
}

// Distance is a method that returns the distance between two points.
func (p Point) Distance(p2 Point) int {
	return int(math.Abs(float64(p.X-p2.X)) + math.Abs(float64(p.Y-p2.Y)))
}
