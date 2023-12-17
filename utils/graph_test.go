package utils

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestNode(t *testing.T) {
	node := Node{Point{1, 2}, 3, nil}

	if node.X != 1 {
		t.Fatalf("Node.X should be 1")
	}

	if node.Y != 2 {
		t.Fatalf("Node.Y should be 2")
	}

	if node.Point != (Point{1, 2}) {
		t.Fatalf("Node.Point should be (Point{1, 2})")
	}

	if node.Value != 3 {
		t.Fatalf("Node.Value should be 3")
	}
}

func TestGraph(t *testing.T) {
	data := Grid[*Node]{
		{
			&Node{Point{0, 0}, 1, nil},
			&Node{Point{0, 1}, 1, nil},
			&Node{Point{0, 2}, 1, nil},
		},
		{
			&Node{Point{1, 0}, 1, nil},
			&Node{Point{1, 1}, 1, nil},
			&Node{Point{1, 2}, 1, nil},
		},
	}

	var graph = Graph[*Node]{
		Nodes: data,
	}

	if graph.Nodes.Width() != 3 {
		t.Fatalf("graph.Nodes.Width() should be 3, but is %d", graph.Nodes.Width())
	}

	if graph.Nodes.Height() != 2 {
		t.Fatalf("graph.Nodes.Height() should be 2, but is %d", graph.Nodes.Height())
	}
}

func TestGraph_Get(t *testing.T) {
	data := [][]*Node{
		{
			&Node{Point{0, 0}, 1, nil},
			&Node{Point{0, 1}, 1, nil},
			&Node{Point{0, 2}, 1, nil},
		},
		{
			&Node{Point{1, 0}, 1, nil},
			&Node{Point{1, 1}, 1, nil},
			&Node{Point{1, 2}, 1, nil},
		},
	}

	var graph = Graph[*Node]{
		Nodes: data,
	}

	node := graph.Get(1, 1)

	if node.X != 1 {
		t.Fatalf("node.X should be 1, but is %d", node.X)
	}

	if node.Y != 1 {
		t.Fatalf("node.Y should be 1, but is %d", node.Y)
	}

	if node.Value != 1 {
		t.Fatalf("node.Value should be 1, but is %d", node.Value)
	}
}

func TestNewNode(t *testing.T) {
	node := NewNode(1, 2, 3)

	if node.X != 1 {
		t.Fatalf("Node.X should be 1")
	}

	if node.Y != 2 {
		t.Fatalf("Node.Y should be 2")
	}

	if node.Point != (Point{1, 2}) {
		t.Fatalf("Node.Point should be (Point{1, 2})")
	}

	if node.Value != 3 {
		t.Fatalf("Node.Value should be 3")
	}
}

func TestNewGraph(t *testing.T) {
	data := [][]int{
		{1, 1, 1},
		{1, 1, 1},
	}

	graph := NewGraph(data)

	if graph.Nodes.Width() != 3 {
		t.Fatalf("graph.Nodes.Width() should be 3, but is %d", graph.Nodes.Width())
	}

	if graph.Nodes.Height() != 2 {
		t.Fatalf("graph.Nodes.Height() should be 2, but is %d", graph.Nodes.Height())
	}

	for y, row := range data {
		for x, value := range row {
			if gotValue := graph.Nodes[y][x].Value; gotValue != value {
				t.Fatalf("graph.Nodes[%d][%d].Value should be %d, but is %d", y, x, value, gotValue)
			}

			if gotX := graph.Nodes[y][x].X; gotX != x {
				t.Fatalf("graph.Nodes[%d][%d].X should be %d, but is %d", y, x, x, gotX)
			}

			if gotY := graph.Nodes[y][x].Y; gotY != y {
				t.Fatalf("graph.Nodes[%d][%d].Y should be %d, but is %d", y, x, y, gotY)
			}
		}
	}
}

func TestNeighbors(t *testing.T) {
	data := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}

	graph := NewGraph(data)

	testCases := []struct {
		node          *Node
		wantNeighbors []*Node
	}{
		{
			node: graph.Get(0, 0),
			wantNeighbors: []*Node{
				graph.Get(0, 1),
				graph.Get(1, 0),
			},
		},
		{
			node: graph.Get(1, 1),
			wantNeighbors: []*Node{
				graph.Get(1, 0),
				graph.Get(1, 2),
				graph.Get(0, 1),
				graph.Get(2, 1),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.node.String(), func(t *testing.T) {
			gotNeighbors := Neighbors(tc.node, graph)

			if len(gotNeighbors) != len(tc.wantNeighbors) {
				t.Fatalf("len(gotNeighbors) should be %d, but is %d", len(tc.wantNeighbors), len(gotNeighbors))
			}

			for i, wantNeighbor := range tc.wantNeighbors {
				if gotNeighbors[i].Point != wantNeighbor.Point {
					t.Fatalf("gotNeighbors[%d] should be %v, but is %v", i, wantNeighbor, gotNeighbors[i])
				}
			}
		})
	}

}
func TestNewNodeQueue(t *testing.T) {
	nq := NewNodeQueue[*Node]()
	heap.Init(nq)
	if nq.Len() != 0 {
		t.Fatalf("nq.Len() should be 0, but is %d", nq.Len())
	}
}

func TestNodeQueue_Len(t *testing.T) {
	nq := NewNodeQueue[*Node]()
	heap.Init(nq)

	if nq.Len() != 0 {
		t.Fatalf("nq.Len() should be 0, but is %d", nq.Len())
	}

	heap.Push(nq, &Node{Point{0, 0}, 1, nil})
	if nq.Len() != 1 {
		t.Fatalf("nq.Len() should be 1, but is %d", nq.Len())
	}
}

func TestNodeQueue_Less(t *testing.T) {
	nq := NewNodeQueue[*Node]()
	heap.Init(nq)

	heap.Push(nq, &Node{Point{0, 0}, 1, nil})
	heap.Push(nq, &Node{Point{0, 0}, 2, nil})

	if !nq.Less(0, 1) {
		t.Fatalf("nq.Less(0, 1) should be false, but is %t", nq.Less(0, 1))
	}

	if nq.Less(1, 0) {
		t.Fatalf("nq.Less(1, 0) should be true")
	}
}

func TestNodeQueue_Swap(t *testing.T) {
	nq := NewNodeQueue[*Node]()
	heap.Init(nq)

	heap.Push(nq, &Node{Point{0, 0}, 1, nil})
	heap.Push(nq, &Node{Point{0, 0}, 2, nil})

	nq.Swap(0, 1)

	if nq.Less(0, 1) {
		t.Fatalf("nq.Less(0, 1) should be false, but is %t", nq.Less(0, 1))
	}

	if !nq.Less(1, 0) {
		t.Fatalf("nq.Less(1, 0) should be true")
	}
}

func TestNodeQueue_Push(t *testing.T) {
	nq := NewNodeQueue[*Node]()
	heap.Init(nq)

	heap.Push(nq, &Node{Point{0, 0}, 1, nil})
	heap.Push(nq, &Node{Point{0, 0}, 2, nil})

	if nq.Len() != 2 {
		t.Fatalf("nq.Len() should be 2, but is %d", nq.Len())
	}
}

func TestNodeQueue_Pop(t *testing.T) {
	nq := NewNodeQueue[*Node]()
	heap.Init(nq)

	heap.Push(nq, &Node{Point{0, 0}, 1, nil})
	heap.Push(nq, &Node{Point{0, 0}, 2, nil})

	if nq.Len() != 2 {
		t.Fatalf("nq.Len() should be 2, but is %d", nq.Len())
	}

	heap.Pop(nq)

	if nq.Len() != 1 {
		t.Fatalf("nq.Len() should be 1, but is %d", nq.Len())
	}
}

func TestReconstructPath(t *testing.T) {
	n1 := NewNode(0, 0, 1)
	n2 := NewNode(0, 1, 2)
	n3 := NewNode(0, 2, 3)
	n4 := NewNode(1, 0, 4)

	n1.Prev = &n2
	n2.Prev = &n3
	n3.Prev = &n4

	path := ReconstructPath(&n1)

	if len(path) != 4 {
		t.Fatalf("len(path) should be 4, but is %d", len(path))
	}

	if path[0].Value != 4 {
		t.Fatalf("path[0].Value should be 4, but is %d", path[0].Value)
	}

	if path[1].Value != 3 {
		t.Fatalf("path[1].Value should be 3, but is %d", path[1].Value)
	}

	if path[2].Value != 2 {
		t.Fatalf("path[2].Value should be 2, but is %d", path[2].Value)
	}

	if path[3].Value != 1 {
		t.Fatalf("path[3].Value should be 1, but is %d", path[3].Value)
	}
}

func TestDijkstra(t *testing.T) {
	g := NewGraph([][]int{
		{1, 2, 3},
		{4, 5, 6},
	})

	path := Dijkstra(g, g.Get(0, 0), g.Get(2, 1))

	if len(path) != 4 {
		t.Fatalf("len(path) should be 4, but is %d", len(path))
	}

	if path[3].Value != 11 {
		t.Fatalf("path[0].Value should be 1, but is %d", path[0].Value)
	}
}

func TestDijkstraShortestDistance(t *testing.T) {
	g1 := NewGraph([][]int{
		{1, 1, 1, 1, 1, 1},
		{4, 5, 6, 7, 8, 1},
		{7, 8, 9, 10, 11, 12},
	})

	g2 := NewGraph([][]int{
		{1, 12, 1, 1, 1, 1},
		{1, 12, 1, 12, 12, 1},
		{1, 1, 1, 12, 12, 12},
	})

	g3 := NewGraph([][]int{
		{1, 2, 3, 4, 5, 6},
		{1, 5, 6, 7, 8, 9},
		{1, 1, 1, 1, 1, 12},
	})

	testCases := []struct {
		graph Graph[*Node]
		want  int
	}{
		{g1, 18},
		{g2, 22},
		{g3, 18},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("%v", idx), func(t *testing.T) {
			got := DijkstraShortestDistance(tc.graph, tc.graph.Get(0, 0), tc.graph.Get(5, 2))
			if got != tc.want {
				t.Fatalf("got should be %d, but is %d", tc.want, got)
			}
		})
	}
}

func TestNode_Equal(t *testing.T) {
	n1 := NewNode(0, 0, 1)
	n2 := NewNode(0, 0, 2)
	n3 := NewNode(0, 1, 3)
	// Equal should only compare Point

	if !n1.Equal(&n2) {
		t.Fatalf("n1.Equal(n1.Point) should be true")
	}

	if n1.Equal(&n3) {
		t.Fatalf("n1.Equal(n2.Point) should be false")
	}
}

func TestNode_GetValue(t *testing.T) {
	n := NewNode(0, 0, 123)

	if n.GetValue() != 123 {
		t.Fatalf("n.GetValue() should be 123, but is %d", n.GetValue())
	}
}

func TestNode_IsEmpty(t *testing.T) {
	n := NewNode(0, 0, 123)
	if n.IsEmpty() {
		t.Fatalf("n.IsEmpty() should be false")
	}

	if !n.Prev.IsEmpty() {
		t.Fatalf("n.IsEmpty() should be true")
	}
}

func TestNode_Next(t *testing.T) {
	n1 := NewNode(0, 0, 1)
	n2 := NewNode(0, 1, 2)
	n3 := NewNode(0, 2, 3)

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

func TestNode_SetValue(t *testing.T) {
	n := NewNode(0, 0, 123)
	n.SetValue(456)

	if n.GetValue() != 456 {
		t.Fatalf("n.GetValue() should be 456, but is %d", n.GetValue())
	}
}

func TestNode_String(t *testing.T) {
	n := NewNode(0, 0, 123)

	if n.String() != "(0, 0)" {
		t.Fatalf("n.String() should be (0, 0), but is %s", n.String())
	}
}
