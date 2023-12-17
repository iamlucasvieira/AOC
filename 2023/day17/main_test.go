package main

import (
	"testing"
)

var mockData = []string{
	"2413432311323",
	"3215453535623",
	"3255245654254",
	"3446585845452",
	"4546657867536",
	"1438598798454",
	"4457876987766",
	"3637877979653",
	"4654967986887",
	"4564679986453",
	"1224686865563",
	"2546548887735",
	"4322674655533",
}

func TestParse(t *testing.T) {
	graph := parse(mockData)

	if graph.Nodes.Width() != 13 {
		t.Fatalf("graph.Nodes.Width() should be 13, but is %d", graph.Nodes.Width())
	}

	if graph.Nodes.Height() != 13 {
		t.Fatalf("graph.Nodes.Height() should be 13, but is %d", graph.Nodes.Height())
	}

	if graph.Get(0, 0).Value != 2 {
		t.Fatalf("graph.Nodes.Get(0, 0).Value should be 2, but is %d", graph.Get(0, 0).Value)
	}
}

func TestNeighbours(t *testing.T) {
	g1 := Graph{
		Nodes: [][]*Node{
			{
				&Node{Value: 1},
				&Node{Point: Point{X: 1}, Value: 1},
				&Node{Point: Point{X: 2}, Value: 1},
				&Node{Point: Point{X: 3}, Value: 1},
			},
			{
				&Node{Point: Point{Y: 1}, Value: 1},
				&Node{Point: Point{X: 1, Y: 1}, Value: 1},
				&Node{Point: Point{X: 2, Y: 1}, Value: 1},
				&Node{Point: Point{X: 3, Y: 1}, Value: 1},
			},
		},
	}

	// first row: Make node 3 be followed by node 2, 1, 0
	g1.Get(3, 0).Prev = g1.Get(2, 0)
	g1.Get(2, 0).Prev = g1.Get(1, 0)
	g1.Get(1, 0).Prev = g1.Get(0, 0)

	// With three horizontal movement, the only possible direction from node 3 is down
	neighbours := Neighbours(g1.Get(3, 0), g1)

	if len(neighbours) != 1 {
		t.Fatalf("len(neighbours) should be 1, but is %d", len(neighbours))
	}

	for _, n := range neighbours {
		if n.Point.Equal(Point{X: 4}) {
			t.Fatalf("neighbours should not contain (0, 0)")
		}
	}
}

func TestShortestPath(t *testing.T) {
	graph := parse(mockData)

	if distance := shortestDistance(graph, Neighbours); distance != 102 {
		t.Fatalf("shortestDistance should be 102, but is %d", distance)
	}
}

func TestShortestPathUltra(t *testing.T) {
	graph := parse(mockData)

	if distance := shortestDistance(graph, NeighboursUltra); distance != 94 {
		t.Fatalf("shortestDistance should be 94, but is %d", distance)
	}
}

func TestNode_Equal(t *testing.T) {
	n1 := Node{Point: Point{X: 1, Y: 2}, Value: 1}
	n2 := Node{Point: Point{X: 1, Y: 2}, Value: 2}
	n3 := Node{Point: Point{X: 1, Y: 3}, Value: 3}

	if !n1.Equal(&n2) {
		t.Fatalf("n1 should be equal to n2.Point")
	}

	if n1.Equal(&n3) {
		t.Fatalf("n1 should not be equal to n3.Point")
	}
}

func TestNode_GetValue(t *testing.T) {
	n := Node{Value: 123}

	if n.GetValue() != 123 {
		t.Fatalf("n.GetValue() should be 123, but is %d", n.GetValue())
	}
}

func TestNode_SetValue(t *testing.T) {
	n := Node{Value: 123}
	n.SetValue(456)

	if n.GetValue() != 456 {
		t.Fatalf("n.GetValue() should be 456, but is %d", n.GetValue())
	}
}

func TestNode_IsEmpty(t *testing.T) {
	n := Node{Value: 123}
	if n.IsEmpty() {
		t.Fatalf("n.IsEmpty() should be false")
	}

	if !n.Prev.IsEmpty() {
		t.Fatalf("n.IsEmpty() should be true")
	}
}

func TestNode_Next(t *testing.T) {
	n1 := Node{Value: 1}
	n2 := Node{Value: 2}
	n3 := Node{Value: 3}

	n1.Prev = &n2
	n2.Prev = &n3

	if n1.Next() != &n2 {
		t.Fatalf("n1.Next() should be &n2")
	}

	if n2.Next() != &n3 {
		t.Fatalf("n2.Next() should be &n3")
	}

	if !n3.Next().IsEmpty() {
		t.Fatalf("n3.Next() should be nil")
	}
}

func TestNode_String(t *testing.T) {
	n := Node{Point: Point{X: 1, Y: 2}, Value: 123, Direction: Point{X: 3, Y: 4}, MoveCount: 5}

	if n.String() != "(1, 2, (3, 4), 5)" {
		t.Fatalf("n.String() should be (1, 2), but is %s", n.String())
	}
}
