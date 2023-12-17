package main

import (
	"fmt"
	"github.com/iamlucasvieira/aoc/utils"
	"strconv"
)

type (
	Graph = utils.Graph[*Node]
	Point = utils.Point
)

// Node is a struct that represents a node in a graph.
type Node struct {
	Point
	Value     int
	Prev      *Node
	Direction Point
	MoveCount int
}

func (n *Node) GetValue() int {
	return n.Value
}

func (n *Node) SetValue(i int) {
	n.Value = i
}

// String is a method that returns a string representation of a Node.
// (x, y, direction, moveCount)
func (n *Node) String() string {
	return fmt.Sprintf("(%d, %d, %v, %d)", n.X, n.Y, n.Direction, n.MoveCount)
}

// Equal is a method that checks if a node is equal to another node.
func (n *Node) Equal(other utils.NodeInterface) bool {
	if n == nil || other == nil {
		return false
	}
	otherNode, ok := other.(*Node)
	if !ok {
		return false
	}
	return n.Point.Equal(otherNode.Point)
}

// Next is a method that returns the next node in a graph.
func (n *Node) Next() utils.NodeInterface {
	return n.Prev
}

// IsEmpty is a method that checks if a node is empty.
func (n *Node) IsEmpty() bool {
	return n == nil
}

// NeighboursTemplate is a function that returns the neighbors of a node in a graph.
func NeighboursTemplate(n *Node, g Graph, maxMove, turnAfter int) []*Node {
	directions := []Point{
		{0, -1}, // up
		{0, 1},  // down
		{-1, 0}, // left
		{1, 0},  // right
	}

	var currentDirection = n.Direction
	var previous = n
	if n.Prev != nil {
		previous = n.Prev
		currentDirection = n.Sub(previous.Point)
	}

	// Moving in same direction is only allowed three times in a row
	var possibleNeighbors []*Node
	for _, d := range directions {
		var moveCount = 1

		if d == currentDirection {
			if n.MoveCount == maxMove {
				continue
			} else {
				moveCount = n.MoveCount + 1
			}
		} else if n.MoveCount < turnAfter && !currentDirection.Equal(Point{}) {
			continue
		}

		neighbor := n.Add(d)

		if g.Nodes.Contains(neighbor) && !neighbor.Equal(previous.Point) {
			possibleNeighbors = append(possibleNeighbors, &Node{
				Point:     neighbor,
				Value:     n.Value + g.Get(neighbor.X, neighbor.Y).Value,
				Prev:      n,
				Direction: d,
				MoveCount: moveCount,
			})
		}
	}
	return possibleNeighbors
}

// Neighbours is a function that returns the neighbors of a node in a graph.
// Rule: Only possible to move three times in the same direction and not possible to return to previous node.
func Neighbours(n *Node, g Graph) []*Node {
	return NeighboursTemplate(n, g, 3, 0)
}

// NeighboursUltra is a function that returns the neighbors of a node in a graph.
// Rule: Only possible to turn after moving 4 times in the same direction and not possible to move than more than 10 times in the same direction.
func NeighboursUltra(n *Node, g Graph) []*Node {
	return NeighboursTemplate(n, g, 10, 4)
}

// parse is a function that parses the input into a Graph.
func parse(input []string) Graph {
	grid := make(utils.Grid[*Node], 0)
	for y, line := range input {
		if len(line) == 0 {
			continue
		}
		row := make([]*Node, len(line))
		for x, char := range line {
			// String to int
			val, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			row[x] = &Node{Point: Point{X: x, Y: y}, Value: val, Direction: Point{}}
		}
		grid = append(grid, row)
	}
	return Graph{Nodes: grid}
}

// shortestDistance is a function that returns the shortest distance between two points in a graph.
func shortestDistance(g Graph, neighbourFunc func(*Node, Graph) []*Node) int {
	h, w := g.Nodes.Height(), g.Nodes.Width()

	start := g.Get(0, 0)
	end := g.Get(w-1, h-1)

	path := utils.DijkstraTemplate[*Node](g, start, end, neighbourFunc)
	if len(path) == 0 {
		return -1
	}
	return path[len(path)-1].Value
}

func part1() {
	fmt.Println("Part 1:")
	g := parse(utils.ReadFile("input.txt"))
	loss := shortestDistance(g, Neighbours)
	fmt.Printf("The shortest path is %d\n", loss)
}

func part2() {
	fmt.Println("Part 2:")
	g := parse(utils.ReadFile("input2.txt"))
	loss := shortestDistance(g, NeighboursUltra)
	fmt.Printf("The shortest path is %d\n", loss)
}

func main() {
	part1()
	part2()
}
